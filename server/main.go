package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

type Task struct {
  gorm.Model
  Description string
  Completed bool
}

func main() {
  e := echo.New()
	e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.Use(middleware.CORS())

  db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

  if err != nil {
    panic("error accessing db")
  }
  db.AutoMigrate(&Task {})

  e.GET("/", func(c echo.Context) error {
		var tasks = get_tasks()
    return c.JSON(http.StatusOK, tasks)
	})

  e.GET("/:desc", func(c echo.Context) error {
    update_task(c.Param("desc"))
    return c.String(http.StatusOK, "OK")
  })

  e.GET("/new/:desc", func(c echo.Context) error {
    add_task(c.Param("desc"))
    return c.String(http.StatusOK, "OK")
  })

	e.Logger.Fatal(e.Start(":5050"))
}

func update_task(t string) {
  db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

  if err != nil {
    panic("error accessing db")
  }

  var task Task
  db.First(&task, "Description = ?", string(t))
  db.Model(&task).Update("Completed", !task.Completed)
}

func add_task(desc string)  {
  db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config {})

  if err != nil {
    panic("error accessing db")
  }

  var task Task = Task {Description: desc, Completed: false}
  db.Create(&task)
}

func get_tasks() []Task {
  db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

  if err != nil {
    panic("error accessing db")
  }

  var tasks [] Task
  db.Find(&tasks)
  return tasks
}
