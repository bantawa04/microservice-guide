package repository

import (
	"github.com/bantawa04/go-mq/app/model"
	"github.com/bantawa04/go-mq/config"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(product *model.BookModel) error
	WithTrx(tx *gorm.DB) BookRepository
}

func (r *bookRepository) WithTrx(tx *gorm.DB) BookRepository {
	if tx == nil {
		return r
	}
	newRepo := &bookRepository{
		db: tx,
	}
	return newRepo
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository() BookRepository {
	return &bookRepository{
		db: config.DB.Db,
	}
}

func (r *bookRepository) CreateBook(product *model.BookModel) error {
	err := r.db.Create(product).Error
	if err != nil {
		return err
	}
	return nil
}
