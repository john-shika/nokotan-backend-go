package extras

import (
	"example/app/cores"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"strings"
)

type RouterImpl interface {
	Use(middleware ...echo.MiddlewareFunc)
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Any(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route
	Match(methods []string, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route
	Group(prefix string, middleware ...echo.MiddlewareFunc) (sg *echo.Group)
	RouteNotFound(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
}

func GetJwtToken(c echo.Context) (*jwt.Token, error) {
	var ok bool
	var err error
	var token string
	var jwtToken *jwt.Token
	if token = strings.Trim(c.Request().Header.Get("Authorization"), " "); token == "" {
		return nil, cores.ErrJwtTokenNotFound
	}
	if token, ok = strings.CutPrefix(token, "Bearer "); !ok {
		return nil, cores.ErrJwtTokenNotFound
	}
	if token = strings.Trim(token, " "); token == "" {
		return nil, cores.ErrJwtTokenNotFound
	}
	if jwtToken, err = ParseJwtToken(token); err != nil {
		return nil, err
	}
	return jwtToken, nil
}
