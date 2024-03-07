package initializers

import "auth/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {

	}
}
