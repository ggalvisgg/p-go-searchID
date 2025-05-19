package services

import (
    //"fmt"
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/repositories"
)

type BookServiceInterface interface {
    GetBookByID(id string) (*models.Book, error)
}

type BookService struct {
    repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
    return &BookService{repo}
}

func (s *BookService) GetBookByID(id string) (*models.Book, error) {
    return s.repo.GetBookByID(id)
}
