package response

import (
	"github.com/rentaro-m-b/ai-model-exam/db"
)

type FetchBooksResponses struct {
	Books []FetchBooksResponse `json:"books"`
}

type FetchBooksResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Price     int    `json:"price"`
}

func ParseFetchBooksResponse(books []db.Book) *FetchBooksResponses {
	var res FetchBooksResponses
	for _, book := range books {
		res.Books = append(res.Books, FetchBooksResponse{
			ID:        int(book.ID),
			Title:     book.Title.String,
			Author:    book.Author.String,
			Publisher: book.Publisher.String,
			Price:     int(book.Price.Int32),
		})
	}

	return &res
}

type CreateBookErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

type FindBookByIdResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Price     int    `json:"price"`
}

func ParseFindBookByIdResponse(book *db.Book) *FindBookByIdResponse {
	return &FindBookByIdResponse{
		ID:        int(book.ID),
		Title:     book.Title.String,
		Author:    book.Author.String,
		Publisher: book.Publisher.String,
		Price:     int(book.Price.Int32),
	}
}
