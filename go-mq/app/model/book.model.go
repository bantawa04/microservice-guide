package model

import (
	"github.com/bantawa04/go-mq/app/dao"
)

type BookModel struct {
	dao.Book
}

// func (u *BookModel) BeforeCreate(db *gorm.DB) error {

// 	u.UserID = "a1e001be-55e9-415d-8b52-31f28a574a1b"

// 	return nil
// }
