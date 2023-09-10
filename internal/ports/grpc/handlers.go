package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	LibraryApp "libraryService/internal/libraryApp"
)

type LibraryService struct {
	libraryApp LibraryApp.App
	UnimplementedLibraryServiceServer
}

func NewServiceServer(la LibraryApp.App) LibraryServiceServer {
	return &LibraryService{libraryApp: la}
}

func getGRPCStatus(err error) codes.Code {
	if errors.Is(err, LibraryApp.ErrBadRequest) {
		return codes.NotFound
	}
	return codes.Unknown
}

func (ls *LibraryService) GetBooksByAuthorName(ctx context.Context, req *BooksByAuthorNameRequest) (*BooksResponse, error) {
	books, err := ls.libraryApp.GetBooksByAuthorName(ctx, req.AuthorName)
	if err != nil {
		return nil, status.Error(getGRPCStatus(err), err.Error())
	}
	response := BooksResponse{}
	for _, book := range books {
		bookResponse := &Book{Id: book.ID, Title: book.Title, PageCount: book.PageCount, PublishingYear: book.PublishingYear}
		response.Books = append(response.Books, bookResponse)
	}
	return &response, status.Error(codes.OK, "")
}

func (ls *LibraryService) GetAuthorsByBookTitle(ctx context.Context, req *AuthorsByBookTitleRequest) (*AuthorsResponse, error) {
	authors, err := ls.libraryApp.GetAuthorsByBookTitle(ctx, req.Title)
	if err != nil {
		return nil, status.Error(getGRPCStatus(err), err.Error())
	}
	response := AuthorsResponse{}
	for _, author := range authors {
		authorResponse := &Author{Id: author.ID, Name: author.Name, BookCount: author.BookCount, BirthDate: author.BirthDate}
		response.Authors = append(response.Authors, authorResponse)
	}
	return &response, status.Error(codes.OK, "")
}
