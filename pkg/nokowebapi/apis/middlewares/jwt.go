package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/apis/schemas"
	"nokowebapi/cores"
	"nokowebapi/globals"
)

func AuthJWT(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			var token string
			var jwtToken cores.JwtTokenImpl
			var session *models.Session
			cores.KeepVoid(err, token, jwtToken, session)

			if token, err = extras.GetJwtTokenFromEchoContext(c); err != nil {
				return c.JSON(http.StatusUnauthorized, schemas.NewMessageBodyUnauthorized(err.Error(), nil))
			}

			jwtConfig := globals.Globals().GetJwtConfig()
			if jwtToken, err = cores.ParseJwtToken(token, jwtConfig.SecretKey); err != nil {
				return c.JSON(http.StatusUnauthorized, schemas.NewMessageBodyUnauthorized(err.Error(), nil))
			}

			jwtClaims := cores.Unwrap(cores.GetJwtClaimsFromJwtToken(jwtToken))
			jwtClaimsAccessData := cores.CvtJwtClaimsToJwtClaimsAccessData(jwtClaims)

			fmt.Println(jwtClaimsAccessData)

			c.Set("token", token)
			c.Set("jwt_token", jwtToken)
			c.Set("jwt_claims", jwtClaims)
			c.Set("jwt_claims_access_data", jwtClaimsAccessData)
			return next(c)
		}
	}
}
