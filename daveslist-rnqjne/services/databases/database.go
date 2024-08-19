package databases

import (
	"Daveslist/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	DB, err = gorm.Open(sqlite.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		return err
	}
	DB.AutoMigrate(&models.CarList{})
	DB.AutoMigrate(&models.Categories{})
	DB.AutoMigrate(&models.File{})
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.Reply{})
	return nil
}
