package main

import (
  "gorm.io/gorm"
  "github.com/joho/godotenv"
  "gorm.io/driver/mysql"
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

  envs, err := godotenv.Read(".env")

  if err != nil {
    panic("error reading env")
  }

  dburl := envs["DBURL"]

  db, err := gorm.Open(mysql.Open(dburl), &gorm.Config{})

  if err != nil {
    panic("error accessing db")
  }

  db.AutoMigrate(&Task {})

  tasks := init_tasks(db)

  e.GET("/", func(c echo.Context) error {
    t := map_to_arr(tasks)
    return c.JSON(http.StatusOK, t)
	})

  e.GET("/:desc", func(c echo.Context) error {
    update_task(db, c.Param("desc"), tasks)
    return c.String(http.StatusOK, "OK")
  })

  e.GET("/new/:desc", func(c echo.Context) error {
    add_task(db, c.Param("desc"), tasks)
    return c.String(http.StatusOK, "OK")
  })

  e.GET("/del/:desc", func(c echo.Context) error {
    del_task(db, c.Param("desc"), tasks)
    return c.JSON(http.StatusOK, "OK")
  })

	e.Logger.Fatal(e.Start(":5050"))
}

func update_task(db *gorm.DB, desc string, tasks map[string]bool) {
  var task Task
  tasks[desc] = !tasks[desc]
  go db.First(&task, "Description = ?", string(desc))
  go db.Model(&task).Update("Completed", !task.Completed)
}

func add_task(db  *gorm.DB, desc string, tasks map[string]bool)  {
  tasks[desc] = false
  task := Task { Description: desc, Completed: false }
  go db.Create(&task)
}

func init_tasks(db *gorm.DB) map[string]bool {
  var tasks [] Task
  db.Find(&tasks)

  m := make(map[string]bool)
  
  for _, val := range tasks {
    m[val.Description] = val.Completed
  }
  return m
}

func del_task(db *gorm.DB, desc string, tasks map[string]bool) {
  var task Task = Task { Description: desc, Completed: true }
  delete(tasks, desc)
  go db.Where("Description = ?", desc).Delete(&task)
}

func map_to_arr(tasks map[string]bool) []Task {
  t := make([]Task, len(tasks))
  var i = 0
  for key, value := range tasks {
    t[i] = Task{ Description: key, Completed: value }
    i+=1
  }
  return t
}
