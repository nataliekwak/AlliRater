package database

import (
	"go-admin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("mydb_admin:rootroot@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
