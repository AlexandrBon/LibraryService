package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	LibraryApp "libraryService/internal/libraryApp"
	grpcPort "libraryService/internal/ports/grpc"
	"libraryService/internal/storage/mysql"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := ":" + os.Getenv("PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	eg, ctx := errgroup.WithContext(context.Background())
	sigQuit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			log.Printf("captured signal: %v\n", s)
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	storage, err := mysql.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpcPort.NewGRPCServer(lis, LibraryApp.NewLibraryApp(storage))

	eg.Go(func() error {
		log.Printf("starting server, listening on %s\n", port)
		defer log.Printf("close server listening on %s\n", port)

		errCh := make(chan error)

		defer func() {
			grpcServer.GetServer().GracefulStop()
			_ = lis.Close()

			close(errCh)
		}()

		go func() {
			if err := grpcServer.Listen(); err != nil {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("server can't listen and serve requests: %w", err)
		}
	})

	if err := eg.Wait(); err != nil {
		log.Printf("gracefully shutting down the server: %s\n", err.Error())
	}

	log.Println("server was successfully shutdown")
}
