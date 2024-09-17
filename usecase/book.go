package usecase

import (
	"context"
	"log"

	"github.com/rentaro-m-b/ai-model-exam/db"
	"github.com/rentaro-m-b/ai-model-exam/repository"
)

type BookUsecase interface {
	FetchBooks(ctx context.Context) ([]db.Book, error)
	CreateBook(ctx context.Context, param *db.CreateBookParams) (*db.Book, error)
	FindBookById(ctx context.Context, id int) (*db.Book, error)
}

type bookUsecaseImpl struct {
	repository repository.BookRepository
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return &bookUsecaseImpl{
		repository: repository,
	}
}

func (u *bookUsecaseImpl) FetchBooks(ctx context.Context) ([]db.Book, error) {
	books, err := u.repository.ListBooks(ctx)
	if err != nil {
		log.Printf("Unable to execute BookUsecaseFetchBooks: %d\n", err)
		return nil, err
	}

	return books, nil
}

func (u *bookUsecaseImpl) CreateBook(ctx context.Context, param *db.CreateBookParams) (*db.Book, error) {
	book, err := u.repository.CreateBook(ctx, param)
	if err != nil {
		log.Printf("Unable to execute BookUsecaseCreateBook: %d\n", err)
		return nil, err
	}

	return book, nil
}

func (u *bookUsecaseImpl) FindBookById(ctx context.Context, id int) (*db.Book, error) {
	book, err := u.repository.GetBookById(ctx, id)
	if err != nil {
		log.Printf("Unable to execute BookUsecaseFindBookById: %d\n", err)
		return nil, err
	}

	return book, nil
}
