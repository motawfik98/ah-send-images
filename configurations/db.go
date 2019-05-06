package configurations

import (
	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mssql", "sqlserver://remote:mohamed@localhost:1433?database=ah_images_go")
	if err == nil {
		db.AutoMigrate(&models.MainDepartment{}, &models.SecondaryDepartment{}, &models.User{})
	}
	return db, err
}
