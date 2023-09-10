package cache

import (
	"context"
	lru "github.com/hashicorp/golang-lru/v2"
	"libraryService/internal/entity"
	"libraryService/internal/storage"
	"log"
)

type Storage struct {
	mainStorage storage.IStorage
	cache       *lru.Cache[string, any]
}

func NewStorage(s storage.IStorage) (storage.IStorage, error) {
	cache, err := lru.New[string, any](128)
	if err != nil {
		log.Fatal(err)
	}
	return Storage{mainStorage: s, cache: cache}, nil
}

func (s Storage) GetBooksByAuthorName(ctx context.Context, name string) []entity.Book {
	booksAny, ok := s.cache.Get(name)

	if !ok {
		books := s.mainStorage.GetBooksByAuthorName(ctx, name)
		s.cache.Add(name, books)
		return books
	}

	books, ok := booksAny.([]entity.Book)
	if !ok {
		return nil
	}
	return books
}

func (s Storage) CheckIfBookExists(ctx context.Context, title string) bool {
	ok := s.cache.Contains(title)
	if !ok {
		return s.mainStorage.CheckIfBookExists(ctx, title)
	}
	return true
}

func (s Storage) GetAuthorsByBookTitle(ctx context.Context, title string) []entity.Author {
	authorsAny, ok := s.cache.Get(title)

	if !ok {
		authors := s.mainStorage.GetAuthorsByBookTitle(ctx, title)
		s.cache.Add(title, authors)
		return authors
	}

	authors, ok := authorsAny.([]entity.Author)
	if !ok {
		return nil
	}
	return authors
}

func (s Storage) CheckIfAuthorExists(ctx context.Context, name string) bool {
	ok := s.cache.Contains(name)
	if !ok {
		return s.mainStorage.CheckIfAuthorExists(ctx, name)
	}
	return true
}
