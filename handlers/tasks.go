package handlers

import (
  "database/sql"
  "net/http"
  "strconv"

  "go-echo-vue/models"

  "github.com/labstack/echo"
)


// Небольшой трюк, который позволит возвращать произвольный JSON. Это просто map со ключами типа string и любым типом значения.
type H map[string]interface{}


// конечная точка GetTasks
func GetTasks(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    // получаем задачи из модели
    return c.JSON(http.StatusOK, models.GetTasks(db))
  }
}

// конечная точка PutTask
func PutTask(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    // создаём новую задачу
    var task models.Task
    // привязываем пришедший JSON в новую задачу
    c.Bind(&task)
    // добавим задачу с помощью модели
    id, err := models.PutTask(db, task.Name)
    // вернём ответ JSON при успехе
    if err == nil {
      return c.JSON(http.StatusCreated, H{
        "created": id,
      })
    // обработка ошибок
    } else {
      return err
    }
  }
}

// конечная точка DeleteTask
func DeleteTask(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    // используем модель для удаления задачи
    _, err := models.DeleteTask(db, id)
    // вернём ответ JSON при успехе
    if err == nil {
      return c.JSON(http.StatusOK, H{
        "deleted": id,
      })
    // обработка ошибок
    } else {
      return err
    }
  }
}
