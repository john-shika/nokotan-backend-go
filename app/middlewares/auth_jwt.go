package middlewares

import (
	"example/app/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func AuthJWT(db *gorm.DB) echo.MiddlewareFunc {
	//var secretKey string
	//if secretKey = strings.Trim(viper.GetString("jwt_secret_key"), " "); secretKey == "" {
	//	panic("jwt secret key not found")
	//}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			var jwtToken *jwt.Token
			var session *models.Session
			if jwtToken, err = cores.GetJwtToken(c); err != nil {
				return c.JSON(http.StatusUnauthorized, cores.NewMessageUnauthorized(err.Error()))
			}
			if session, jwtToken, err = cores.ValidateJwtToken(db, jwtToken); err != nil {
				return c.JSON(http.StatusUnauthorized, cores.NewMessageUnauthorized(err.Error()))
			}
			cores.KeepVoid(jwtToken, session)
			c.Set("jwt_token", jwtToken)
			c.Set("session", session)
			return next(c)
		}
	}
}
