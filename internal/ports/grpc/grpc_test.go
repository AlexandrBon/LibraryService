package grpc

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	LibraryApp "libraryService/internal/libraryApp"
	"libraryService/internal/storage/mysql"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func getClient(t *testing.T) (LibraryServiceClient, context.Context) {
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		_ = lis.Close()
	})

	storage, err := mysql.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	srv := NewGRPCServer(lis, LibraryApp.NewLibraryApp(storage))
	t.Cleanup(func() {
		srv.Stop()
	})

	go func() {
		assert.NoError(t, srv.Listen(), "srv.Serve")
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err, "main.DialContext")

	t.Cleanup(func() {
		_ = conn.Close()
	})

	return NewLibraryServiceClient(conn), ctx
}

func TestGRPC_GetBooksByAuthorName(t *testing.T) {
	client, ctx := getClient(t)

	response, err := client.GetBooksByAuthorName(ctx, &BooksByAuthorNameRequest{AuthorName: "a_1"})
	if err != nil {
		return
	}
	assert.NoError(t, err)

	assert.Len(t, response.Books, 1)
	book := response.Books[0]
	assert.Equal(t, book.Id, int32(1))
	assert.Equal(t, book.Title, "t_1")
	assert.Equal(t, book.PageCount, int32(40))
	assert.Equal(t, book.PublishingYear, int32(1988))
}

func TestGRPC_GetAuthorsByBookTitle_NoBook(t *testing.T) {
	client, ctx := getClient(t)

	_, err := client.GetAuthorsByBookTitle(ctx, &AuthorsByBookTitleRequest{Title: "t_100"})
	if err != nil {
		return
	}
	assert.Error(t, err, LibraryApp.ErrBadRequest)
}
