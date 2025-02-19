package model

import (
	"github.com/bantawa04/go-mq/app/dao"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	dao.User
}

func (u *UserModel) BeforeCreate(db *gorm.DB) error {

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(password)
	if err != nil {
		return err
	}
	return nil
}
