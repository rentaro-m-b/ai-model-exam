package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/rentaro-m-b/ai-model-exam/db"
	"github.com/rentaro-m-b/ai-model-exam/handler/request"
	"github.com/rentaro-m-b/ai-model-exam/handler/response"
	"github.com/rentaro-m-b/ai-model-exam/usecase"
)

type BookHandler interface {
	FetchBooks(c echo.Context) error
	CreateBook(c echo.Context) error
	FindBookById(c echo.Context) error
}

type bookHandlerImpl struct {
	usecase usecase.BookUsecase
}

func NewBookHandler(usecase usecase.BookUsecase) BookHandler {
	return &bookHandlerImpl{
		usecase: usecase,
	}
}

func (h *bookHandlerImpl) FetchBooks(c echo.Context) error {
	books, err := h.usecase.FetchBooks(context.Background())
	if err != nil {
		log.Printf("Unable to execute BookHandlerFetchBooks: %d\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, response.ParseFetchBooksResponse(books))
}

// メモ：レスポンス値に改修の余地あり
func (h *bookHandlerImpl) CreateBook(c echo.Context) error {
	body := new(request.CreateBookRequest)
	if err := c.Bind(body); err != nil {
		log.Printf("Unable to execute BookHandlerCreateBook: %d\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Internal server error",
		})
	}
	vs, ve := body.Validate()
	if ve != -1 {
		var errRes response.CreateBookErrorResponse
		if ve == request.ValidationErrRequestFieldMissing {
			errRes.Type = "about:none"
			errRes.Detail = fmt.Sprintf("%s must not be none.", vs)
		} else if ve == request.ValidationErrRequestFieldEmpty {
			errRes.Type = "about:blank"
			errRes.Detail = fmt.Sprintf("%s must not be blank.", vs)
		}
		errRes.Title = "request validation error is occurred."
		errRes.Instance = "/books"

		return c.JSON(http.StatusBadRequest, errRes)
	}

	param := db.CreateBookParams{
		Title:     pgtype.Text{String: body.Title.String, Valid: true},
		Author:    pgtype.Text{String: body.Author.String, Valid: true},
		Publisher: pgtype.Text{String: body.Publisher.String, Valid: true},
		Price:     pgtype.Int4{Int32: int32(body.Price.Int64), Valid: true},
	}

	book, err := h.usecase.CreateBook(context.Background(), &param)
	if err != nil {
		log.Printf("Unable to execute BookHandlerCreateBook: %d\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Internal server error",
		})
	}
	location := fmt.Sprintf("%s/books/%d", c.Scheme()+"://"+c.Request().Host, book.ID)
	c.Response().Header().Set("Location", location)

	return c.JSON(http.StatusCreated, nil)
}

func (h *bookHandlerImpl) FindBookById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Unable to execute BookHandlerFindBookById: %d\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid book ID",
		})
	}

	book, err := h.usecase.FindBookById(context.Background(), id)
	if err != nil {
		log.Printf("Unable to execute BookHandlerFindBookById: %d\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, response.ParseFindBookByIdResponse(book))
}
