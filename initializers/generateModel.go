package initializers

import "pagination/models"

func GenerateModel() {
	DB.AutoMigrate(&models.User{})
}
