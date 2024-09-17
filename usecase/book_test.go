package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rentaro-m-b/ai-model-exam/db"
	mock_repository "github.com/rentaro-m-b/ai-model-exam/repository/mock"
	"github.com/rentaro-m-b/ai-model-exam/usecase"
	"github.com/stretchr/testify/assert"
)

func TestFetchBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockBookRepository(ctrl)
	expects := []db.Book{
		{
			ID:        1,
			Title:     pgtype.Text{String: "test title 1", Valid: true},
			Author:    pgtype.Text{String: "test author 1", Valid: true},
			Publisher: pgtype.Text{String: "test publisher 1", Valid: true},
			Price:     pgtype.Int4{Int32: 100, Valid: true},
		},
		{
			ID:        2,
			Title:     pgtype.Text{String: "test title 2", Valid: true},
			Author:    pgtype.Text{String: "test author 2", Valid: true},
			Publisher: pgtype.Text{String: "test publisher 2", Valid: true},
			Price:     pgtype.Int4{Int32: 200, Valid: true},
		},
	}
	mockRepo.EXPECT().ListBooks(gomock.Any()).Return(expects, nil)

	uc := usecase.NewBookUsecase(mockRepo)
	books, err := uc.FetchBooks(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expects, books)
}

func TestListBooksFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockBookRepository(ctrl)
	uc := usecase.NewBookUsecase(mockRepo)

	mockRepo.EXPECT().ListBooks(gomock.Any()).Return(nil, errors.New("error"))

	books, err := uc.FetchBooks(context.Background())
	assert.Error(t, err)
	assert.Nil(t, books)
}

func TestCreateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockBookRepository(ctrl)
	uc := usecase.NewBookUsecase(mockRepo)

	param := db.CreateBookParams{
		Title:     pgtype.Text{String: "test title 1", Valid: true},
		Author:    pgtype.Text{String: "test author 1", Valid: true},
		Publisher: pgtype.Text{String: "test publisher 1", Valid: true},
		Price:     pgtype.Int4{Int32: 200, Valid: true},
	}
	expect := db.Book{
		ID:        1,
		Title:     pgtype.Text{String: "test title 1", Valid: true},
		Author:    pgtype.Text{String: "test author 1", Valid: true},
		Publisher: pgtype.Text{String: "test publisher 1", Valid: true},
		Price:     pgtype.Int4{Int32: 200, Valid: true},
	}

	mockRepo.EXPECT().CreateBook(gomock.Any(), &param).Return(&expect, nil)

	book, err := uc.CreateBook(context.Background(), &param)
	assert.NoError(t, err)
	assert.Equal(t, &expect, book)
}

func TestCreateBookFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockBookRepository(ctrl)
	uc := usecase.NewBookUsecase(mockRepo)

	param := db.CreateBookParams{
		Title:     pgtype.Text{String: "test title 1", Valid: true},
		Author:    pgtype.Text{String: "test author 1", Valid: true},
		Publisher: pgtype.Text{String: "test publisher 1", Valid: true},
		Price:     pgtype.Int4{Int32: 200, Valid: true},
	}

	mockRepo.EXPECT().CreateBook(gomock.Any(), &param).Return(nil, errors.New("error"))

	book, err := uc.CreateBook(context.Background(), &param)
	assert.Error(t, err)
	assert.Nil(t, book)
}

func TestFindBookById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockBookRepository(ctrl)
	uc := usecase.NewBookUsecase(mockRepo)

	id := 1
	expect := db.Book{
		ID:        1,
		Title:     pgtype.Text{String: "test title 1", Valid: true},
		Author:    pgtype.Text{String: "test author 1", Valid: true},
		Publisher: pgtype.Text{String: "test publisher 1", Valid: true},
		Price:     pgtype.Int4{Int32: 200, Valid: true},
	}

	mockRepo.EXPECT().GetBookById(gomock.Any(), id).Return(&expect, nil)

	book, err := uc.FindBookById(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, &expect, book)
}

func TestFindBookByIdFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockBookRepository(ctrl)
	uc := usecase.NewBookUsecase(mockRepo)

	id := 1

	mockRepo.EXPECT().GetBookById(gomock.Any(), id).Return(nil, errors.New("error"))

	book, err := uc.FindBookById(context.Background(), id)
	assert.Error(t, err)
	assert.Nil(t, book)
}
