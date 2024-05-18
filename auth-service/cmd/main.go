package main

import (
	"database/sql"
	"log"

	"github.com/ALIK/project/auth-service/internal/repository"
	"github.com/ALIK/project/auth-service/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных
	db, err := sql.Open("postgres", "postgres://user:password@localhost/auth_service?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Создание экземпляра репозитория
	userRepo := repository.NewUserRepository(db)

	// Создание экземпляра сервиса и передача ему репозитория
	authService := service.NewAuthService(userRepo)

	// Здесь может быть настройка HTTP-сервера и обработка запросов
	// ...
}
