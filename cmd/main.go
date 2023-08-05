package main

import (
	"fmt"
	"go-restapi/internal/adapter/database"
	apphttp "go-restapi/internal/adapter/http"
	"go-restapi/internal/domain"
	"go-restapi/internal/usecase"
	"go-restapi/pkg/server"
	"net/http"
)

func main() {
	// Postgres DSN
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=tasks sslmode=disable timezone=UTC connect_timeout=5"

	// Database bağlantısı oluştur
	db, err := database.NewDBConnection(dsn)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Domain ve Repository'lerin oluşturulması
	taskRepo := domain.NewTaskRepository(db)
	userRepo := domain.NewUserRepository(db)

	// Use case'lerin oluşturulması
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// HTTP handler'ların oluşturulması
	taskHandler := apphttp.NewTaskHandler(taskUsecase)
	userHandler := apphttp.NewUserHandler(userUsecase)

	// Router ve Server oluşturma işlemleri
	router := apphttp.NewRouter(taskHandler, userHandler)
	server := server.NewServer(":8080", router)

	fmt.Println("Server is running on :8080")
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("Failed to start server:", err)
	}
}
