package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {

	users := e.Group("/usuarios")
	form := e.Group("/formularios")
	controls := e.Group("/controles")
	controls.POST("/registrar", a.RegisterControl)
	users.GET("", a.GetUsers)
	users.POST("/registrar", a.RegisterUser)
	users.POST("/login", a.LoginUser)
	form.POST("/registrar", a.RegisterFrom)
	controls.GET("", a.GetContorls)
	form.GET("", a.GetForms)

}
