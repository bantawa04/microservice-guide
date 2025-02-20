package service

import (
	"github.com/bantawa04/go-mq/app/errors"
	"github.com/bantawa04/go-mq/app/model"
	"github.com/bantawa04/go-mq/app/repository"
	"gorm.io/gorm"
)

type BookService interface {
	CreateBook(book *model.BookModel) error
	WithTrx(tx *gorm.DB) BookService
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{bookRepo: bookRepo}
}

func (s *bookService) WithTrx(tx *gorm.DB) BookService {
	newService := &bookService{
		bookRepo: s.bookRepo.WithTrx(tx),
	}
	return newService
}

func (s *bookService) CreateBook(book *model.BookModel) error {
	if book == nil {
		return errors.NewBadRequestError("Product data cannot be empty")
	}

	err := s.bookRepo.CreateBook(book)
	if err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}
