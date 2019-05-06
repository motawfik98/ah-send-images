package repositories

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func FindAllMainDepartments(db *gorm.DB) []models.MainDepartment {
	var mainDepartments []models.MainDepartment
	db.Find(&mainDepartments)
	return mainDepartments
}
