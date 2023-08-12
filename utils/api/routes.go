package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {

	users := e.Group("/usuarios")
	form := e.Group("/formularios")
	controls := e.Group("/controles")
	obra := e.Group("/obras")
	etapa := e.Group("/etapas")
	piso := e.Group("/pisos")
	check := e.Group("/checks")
	controls.POST("/registrar", a.RegisterControl)
	controls.POST("/addForm", a.AddControlForm)
	users.GET("", a.GetUsers)
	users.POST("/registrar", a.RegisterUser)
	users.POST("/login", a.LoginUser)
	form.POST("/registrar", a.RegisterFrom)
	form.GET("/getByid", a.GetForm)
	controls.GET("", a.GetContorls)
	form.GET("", a.GetForms)
	obra.POST("/registrar", a.RegisterObra)
	obra.DELETE("/eliminar", a.DeleteObra)
	etapa.POST("/registrar", a.RegisterEtapa)
	piso.POST("/registrar", a.RegisterPiso)
	piso.POST("/addObra", a.RegisterObraPiso)
	check.POST("/registrar", a.RegisterCheck)
	check.POST("/addForm", a.RegisterCheckForm)

}
