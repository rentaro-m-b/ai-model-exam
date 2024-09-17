package request

import (
	"github.com/guregu/null"
)

type CreateBookRequest struct {
	Title     null.String `json:"title"`
	Author    null.String `json:"author"`
	Publisher null.String `json:"publisher"`
	Price     null.Int    `json:"price"`
}

func (rec *CreateBookRequest) Validate() (string, ValidationError) {
	if !rec.Title.Valid {
		return "title", ValidationErrRequestFieldMissing
	} else if rec.Title.String == "" {
		return "title", ValidationErrRequestFieldEmpty
	}

	if !rec.Author.Valid {
		return "author", ValidationErrRequestFieldMissing
	} else if rec.Author.String == "" {
		return "author", ValidationErrRequestFieldEmpty
	}

	if !rec.Publisher.Valid {
		return "publisher", ValidationErrRequestFieldMissing
	} else if rec.Publisher.String == "" {
		return "publisher", ValidationErrRequestFieldEmpty
	}

	if !rec.Price.Valid {
		return "price", ValidationErrRequestFieldMissing
	}

	return "", -1
}
