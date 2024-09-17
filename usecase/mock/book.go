// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/book.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/rentaro-m-b/ai-model-exam/db"
)

// MockBookUsecase is a mock of BookUsecase interface.
type MockBookUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockBookUsecaseMockRecorder
}

// MockBookUsecaseMockRecorder is the mock recorder for MockBookUsecase.
type MockBookUsecaseMockRecorder struct {
	mock *MockBookUsecase
}

// NewMockBookUsecase creates a new mock instance.
func NewMockBookUsecase(ctrl *gomock.Controller) *MockBookUsecase {
	mock := &MockBookUsecase{ctrl: ctrl}
	mock.recorder = &MockBookUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookUsecase) EXPECT() *MockBookUsecaseMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookUsecase) CreateBook(ctx context.Context, param *db.CreateBookParams) (*db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", ctx, param)
	ret0, _ := ret[0].(*db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookUsecaseMockRecorder) CreateBook(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookUsecase)(nil).CreateBook), ctx, param)
}

// FetchBooks mocks base method.
func (m *MockBookUsecase) FetchBooks(ctx context.Context) ([]db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBooks", ctx)
	ret0, _ := ret[0].([]db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBooks indicates an expected call of FetchBooks.
func (mr *MockBookUsecaseMockRecorder) FetchBooks(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBooks", reflect.TypeOf((*MockBookUsecase)(nil).FetchBooks), ctx)
}

// FindBookById mocks base method.
func (m *MockBookUsecase) FindBookById(ctx context.Context, id int) (*db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBookById", ctx, id)
	ret0, _ := ret[0].(*db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBookById indicates an expected call of FindBookById.
func (mr *MockBookUsecaseMockRecorder) FindBookById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBookById", reflect.TypeOf((*MockBookUsecase)(nil).FindBookById), ctx, id)
}