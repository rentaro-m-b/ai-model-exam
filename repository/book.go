package repository

import (
	"context"
	"log"

	"github.com/rentaro-m-b/ai-model-exam/db"
)

type BookRepository interface {
	ListBooks(ctx context.Context) ([]db.Book, error)
	CreateBook(ctx context.Context, param *db.CreateBookParams) (*db.Book, error)
	GetBookById(ctx context.Context, id int) (*db.Book, error)
}

type bookRepositoryImpl struct {
	queries *db.Queries
}

func NewBookRepository(db *db.Queries) BookRepository {
	return &bookRepositoryImpl{
		queries: db,
	}
}

func (r *bookRepositoryImpl) ListBooks(ctx context.Context) ([]db.Book, error) {
	books, err := r.queries.ListBooks(ctx)
	if err != nil {
		log.Printf("Unable to execute BookRepositoryListBooks: %d\n", err)
		return nil, err
	}

	return books, nil
}

func (r *bookRepositoryImpl) CreateBook(ctx context.Context, param *db.CreateBookParams) (*db.Book, error) {
	book, err := r.queries.CreateBook(ctx, *param)
	if err != nil {
		log.Printf("Unable to execute BookRepositoryCreateBook: %d\n", err)
		return nil, err
	}

	return &book, nil
}

func (r *bookRepositoryImpl) GetBookById(ctx context.Context, id int) (*db.Book, error) {
	book, err := r.queries.GetBookByID(ctx, int32(id))
	if err != nil {
		log.Printf("Unable to execute BookRepositoryGetBookById: %d\n", err)
		return nil, err
	}

	return &book, nil
}
