package main

import (
	"inventory/internal/database"
	"inventory/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к БД
	database.Connect()

	// Создание роутера Gin
	r := gin.Default()

	// Отдаём HTML-файлы из папки frontend
	r.Static("/static", "./frontend/static")
	r.StaticFile("/", "./frontend/index.html")
	r.StaticFile("/instruments", "./frontend/instruments.html")
	r.StaticFile("/employees", "./frontend/employees.html")
	r.StaticFile("/issues", "./frontend/issues.html")

	// ─── API маршруты ───────────────────────────────

	api := r.Group("/api")
	{
		// Сотрудники - CRUD
		api.GET("/employees",       handlers.GetEmployees)   // R - список
		api.GET("/employees/:id",   handlers.GetEmployee)    // R - один
		api.POST("/employees",      handlers.CreateEmployee) // C - создать
		api.PUT("/employees/:id",   handlers.UpdateEmployee) // U - изменить
		api.DELETE("/employees/:id",handlers.DeleteEmployee) // D - удалить

		// Инструменты - CRUD
		api.GET("/instruments",        handlers.GetInstruments)   // R - список
		api.GET("/instruments/:id",    handlers.GetInstrument)    // R - один
		api.POST("/instruments",       handlers.CreateInstrument) // C - создать
		api.PUT("/instruments/:id",    handlers.UpdateInstrument) // U - изменить
		api.DELETE("/instruments/:id", handlers.DeleteInstrument) // D - удалить

		// Выдачи - CRUD
		api.GET("/issues",           handlers.GetIssues)    // R - список
		api.POST("/issues",          handlers.CreateIssue)  // C - выдать
		api.PUT("/issues/:id/return",handlers.ReturnIssue)  // U - вернуть
		api.DELETE("/issues/:id",    handlers.DeleteIssue)  // D - удалить
	}

	// Запуск сервера на порту 8080
	r.Run(":8080")
}
