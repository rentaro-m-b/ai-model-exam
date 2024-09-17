package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rentaro-m-b/ai-model-exam/db"
	"github.com/rentaro-m-b/ai-model-exam/handler"
	"github.com/rentaro-m-b/ai-model-exam/repository"
	"github.com/rentaro-m-b/ai-model-exam/usecase"
)

func Init(e *echo.Echo, db *db.Queries) {
	bookRepository := repository.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	bookHandler := handler.NewBookHandler(bookUsecase)

	e.GET("/books", bookHandler.FetchBooks)
	e.POST("/books", bookHandler.CreateBook)
	e.GET("/books/:id", bookHandler.FindBookById)
}
