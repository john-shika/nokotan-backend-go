package todo

import (
	"example/app/extras"
	"example/app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func NewController(router extras.RouterImpl, db *gorm.DB) extras.ControllerImpl {
	todos := router.Group("/todos")

	// all routers
	todos.GET("", getTodos(db))
	todos.POST("", createTodo(db))
	todos.GET("/:id", getTodo(db))
	todos.PUT("/:id", updateTodo(db))
	todos.DELETE("/:id", removeTodo(db))

	// fallback
	return extras.NewController(todos, db)
}

func getTodos(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var todos []models.Todo
		if tx := db.Find(&todos); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, extras.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, extras.NewMessageSuccess("successful get all todos", todos))
	}
}

func createTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var todo models.Todo
		if err = c.Bind(&todo); err != nil {
			return c.JSON(http.StatusBadRequest, extras.NewMessageBadRequest(err.Error()))
		}
		if tx := db.Create(&todo); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, extras.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusCreated, extras.NewMessageCreated("successful create todo", todo))
	}
}

func getTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo models.Todo
		id := c.Param("id")
		if tx := db.First(&todo, id); tx.Error != nil {
			return c.JSON(http.StatusNotFound, extras.NewMessageNotFound(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, extras.NewMessageSuccess("successful get todo", todo))
	}
}

func updateTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var todo models.Todo
		if err = c.Bind(&todo); err != nil {
			return c.JSON(http.StatusBadRequest, extras.NewMessageBadRequest(err.Error()))
		}
		id := c.Param("id")
		if tx := db.First(&todo, id); tx.Error != nil {
			return c.JSON(http.StatusNotFound, extras.NewMessageNotFound(tx.Error.Error()))
		}
		if tx := db.Save(&todo); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, extras.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, extras.NewMessageSuccess("successful update todo", nil))
	}
}

func removeTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var todo models.Todo
		if tx := db.First(&todo, id); tx.Error != nil {
			return c.JSON(http.StatusNotFound, extras.NewMessageNotFound(tx.Error.Error()))
		}
		if tx := db.Delete(&todo); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, extras.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, extras.NewMessageSuccess("successful delete todo", nil))
	}
}
