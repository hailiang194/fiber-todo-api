package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hailiang194/fiber-todo-api/controllers"
)

/**
 * configure route
 */
func TodoRoute(route fiber.Router) {
	//Get all todos
  route.Get("", controllers.GetTodos)
	//Create new todo
	route.Post("", controllers.CreateTodo)
	//Get todo by id
	route.Get("/:id", controllers.GetTodo)
	//Update totod
	route.Post("/:id", controllers.UpdateTodo)
	//Delete todos
	route.Delete("/:id", controllers.DeleteTodo)
}
