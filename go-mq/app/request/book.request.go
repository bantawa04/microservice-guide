package request

import (
	"github.com/bantawa04/go-mq/app/dao"
	"github.com/bantawa04/go-mq/app/model"
)

type CreateBookRequestData struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required"`
}

// ToModel converts the request data to a UserModel
func (r *CreateBookRequestData) ToModel() *model.BookModel {
	return &model.BookModel{
		Book: dao.Book{
			Name:        r.Name,
			Description: &r.Description,
			Price:       r.Price,
		},
	}
}
