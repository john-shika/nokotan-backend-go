package middlewares

import (
	"example/app/cores"
	"example/app/cores/extras"
	"example/app/cores/schemas"
	"example/app/globals"
	"example/app/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func AuthJWT(db *gorm.DB) echo.MiddlewareFunc {
	//var secretKey string
	//if secretKey = strings.Trim(viper.GetString("jwt_secret_key"), " "); secretKey == "" {
	//	panic("jwt secret key not found")
	//}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			var token string
			var jwtToken cores.JwtTokenImpl
			var session *models.Session
			cores.KeepVoid(err, token, jwtToken, session)

			if token, err = extras.GetJwtToken(c); err != nil {
				return c.JSON(http.StatusUnauthorized, schemas.NewMessageBodyUnauthorized(err.Error(), nil))
			}

			jwtConfig := globals.GetGlobalJwtConfig()
			secretKey := jwtConfig.SecretKey
			if jwtToken, err = cores.ParseJwtToken(token, secretKey); err != nil {
				return c.JSON(http.StatusUnauthorized, schemas.NewMessageBodyUnauthorized(err.Error(), nil))
			}

			jwtClaims := cores.Unwrap(cores.GetJwtClaimsFromJwtToken[time.Time](jwtToken))
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
