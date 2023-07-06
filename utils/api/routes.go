package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {

	users := e.Group("/usuarios")

	users.POST("/registrar", a.RegisterUser)
	users.POST("/login", a.LoginUser)

}
