package cache

import (
	"context"
	"libraryService/internal/entity"
)

func NewStorage() (Storage, error) {

}

func (s Storage) GetBooksByAuthorName(ctx context.Context, name string) []entity.Book {

}

func (s Storage) CheckIfBookExists(ctx context.Context, title string) bool {

}

func (s Storage) GetAuthorsByBookTitle(ctx context.Context, title string) []entity.Author {

}

func (s Storage) CheckIfAuthorExists(ctx context.Context, name string) bool {

}
