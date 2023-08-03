package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {

	users := e.Group("/usuarios")
	form := e.Group("/formularios")
	obra := e.Group("/obras")
	users.POST("/registrar", a.RegisterUser)
	users.POST("/login", a.LoginUser)
	form.GET("/auditoria", a.GetFormByDate)
	form.POST("/registrar", a.RegisterFrom)
	obra.POST("/registrar", a.RegisterObra)

}
