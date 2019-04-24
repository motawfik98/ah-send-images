package main

import (
	"github.com/jinzhu/gorm"

)

type User struct {
	gorm.Model
	Username              string `gorm:"not null" form:"username"`
	Password              string `gorm:"not null" form:"password"`
	Hash                  string
	MainDepartmentID      int `gorm:"index" form:"main"`
	SecondaryDepartmentID int `gorm:"index" form:"secondary"`
}

func (user *User) AfterCreate(scope *gorm.Scope) error {
	ID := int(user.ID)
	hash := generateHash(ID)
	scope.DB().Model(user).Update("Hash", hash)
	return nil
}
