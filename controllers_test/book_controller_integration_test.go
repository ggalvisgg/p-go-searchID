package controllers_test

import (
	//"bytes"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-mongo-app/models"
	"example.com/go-mongo-app/controllers"
	"example.com/go-mongo-app/repositories"
	"example.com/go-mongo-app/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupRouter() *mux.Router {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)
	controller := controllers.NewBookController(service)

	r := mux.NewRouter()
	r.HandleFunc("/books/{id}", controller.GetBookByID).Methods("GET")

	return r
}

func TestBookCRUDIntegration(t *testing.T) {
	
	router := setupRouter()

	createdBook := &models.Book {
    ID:     primitive.NewObjectID(),
    Title:  "Test Book",
    Author: "Author Test",
    ISBN:   "1234567890",
    }

	// 3. Get book by ID
	urlByID := fmt.Sprintf("/books/%s", createdBook.ID.Hex())
	reqGet, _ := http.NewRequest("GET", urlByID, nil)
	rrGet := httptest.NewRecorder()
	router.ServeHTTP(rrGet, reqGet)
	assert.Equal(t, http.StatusOK, rrGet.Code)
}
