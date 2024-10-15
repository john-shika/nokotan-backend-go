package todo

import (
	"example/app/middlewares"
	"example/app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func NewController(router vendor.RouterImpl, db *gorm.DB) cores.ControllerImpl {
	todos := router.Group("/todos")
	auth := router.Group("/todos", middlewares.AuthJWT(db))

	// all routers
	todos.GET("", getTodos(db))
	auth.POST("", createTodo(db))
	todos.GET("/:id", getTodo(db))
	auth.PUT("/:id", updateTodo(db))
	auth.DELETE("/:id", removeTodo(db))

	// fallback
	return cores.NewController(todos, db)
}

func getTodos(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var todos []models.Todo
		if tx := db.Find(&todos); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, cores.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, cores.NewMessageSuccess("successful get all todos", todos))
	}
}

func createTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var todo models.Todo
		if err = c.Bind(&todo); err != nil {
			return c.JSON(http.StatusBadRequest, cores.NewMessageBadRequest(err.Error()))
		}
		if tx := db.Create(&todo); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, cores.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusCreated, cores.NewMessageCreated("successful create todo", todo))
	}
}

func getTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo models.Todo
		id := c.Param("id")
		if tx := db.First(&todo, id); tx.Error != nil {
			return c.JSON(http.StatusNotFound, cores.NewMessageNotFound(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, cores.NewMessageSuccess("successful get todo", todo))
	}
}

func updateTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var todo models.Todo
		if err = c.Bind(&todo); err != nil {
			return c.JSON(http.StatusBadRequest, cores.NewMessageBadRequest(err.Error()))
		}
		id := c.Param("id")
		if tx := db.First(&todo, id); tx.Error != nil {
			return c.JSON(http.StatusNotFound, cores.NewMessageNotFound(tx.Error.Error()))
		}
		if tx := db.Save(&todo); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, cores.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, cores.NewMessageSuccess("successful update todo", nil))
	}
}

func removeTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var todo models.Todo
		if tx := db.First(&todo, id); tx.Error != nil {
			return c.JSON(http.StatusNotFound, cores.NewMessageNotFound(tx.Error.Error()))
		}
		if tx := db.Delete(&todo); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, cores.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, cores.NewMessageSuccess("successful delete todo", nil))
	}
}
