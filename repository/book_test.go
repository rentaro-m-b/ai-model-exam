package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/rentaro-m-b/ai-model-exam/db"
	"github.com/rentaro-m-b/ai-model-exam/repository"
	"github.com/stretchr/testify/assert"
)

func TestListBooks(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("the error '%s' when opening a stub database connection", err)
	}
	defer mock.Close()

	columns := []string{
		"id",
		"title",
		"author",
		"publisher",
		"price",
	}
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

	rows := pgxmock.NewRows(columns)
	for _, expect := range expects {
		rows.AddRow(
			expect.ID,
			expect.Title,
			expect.Author,
			expect.Publisher,
			expect.Price,
		)
	}
	sql := `
	-- name: ListBooks :many
	SELECT id, title, author, publisher, price
    FROM books
	`
	mock.ExpectQuery(sql).
		WillReturnRows(rows)

	repo := repository.NewBookRepository(db.New(mock))
	books, err := repo.ListBooks(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expects[0], books[0])
	assert.Equal(t, expects[1], books[1])

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("didn't execute query: %v", err)
	}
}

func TestListBooksFailure(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("the error '%s' when opening a stub database connection", err)
	}
	defer mock.Close()

	sql := `
	-- name: ListBooks :many
	SELECT id, title, author, publisher, price
    FROM books
	`
	mock.ExpectQuery(sql).
		WillReturnError(fmt.Errorf("query error"))

	repo := repository.NewBookRepository(db.New(mock))
	books, err := repo.ListBooks(context.Background())
	assert.Error(t, err)
	assert.Nil(t, books)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("didn't execute query: %v", err)
	}
}

func TestCreateBook(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("the error '%s' when opening a stub database connection", err)
	}
	defer mock.Close()

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

	columns := []string{
		"id",
		"title",
		"author",
		"publisher",
		"price",
	}
	rows := pgxmock.NewRows(columns).AddRow(expect.ID, expect.Title, expect.Author, expect.Publisher, expect.Price)

	sql := `
	-- name: CreateBook :one
	INSERT INTO books \(id, title, author, publisher, price\)
    VALUES \(nextval\('BOOK_ID_SEQ'\), \$1, \$2, \$3, \$4\)
    RETURNING id, title, author, publisher, price
	`
	mock.ExpectQuery(sql).
		WithArgs(param.Title, param.Author, param.Publisher, param.Price).
		WillReturnRows(rows)

	repo := repository.NewBookRepository(db.New(mock))
	book, err := repo.CreateBook(context.Background(), &param)
	assert.NoError(t, err)
	assert.Equal(t, &expect, book)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("didn't execute query: %v", err)
	}
}

func TestCreateBookFailure(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("the error '%s' when opening a stub database connection", err)
	}
	defer mock.Close()

	param := db.CreateBookParams{
		Title:     pgtype.Text{String: "test title 1", Valid: true},
		Author:    pgtype.Text{String: "test author 1", Valid: true},
		Publisher: pgtype.Text{String: "test publisher 1", Valid: true},
		Price:     pgtype.Int4{Int32: 200, Valid: true},
	}

	sql := `
	-- name: CreateBook :one
	INSERT INTO books \(id, title, author, publisher, price\)
    VALUES \(nextval\('BOOK_ID_SEQ'\), \$1, \$2, \$3, \$4\)
    RETURNING id, title, author, publisher, price
	`
	mock.ExpectQuery(sql).
		WithArgs(param.Title, param.Author, param.Publisher, param.Price).
		WillReturnError(fmt.Errorf("query error"))

	repo := repository.NewBookRepository(db.New(mock))
	book, err := repo.CreateBook(context.Background(), &param)
	assert.Error(t, err)
	assert.Nil(t, book)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("didn't execute query: %v", err)
	}
}

func TestGetBookById(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("the error '%s' when opening a stub database connection", err)
	}
	defer mock.Close()

	id := 1
	expect := db.Book{
		ID:        1,
		Title:     pgtype.Text{String: "test title 1", Valid: true},
		Author:    pgtype.Text{String: "test author 1", Valid: true},
		Publisher: pgtype.Text{String: "test publisher 1", Valid: true},
		Price:     pgtype.Int4{Int32: 200, Valid: true},
	}

	columns := []string{
		"id",
		"title",
		"author",
		"publisher",
		"price",
	}
	rows := pgxmock.NewRows(columns).AddRow(expect.ID, expect.Title, expect.Author, expect.Publisher, expect.Price)

	sql := `-- name: GetBookByID :one
	SELECT id, title, author, publisher, price
		FROM books
		WHERE id = \$1
	`
	mock.ExpectQuery(sql).
		WithArgs(int32(id)).
		WillReturnRows(rows)

	repo := repository.NewBookRepository(db.New(mock))
	book, err := repo.GetBookById(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, &expect, book)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("didn't execute query: %v", err)
	}
}

func TestGetBookByIdFailure(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("the error '%s' when opening a stub database connection", err)
	}
	defer mock.Close()

	id := 1

	sql := `-- name: GetBookByID :one
	SELECT id, title, author, publisher, price
		FROM books
		WHERE id = \$1
	`
	mock.ExpectQuery(sql).
		WithArgs(int32(id)).
		WillReturnError(fmt.Errorf("query error"))
	repo := repository.NewBookRepository(db.New(mock))
	book, err := repo.GetBookById(context.Background(), id)
	assert.Error(t, err)
	assert.Nil(t, book)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("didn't execute query: %v", err)
	}
}
