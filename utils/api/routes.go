package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {

	users := e.Group("/usuarios")
	form := e.Group("/formularios")
	obra := e.Group("/obras")
	etapa := e.Group("/etapas")
	piso := e.Group("/pisos")
	check := e.Group("/checks")
	users.POST("/registrar", a.RegisterUser)
	users.POST("/login", a.LoginUser)
	form.GET("/auditoria", a.GetFormByDate)
	form.POST("/registrar", a.RegisterFrom)
	obra.POST("/registrar", a.RegisterObra)
	etapa.POST("/registrar", a.RegisterEtapa)
	piso.POST("/registrar", a.RegisterPiso)
	piso.POST("/addObra", a.RegisterObraPiso)
	check.POST("/registrar", a.RegisterCheck)
	check.POST("/addForm", a.RegisterCheckForm)

}
