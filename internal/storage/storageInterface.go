package storage

import (
	"context"
	"libraryService/internal/entity"
)

type IBook interface {
	GetBooksByAuthorName(ctx context.Context, name string) []entity.Book
	CheckIfBookExists(ctx context.Context, title string) bool
}

type IAuthor interface {
	GetAuthorsByBookTitle(ctx context.Context, title string) []entity.Author
	CheckIfAuthorExists(ctx context.Context, name string) bool
}

type IStorage interface {
	IBook
	IAuthor
}
