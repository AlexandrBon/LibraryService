package LibraryApp

import (
	"context"
	"fmt"
	"libraryService/internal/entity"
	"libraryService/internal/storage"
)

var ErrBadRequest = fmt.Errorf("bad request")

// App provides functions for getting authors/books by name/title
type App interface {
	GetBooksByAuthorName(ctx context.Context, name string) ([]entity.Book, error)
	GetAuthorsByBookTitle(ctx context.Context, title string) ([]entity.Author, error)
}

type LibraryApp struct {
	storage storage.IStorage
}

func NewLibraryApp(s storage.IStorage) App {
	return LibraryApp{storage: s}
}

func (a LibraryApp) GetBooksByAuthorName(ctx context.Context, name string) ([]entity.Book, error) {
	if !a.storage.CheckIfAuthorExists(ctx, name) {
		return nil, ErrBadRequest
	}

	books := a.storage.GetBooksByAuthorName(ctx, name)
	return books, nil
}

func (a LibraryApp) GetAuthorsByBookTitle(ctx context.Context, title string) ([]entity.Author, error) {
	if !a.storage.CheckIfBookExists(ctx, title) {
		return nil, ErrBadRequest
	}

	authors := a.storage.GetAuthorsByBookTitle(ctx, title)
	return authors, nil
}
