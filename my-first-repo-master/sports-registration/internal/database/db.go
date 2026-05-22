package database

import (
	"fmt"
	"log"
	"os"
	"sports-registration/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // Каскадное удаление запрещено
	})
	if err != nil {
		log.Fatal("❌ Ошибка подключения к БД:", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Service{}, &models.Application{}, &models.ApplicationService{})
	if err != nil {
		log.Fatal("❌ Ошибка миграции:", err)
	}
	log.Println("✅ БД подключена и мигрирована")
}
