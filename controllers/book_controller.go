package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    //"example.com/go-mongo-app/models"
    "example.com/go-mongo-app/services"
    //"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookController struct {
    Service services.BookServiceInterface
}

func NewBookController(service services.BookServiceInterface) *BookController {
    return &BookController{Service: service}
}

func (c *BookController) GetBookByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        http.Error(w, "ID no proporcionado", http.StatusBadRequest)
        return
    }

    book, err := c.Service.GetBookByID(id)
    if err != nil {
        http.Error(w, "Libro no encontrado", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Libro encontrado",
        "book":    book,
    })
}
