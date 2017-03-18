package main

import (
  "database/sql"
  "go-echo-vue/handlers"

  "github.com/labstack/echo"
  _ "github.com/mattn/go-sqlite3"
)

func main() {

  db := initDB("storage.db")
  migrate(db)

  // Create a new instance of Echo
  e := echo.New()

  e.File("/", "public/index.html")
  e.GET("/tasks", handlers.GetTasks(db))
  e.PUT("/tasks", handlers.PutTask(db))
  e.DELETE("/tasks/:id", handlers.DeleteTask(db))

  // Start as a web server
  e.Logger.Fatal(e.Start(":8000"))
}


func initDB(filepath string) *sql.DB {
  //откроем файл или создадим его
  db, err := sql.Open("sqlite3", filepath)

  // проверяем ошибки и выходим при их наличии
  if err != nil {
    panic(err)
  }

  // если ошибок нет, но не можем подключиться к базе данных,
  // то так же выходим
  if db == nil {
    panic("db nil")
  }
  return db
}


func migrate(db *sql.DB) {
  sql := `
  CREATE TABLE IF NOT EXISTS tasks(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name VARCHAR NOT NULL
  );
  `

  _, err := db.Exec(sql)
  // выходим, если будут ошибки с SQL запросом выше
  if err != nil {
    panic(err)
  }
}
