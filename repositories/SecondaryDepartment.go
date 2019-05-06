package repositories

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func FindAllSecondaryDepartments(db *gorm.DB) []models.SecondaryDepartment {
	var secondaryDepartments []models.SecondaryDepartment
	db.Find(&secondaryDepartments)
	return secondaryDepartments
}
