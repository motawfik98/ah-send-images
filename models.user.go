package main

import (
	"github.com/jinzhu/gorm"

)

type User struct {
	gorm.Model
	Username              string `gorm:"not null" json:"username"`
	Password              string `gorm:"not null" json:"password"`
	Hash                  string
	MainDepartmentID      int `gorm:"index" json:"main"`
	SecondaryDepartmentID int `gorm:"index" json:"secondary"`
}

func (user *User) AfterCreate(scope *gorm.Scope) error {
	ID := int(user.ID)
	hash := generateHash(ID)
	scope.DB().Model(user).Update("Hash", hash)
	return nil
}
