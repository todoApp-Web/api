package controllers

import (
    "context"
    "log"
	"time"
    "github.com/gofiber/fiber/v2"
    "api/config"
    "api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTodos(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection("todos")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var todos []models.Todo
	
	cursor, err := todoCollection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	if err != nil {
    	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "success": false,
            "message": "Todos Not found",
            "error":   err,
        })
	}
	
	for cursor.Next(ctx) {
        var todo models.Todo
        cursor.Decode(&todo)
        todos = append(todos, todo)
    }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "data":    todos,
        "success": true,
        "message": "Todo show successfully",
    })

}

func AddTodo(c *fiber.Ctx) error {
    todoCollection := config.MI.DB.Collection("todos")
    todo := new(models.Todo)

    if err := c.BodyParser(todo); err != nil {
        log.Println(err)
        return c.Status(400).JSON(fiber.Map{
            "success": false,
            "message": "Failed to parse body",
            "error":   err,
        })
    }

    result, err := todoCollection.InsertOne(context.TODO(), todo)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "Todo failed to insert",
            "error":   err,
        })
    }
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "data":    result,
        "success": true,
        "message": "Todo inserted successfully",
    })

}