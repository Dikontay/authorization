package pkg

import (
	"auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

}

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{}, &models.Role{})
	if err != nil {
		log.Printf("Migration db error %e", err)
		//os.Exit(1)
	}
}
