package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/guregu/null"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/rentaro-m-b/ai-model-exam/db"
	"github.com/rentaro-m-b/ai-model-exam/handler"
	"github.com/rentaro-m-b/ai-model-exam/handler/request"
	"github.com/rentaro-m-b/ai-model-exam/handler/response"
	mock_usecase "github.com/rentaro-m-b/ai-model-exam/usecase/mock"
	"github.com/stretchr/testify/assert"
)

func TestFetchBooks(t *testing.T) {
	// モックコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ユースケースのモックを作成し、期待値を設定
	mockUc := mock_usecase.NewMockBookUsecase(ctrl)
	expectsUc := []db.Book{
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
	mockUc.EXPECT().FetchBooks(gomock.Any()).Return(expectsUc, nil)

	// Echoのインスタンス、リクエスト、レスポンスを作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラを作成し、テスト項目を検証
	h := handler.NewBookHandler(mockUc)
	expects := response.ParseFetchBooksResponse(expectsUc)
	assert.NoError(t, h.FetchBooks(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	var res *response.FetchBooksResponses
	err := json.NewDecoder(rec.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, expects, res)
}

func TestFetchBooksFailure(t *testing.T) {
	// モックコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ユースケースのモックを作成し、期待値を設定
	mockUc := mock_usecase.NewMockBookUsecase(ctrl)
	mockUc.EXPECT().FetchBooks(gomock.Any()).Return(nil, fmt.Errorf("error"))

	// Echoのインスタンス、リクエスト、レスポンスを作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラを作成し、テスト項目を検証
	h := handler.NewBookHandler(mockUc)
	assert.NoError(t, h.FetchBooks(c))
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	expectErrorMessage := `{"message": "Internal server error"}`
	assert.JSONEq(t, expectErrorMessage, rec.Body.String())
}

func TestCreateBook(t *testing.T) {
	// モックコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ユースケースのモックを作成し、期待値を設定
	mockUc := mock_usecase.NewMockBookUsecase(ctrl)
	paramUc := db.CreateBookParams{
		Title:     pgtype.Text{String: "test title 1", Valid: true},
		Author:    pgtype.Text{String: "test author 1", Valid: true},
		Publisher: pgtype.Text{String: "test publisher 1", Valid: true},
		Price:     pgtype.Int4{Int32: 100, Valid: true},
	}
	expectUc := db.Book{
		ID:        1,
		Title:     pgtype.Text{String: "test title 1", Valid: true},
		Author:    pgtype.Text{String: "test author 1", Valid: true},
		Publisher: pgtype.Text{String: "test publisher 1", Valid: true},
		Price:     pgtype.Int4{Int32: 100, Valid: true},
	}
	mockUc.EXPECT().CreateBook(gomock.Any(), &paramUc).Return(&expectUc, nil)

	// リクエストボディを設定
	param := request.CreateBookRequest{
		Title:     null.NewString("test title 1", true),
		Author:    null.NewString("test author 1", true),
		Publisher: null.NewString("test publisher 1", true),
		Price:     null.NewInt(100, true),
	}

	// Echoのインスタンス、リクエスト、レスポンスを作成
	e := echo.New()
	reqBody, _ := json.Marshal(param)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラを作成し、テスト項目を検証
	h := handler.NewBookHandler(mockUc)
	assert.NoError(t, h.CreateBook(c))
	assert.Equal(t, http.StatusCreated, rec.Code)
	expectLocation := "http://example.com/books/1"
	assert.Equal(t, expectLocation, rec.Header().Get("Location"))
}

func TestCreateBookFailure(t *testing.T) {
	// モックコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ユースケースのモックを作成し、期待値を設定
	mockUc := mock_usecase.NewMockBookUsecase(ctrl)
	mockUc.EXPECT().FetchBooks(gomock.Any()).Return(nil, fmt.Errorf("error"))

	// Echoのインスタンス、リクエスト、レスポンスを作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラを作成し、テスト項目を検証
	h := handler.NewBookHandler(mockUc)
	assert.NoError(t, h.FetchBooks(c))
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	expectErrorMessage := `{"message": "Internal server error"}`
	assert.JSONEq(t, expectErrorMessage, rec.Body.String())
}

func TestCreateBookFailureValidationNone(t *testing.T) {
	// モックコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ユースケースのモックを作成
	mockUc := mock_usecase.NewMockBookUsecase(ctrl)

	// リクエストボディを設定
	param := request.CreateBookRequest{
		Title:     null.NewString("", false),
		Author:    null.NewString("test author 1", true),
		Publisher: null.NewString("test publisher 1", true),
		Price:     null.NewInt(100, true),
	}

	// Echoのインスタンス、リクエスト、レスポンスを作成
	e := echo.New()
	reqBody, _ := json.Marshal(param)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラを作成し、テスト項目を検証
	h := handler.NewBookHandler(mockUc)
	assert.NoError(t, h.CreateBook(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var res response.CreateBookErrorResponse
	err := json.NewDecoder(rec.Body).Decode(&res)
	assert.NoError(t, err)
	expect := response.CreateBookErrorResponse{
		Type:     "about:none",
		Title:    "request validation error is occurred.",
		Detail:   "title must not be none.",
		Instance: "/books",
	}
	fmt.Println(res)
	assert.Equal(t, expect, res)
}

func TestCreateBookFailureValidationEmpty(t *testing.T) {
	// モックコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ユースケースのモックを作成
	mockUc := mock_usecase.NewMockBookUsecase(ctrl)

	// リクエストボディを設定
	param := request.CreateBookRequest{
		Title:     null.NewString("", true),
		Author:    null.NewString("test author 1", true),
		Publisher: null.NewString("test publisher 1", true),
		Price:     null.NewInt(100, true),
	}

	// Echoのインスタンス、リクエスト、レスポンスを作成
	e := echo.New()
	reqBody, _ := json.Marshal(param)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラを作成し、テスト項目を検証
	h := handler.NewBookHandler(mockUc)
	assert.NoError(t, h.CreateBook(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var res response.CreateBookErrorResponse
	err := json.NewDecoder(rec.Body).Decode(&res)
	assert.NoError(t, err)
	expect := response.CreateBookErrorResponse{
		Type:     "about:blank",
		Title:    "request validation error is occurred.",
		Detail:   "title must not be blank.",
		Instance: "/books",
	}
	assert.Equal(t, expect, res)
}

func TestFindBookById(t *testing.T) {
	// モックコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ユースケースのモックを作成し、期待値を設定
	mockUc := mock_usecase.NewMockBookUsecase(ctrl)
	idUc := 1
	expectUc := db.Book{
		ID:        1,
		Title:     pgtype.Text{String: "test title 1", Valid: true},
		Author:    pgtype.Text{String: "test author 1", Valid: true},
		Publisher: pgtype.Text{String: "test publisher 1", Valid: true},
		Price:     pgtype.Int4{Int32: 200, Valid: true},
	}
	mockUc.EXPECT().FindBookById(gomock.Any(), idUc).Return(&expectUc, nil)

	// パスパラメータを設定
	id := 1

	// Echoのインスタンス、リクエスト、レスポンスを作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))

	// ハンドラを作成し、テスト項目を検証
	h := handler.NewBookHandler(mockUc)
	expect := response.ParseFindBookByIdResponse(&expectUc)
	assert.NoError(t, h.FindBookById(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	var res *response.FindBookByIdResponse
	err := json.NewDecoder(rec.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, expect, res)
}

func TestFindBookByIdFailure(t *testing.T) {
	// モックコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	idUc := 1

	// ユースケースのモックを作成し、期待値を設定
	mockUc := mock_usecase.NewMockBookUsecase(ctrl)
	mockUc.EXPECT().FindBookById(gomock.Any(), idUc).Return(nil, fmt.Errorf("error"))

	// パスパラメータを設定
	id := 1

	// Echoのインスタンス、リクエスト、レスポンスを作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))

	// ハンドラを作成し、テスト項目を検証
	h := handler.NewBookHandler(mockUc)
	assert.NoError(t, h.FindBookById(c))
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	expectErrorMessage := `{"message": "Internal server error"}`
	assert.JSONEq(t, expectErrorMessage, rec.Body.String())
}
