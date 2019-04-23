package main

import (
	"github.com/jinzhu/gorm"
)

type SecondaryDepartment struct {
	gorm.Model
	Hash             string
	Name             string `gorm:"not null"`
	Users            []User `gorm:"foreignkey:SecondaryDepartmentID"`
	MainDepartmentID int    `gorm:"index"`
}

func (secondaryDepartment *SecondaryDepartment) AfterCreate(scope *gorm.Scope) error {
	ID := int(secondaryDepartment.ID)
	hash := generateHash(ID)
	scope.DB().Model(secondaryDepartment).Update("Hash", hash)
	return nil
}
