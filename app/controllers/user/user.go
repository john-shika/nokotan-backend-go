package user

import (
	"crypto/sha256"
	"example/app/cores"
	"example/app/cores/apis"
	"example/app/cores/extras"
	"example/app/cores/schemas"
	"example/app/middlewares"
	"example/app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func NewController(router extras.RouterImpl, db *gorm.DB) apis.ControllerImpl {
	users := router.Group("/users")
	auth := router.Group("/users", middlewares.AuthJWT(db))

	// free authentication
	users.POST("/login", login(db))
	users.POST("/logout", logout(db))
	users.POST("/register", register(db))

	// all routers
	auth.GET("", getUsers(db))
	auth.POST("", createUser(db))
	auth.GET("/:id", getUser(db))
	auth.PUT("/:id", updateUser(db))
	auth.DELETE("/:id", removeUser(db))

	// fallback
	return apis.NewController(users, db)
}

func login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var token string
		var user models.User
		var session *models.Session
		cores.KeepVoid(err, token, user, session)
		if err = c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, schemas.NewMessageBadRequest(err.Error()))
		}

		user.Email = strings.Trim(user.Email, " ")
		user.Username = strings.Trim(user.Username, " ")
		user.Password = strings.Trim(user.Password, " ")

		if user.Email == "" && user.Username == "" {
			return c.JSON(http.StatusBadRequest, schemas.NewMessageBadRequest("email or username is empty"))
		}

		if user.Password == "" {
			return c.JSON(http.StatusBadRequest, schemas.NewMessageBadRequest("password is empty"))
		}

		email := user.Email
		username := user.Username
		password := cores.NewBase64EncodeToString(cores.GetArraySliceSize32(sha256.Sum256([]byte(user.Password))))

		if email != "" {
			if tx := db.First(&user, "email = ? AND password = ?", email, password); tx.Error != nil {
				return c.JSON(http.StatusUnauthorized, schemas.NewMessageUnauthorized(tx.Error.Error()))
			}
		} else {
			if tx := db.First(&user, "username = ? AND password = ?", username, password); tx.Error != nil {
				return c.JSON(http.StatusUnauthorized, schemas.NewMessageUnauthorized(tx.Error.Error()))
			}
		}

		if session, token, err = apis.NewSession(db, &user); err != nil {
			return c.JSON(http.StatusInternalServerError, schemas.NewMessageError(err.Error()))
		}

		return c.JSON(http.StatusOK, schemas.NewMessageSuccess("successful login", map[string]any{
			"token": token,
		}))
	}
}

func logout(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var ok bool
		var session *models.Session
		if session, ok = c.Get("session").(*models.Session); !ok {
			return c.JSON(http.StatusInternalServerError, schemas.NewMessageError("failed to get session"))
		}
		// delete current session
		if tx := db.Delete(&session); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, schemas.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, schemas.NewMessageSuccess("successful logout", nil))
	}
}

func register(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, schemas.NewMessageBadRequest(err.Error()))
		}
		if user.Email == "" || user.Username == "" || user.Password == "" {
			return c.JSON(http.StatusBadRequest, schemas.NewMessageBadRequest("email, username or password is empty"))
		}
		email := user.Email
		username := user.Username
		password := sha256.Sum256([]byte(user.Password))

		// validate email, username, password
		// email regex
		// username regex
		// password regex

		user.Email = email
		user.Username = username
		user.Password = string(password[:])
		user.Role = "user"
		// create user
		if tx := db.Create(&user); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, schemas.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, schemas.NewMessageSuccess("successful register", nil))
	}
}

func getUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []models.User
		if tx := db.Find(&users); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, schemas.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, schemas.NewMessageSuccess("successful get all users", users))
	}
}

func createUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var user models.User
		if err = c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, schemas.NewMessageBadRequest(err.Error()))
		}
		if tx := db.Create(&user); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, schemas.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusCreated, schemas.NewMessageCreated("successful create user", user))
	}
}

func getUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		id := c.Param("id")
		if tx := db.First(&user, id); tx.Error != nil {
			return c.JSON(http.StatusNotFound, schemas.NewMessageNotFound(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, schemas.NewMessageSuccess("successful get user", user))
	}
}

func updateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var user models.User
		if err = c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, schemas.NewMessageBadRequest(err.Error()))
		}
		id := c.Param("id")
		if tx := db.First(&user, id); tx.Error != nil {
			return c.JSON(http.StatusNotFound, schemas.NewMessageNotFound(tx.Error.Error()))
		}
		if tx := db.Save(&user); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, schemas.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, schemas.NewMessageSuccess("successful update user", nil))
	}
}

func removeUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var user models.User
		if tx := db.First(&user, id); tx.Error != nil {
			return c.JSON(http.StatusNotFound, schemas.NewMessageNotFound(tx.Error.Error()))
		}
		if tx := db.Delete(&user); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, schemas.NewMessageError(tx.Error.Error()))
		}
		return c.JSON(http.StatusOK, schemas.NewMessageSuccess("successful remove user", nil))
	}
}
