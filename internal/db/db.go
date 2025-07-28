package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"messenger-api/internal/models"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Ошибка при подключении к базе данных: ", err)
	}

	// Автоматическая миграция всех моделей
	err = database.AutoMigrate(
		&models.User{},
		&models.Chat{},
		&models.ChatMember{},
		&models.Message{},
	)
	if err != nil {
		log.Fatal("❌ Ошибка при миграции базы данных: ", err)
	}

	DB = database
	fmt.Println("✅ Успешное подключение к базе данных и миграция выполнена")
}