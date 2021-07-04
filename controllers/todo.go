package controllers

import (
  "fmt"
  "strconv"
  "github.com/gofiber/fiber/v2"
)

/**
 * Get index of todo by Id
 * @param todos list of todo
 * @param id searched ID
 * @return index if ID is existed in list. Otherwise, -1
 */
func GetIndexById(todos []*Todo, id int) int {
  for index, todo := range todos {
    if todo.Id == id {
      return index 
    }
  }

  return -1
}

/**
 * Todo struct. It's can be used for JSON
 */
type Todo struct {
  Id int `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
}

var todos = []*Todo{
  {
    Id: 1,
    Title: "Walk the dog",
    Completed: false,
  },
  {
    Id: 2,
    Title: "Walk the cat",
    Completed: false,
  },
}

/**
 * Response all todo
 */
func GetTodos(c *fiber.Ctx) error {
  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "success": true,
    "data": fiber.Map{
      "todos": todos,
    },
  })
}

/**
 * Create new todo
 */
func CreateTodo(c *fiber.Ctx) error {
  type Request struct {
    Title string `json:"title"`
  }

  var body Request

  err := c.BodyParser(&body)

  //error handles
  if err != nil {
    fmt.Println(err)
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "success": false,
      "message": "Cannot parse JSON",
    })
  }

  todo := &Todo{
    Id: len(todos) + 1,
    Title: body.Title,
    Completed: false,
  }

  //add new todo
  todos = append(todos, todo)

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "success": true,
    "data": fiber.Map{
      "todo": todo,
    },
  })
}

/**
 * Get todo by ID
 */
func GetTodo(c *fiber.Ctx) error {
  paramId := c.Params("id")

  //convert id to int
  id, err := strconv.Atoi(paramId)

  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "success": false,
      "message": "Cannot parse ID",
    })
  }

  for _, todo := range todos {
    if todo.Id == id {
      return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "success": true,
        "data": fiber.Map{
          "todo": todo,
        },
      })
    }
  }

  return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
    "success": false,
    "message": "Todo not found",
  })
}

/**
 * Update todo
 */
func UpdateTodo(c *fiber.Ctx) error {
  paramId := c.Params("id")

  type Request struct {
    Title string `json:"title"`
    Completed bool `json:"completed"` 
  }

  //convert id to int
  id, err := strconv.Atoi(paramId)

  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "success": false,
      "message": "Cannot parse ID",
    })
  }

  //Get update value
  var body Request
  err = c.BodyParser(&body)

  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "success": false,
      "message": "Cannot parse JSON",
    })
  }

  var index = GetIndexById(todos, id)
  
  if index != -1 {
    var updatedTodo *Todo = todos[index]
    updatedTodo.Title = body.Title
    updatedTodo.Completed = body.Completed
  } else {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "success": false,
      "message": "ID is not found",
    })
  }
  


  //Get updated todo
  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "success": true,
    "message": "Todo has been updated",
  }) 
}

func DeleteTodo(c *fiber.Ctx) error {
  paramId := c.Params("id")

  id, err := strconv.Atoi(paramId)

  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "success": false,
      "message": "Invalid ID",
    })
  }

  var index = GetIndexById(todos, id)
  
  if index != -1 {
    todos = append(todos[:index], todos[index + 1:]...)
  } else {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "success": false,
      "message": "ID is not found",
    })
  }
 

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "success": true,
    "message": "Todo has been deleted",
  })
}
