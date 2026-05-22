package main

import (
	"log"
	"sports-registration/internal/database"
	"sports-registration/server"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("🚀 Запуск приложения Kinetic Sports Registration...")

	// Загрузка переменных окружения из .env
	godotenv.Load()

	// Инициализация базы данных
	database.Init()

	server.StartServer()
}
