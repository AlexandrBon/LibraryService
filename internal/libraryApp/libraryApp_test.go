package LibraryApp

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"libraryService/internal/entity"
	"libraryService/internal/storage/mocks"
	"testing"
)

func TestLibraryApp_GetAuthorsByBookTitle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockStorage := mocks.NewMockIStorage(mockCtrl)
	testLibraryApp := NewLibraryApp(mockStorage)
	ctx := context.Background()

	testAuthor := entity.Author{ID: 1, Name: "a_1", BirthDate: "1908-11-23", BookCount: 32}
	mockStorage.EXPECT().GetAuthorsByBookTitle(ctx, "t_1").Return([]entity.Author{testAuthor}).AnyTimes()
	mockStorage.EXPECT().CheckIfBookExists(ctx, "t_1").Return(true).AnyTimes()

	authors, err := testLibraryApp.GetAuthorsByBookTitle(ctx, "t_1")
	assert.NoError(t, err)

	assert.Len(t, authors, 1)
	author := authors[0]
	assert.Equal(t, author.ID, int32(1))
	assert.Equal(t, author.Name, "a_1")
	assert.Equal(t, author.BookCount, int32(32))
	assert.Equal(t, author.BirthDate, "1908-11-23")
}

func TestLibraryApp_GetBooksByAuthorName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockStorage := mocks.NewMockIStorage(mockCtrl)
	testLibraryApp := NewLibraryApp(mockStorage)
	ctx := context.Background()

	mockStorage.EXPECT().CheckIfAuthorExists(ctx, "a_1").Return(false).AnyTimes()

	_, err := testLibraryApp.GetBooksByAuthorName(ctx, "a_1")
	assert.Error(t, err, ErrBadRequest)
}
