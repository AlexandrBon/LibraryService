package tests

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	LibraryApp "libraryService/internal/libraryApp"
	grpcPort "libraryService/internal/ports/grpc"
	"libraryService/internal/storage/mysql"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func getClient(t *testing.T) (grpcPort.LibraryServiceClient, context.Context) {
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		_ = lis.Close()
	})

	storage, err := mysql.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	srv := grpcPort.NewGRPCServer(lis, LibraryApp.NewLibraryApp(storage))
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

	return grpcPort.NewLibraryServiceClient(conn), ctx
}

func TestGRPC_GetBooksByAuthorName(t *testing.T) {
	client, ctx := getClient(t)

	response, err := client.GetBooksByAuthorName(ctx, &grpcPort.BooksByAuthorNameRequest{AuthorName: "a_1"})
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

func TestGRPC_GetAuthorsByBookTitle(t *testing.T) {
	client, ctx := getClient(t)

	response, err := client.GetBooksByAuthorName(ctx, &grpcPort.BooksByAuthorNameRequest{AuthorName: "a_1"})
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

/*func TestGRPC_ChangeAdStatus(t *testing.T) {
	client, ctx := getClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "oleg@mail.ru"})
	assert.NoError(t, err, "client.ChangeAdStatus")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{UserId: 0, Title: "hello", Text: "world"})
	assert.NoError(t, err)

	response, err := client.UpdateAd(ctx, &grpcPort.UpdateAdRequest{AdId: 0, Title: "super", Text: "cat", UserId: 0})
	assert.NoError(t, err)

	assert.Zero(t, response.Id)
	assert.Equal(t, response.Title, "super")
	assert.Equal(t, response.Text, "cat")
	assert.Equal(t, response.AuthorId, int64(0))
	assert.False(t, response.Published)
}

func TestGRPC_ListAds(t *testing.T) {
	client, ctx := getClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "oleg@mail.ru"})
	assert.NoError(t, err, "client.ListAds")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{UserId: 0, Title: "hello", Text: "world"})
	assert.NoError(t, err)

	publishedAd, err := client.ChangeAdStatus(ctx, &grpcPort.ChangeAdStatusRequest{AdId: 0, UserId: 0, Published: true})
	assert.NoError(t, err)

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{UserId: 0, Title: "banana", Text: "3"})
	assert.NoError(t, err)

	ads, err := client.ListAds(ctx, &grpcPort.ListAdRequest{Filter: map[string]string{}})
	assert.NoError(t, err)
	assert.Len(t, ads.List, 1)
	assert.Equal(t, ads.List[0].Id, publishedAd.Id)
	assert.Equal(t, ads.List[0].Title, publishedAd.Title)
	assert.Equal(t, ads.List[0].Text, publishedAd.Text)
	assert.Equal(t, ads.List[0].AuthorId, publishedAd.AuthorId)
	assert.True(t, ads.List[0].Published)
}

func TestGRPCCreateUser(t *testing.T) {
	client, ctx := getClient(t)

	response, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "oleg@mail.ru"})
	assert.NoError(t, err, "client.CreateUser")
	assert.Equal(t, int64(0), response.Id)
	assert.Equal(t, "Oleg", response.Name)
	assert.Equal(t, "oleg@mail.ru", response.Email)
}

func TestGRPCGetUser(t *testing.T) {
	client, ctx := getClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "oleg@mail.ru"})
	assert.NoError(t, err, "client.GetUser")

	response, err := client.GetUser(ctx, &grpcPort.GetUserRequest{Id: 0})
	assert.NoError(t, err, "client.GetUser")
	assert.Equal(t, int64(0), response.Id)
	assert.Equal(t, "Oleg", response.Name)
	assert.Equal(t, "oleg@mail.ru", response.Email)
}

func TestGRPCDeleteUser(t *testing.T) {
	client, ctx := getClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "oleg@mail.ru"})
	assert.NoError(t, err, "client.DeleteUser")

	_, err = client.DeleteUser(ctx, &grpcPort.DeleteUserRequest{Id: 0})
	assert.NoError(t, err, "client.DeleteUser")
}

func TestGRPCDeleteAd(t *testing.T) {
	client, ctx := getClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "oleg@mail.ru"})
	assert.NoError(t, err, "client.DeleteAd")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{UserId: 0, Title: "banana", Text: "3"})
	assert.NoError(t, err)

	_, err = client.DeleteAd(ctx, &grpcPort.DeleteAdRequest{AdId: 0, AuthorId: 0})
	assert.NoError(t, err, "client.DeleteAd")
}
*/
