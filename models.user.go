package main

import (
	"github.com/jinzhu/gorm"

)

type User struct {
	gorm.Model
	Username              string `gorm:"not null"`
	Password              string `gorm:"not null"`
	Hash                  string
	MainDepartmentID      int `gorm:"index"`
	SecondaryDepartmentID int `gorm:"index"`
}

func (user *User) AfterCreate(scope *gorm.Scope) error {
	ID := int(user.ID)
	hash := generateHash(ID)
	scope.DB().Model(user).Update("Hash", hash)
	return nil
}
