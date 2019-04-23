package main

import (
	"github.com/jinzhu/gorm"
)

type MainDepartment struct {
	gorm.Model
	Hash                 string
	Name                 string                `gorm:"not null"`
	Users                []User                `gorm:"foreignkey:MainDepartmentID"`
	SecondaryDepartments []SecondaryDepartment `gorm:"foreignkey:MainDepartmentID"`
}

func (mainDepartment *MainDepartment) AfterCreate(scope *gorm.Scope) error {
	ID := int(mainDepartment.ID)
	hash := generateHash(ID)
	scope.DB().Model(mainDepartment).Update("Hash", hash)
	return nil
}
