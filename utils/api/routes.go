package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {

	users := e.Group("/usuarios")
	form := e.Group("/formularios")
	users.POST("/registrar", a.RegisterUser)
	users.POST("/login", a.LoginUser)
	form.POST("/auditoria", a.GetFormByDate)

}
