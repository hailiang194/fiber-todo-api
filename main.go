package main

import (
  "os"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/logger"
  "github.com/hailiang194/fiber-todo-api/routes"
)

/**
 * Setup route
 */
func setupRoutes(app *fiber.App) {

  //Root
  app.Get("/", func(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
      "success": true,
      "message": "You're at the endpoint",
    })
  })

  //API group
  api := app.Group("/api")

  api.Get("", func(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
      "success": true,
      "message": "You're at the api endpoint",
    })
  })

  routes.TodoRoute(api.Group("/todos"))
}

func main() {
  app := fiber.New()

  app.Use(logger.New())

  setupRoutes(app)

  port, ok := os.LookupEnv("PORT")

  if ok == false {
    port = "8000"
  }

  err := app.Listen(":"+port)
  if err != nil {
    panic(err)
  }
}
