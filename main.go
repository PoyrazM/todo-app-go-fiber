package main

import (
	"github.com/gofiber/fiber/v2"
	"todo-app-go-fiber/app"
	"todo-app-go-fiber/configs"
	"todo-app-go-fiber/dao"
	"todo-app-go-fiber/service"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()

	collectionName := "todos"
	dbClient := configs.GetCollection(configs.DB, collectionName)

	TodoRepositoryDb := dao.NewTodoRepositoryDB(dbClient)

	td := app.TodoHandler{Service: service.NewTodoService(TodoRepositoryDb)}

	appRoute.Post("/api/todos", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Delete("/api/todo/:requestid", td.DeleteTodoById)
	appRoute.Listen(":4000")
}
