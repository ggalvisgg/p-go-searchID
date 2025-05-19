package controllers_test

import (
    //"bytes"
    //"encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"

    "github.com/gorilla/mux"
    "example.com/go-mongo-app/controllers"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "example.com/go-mongo-app/models"
)

// ---------------------- MOCK DEL SERVICIO ----------------------


type MockBookService struct {
    mock.Mock
}

func (m *MockBookService) GetBookByID(id string) (*models.Book, error) {

    args := m.Called(id)
    return args.Get(0).(*models.Book), args.Error(1)
}

// ---------------------- TESTS ----------------------

func TestGetBookByID_Success(t *testing.T) {
    mockService := new(MockBookService)
    controller := controllers.NewBookController(mockService)

    id := primitive.NewObjectID().Hex()
    book := &models.Book{ID: primitive.NewObjectID(), Title: "Book Title", Author: "Author Name", ISBN: "1234567890"}

    mockService.On("GetBookByID", id).Return(book, nil)

    req := httptest.NewRequest("GET", "/books/"+id, nil)
    req = mux.SetURLVars(req, map[string]string{"id": id})
    resp := httptest.NewRecorder()

    controller.GetBookByID(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    mockService.AssertExpectations(t)
}

func TestGetBookByID_MissingID(t *testing.T) {
    mockService := new(MockBookService)
    controller := controllers.NewBookController(mockService)

    req := httptest.NewRequest("GET", "/books/", nil)
    req = mux.SetURLVars(req, map[string]string{"id": ""})
    resp := httptest.NewRecorder()

    controller.GetBookByID(resp, req)

    assert.Equal(t, http.StatusBadRequest, resp.Code)
    assert.Contains(t, resp.Body.String(), "ID no proporcionado")
}

func TestGetBookByID_NotFound(t *testing.T) {
    mockService := new(MockBookService)
    controller := controllers.NewBookController(mockService)

    id := "507f191e810c19729de860ea"

    mockService.On("GetBookByID", id).Return((*models.Book)(nil), fmt.Errorf("no encontrado"))

    req := httptest.NewRequest("GET", "/books/"+id, nil)
    req = mux.SetURLVars(req, map[string]string{"id": id})
    resp := httptest.NewRecorder()

    controller.GetBookByID(resp, req)

    assert.Equal(t, http.StatusNotFound, resp.Code)
    assert.Contains(t, resp.Body.String(), "Libro no encontrado")
}