package routes

import (
    "github.com/gofiber/fiber/v2"
    "api/controllers"
)

func TodosRoute(route fiber.Router) {
    route.Get("/", controllers.GetAllTodos)
    route.Post("/", controllers.AddTodo)
}