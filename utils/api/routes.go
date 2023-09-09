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
	document := e.Group("/document")
	controls.POST("/registrar", a.RegisterControl)
	controls.POST("/addForm", a.AddControlForm)
	users.GET("", a.GetUsers)
	users.POST("/registrar", a.RegisterUser)
	users.POST("/login", a.LoginUser)
	form.POST("/registrar", a.RegisterFrom)
	form.GET("/getByid", a.GetForm)
	form.POST("/addUser", a.RegisterUserForm)
	form.GET("/user", a.GetFormUser)
	controls.GET("", a.GetContorls)
	controls.GET("/sinF", a.GetControlsSinForm)
	controls.GET("/:id_formulario", a.GetControlsByForm)
	controls.POST("/addForm", a.AddControlForm)
	form.GET("", a.GetForms)
	obra.POST("/registrar", a.RegisterObra)
	obra.GET("", a.GetObras)
	obra.GET("/Byid", a.GetObra)
	obra.GET("/pisos", a.GetObrasPiso)
	obra.DELETE("/eliminar", a.DeleteObra)
	etapa.POST("/registrar", a.RegisterEtapa)
	piso.POST("/registrar", a.RegisterPiso)
	piso.GET("", a.GetPisos)
	piso.GET("/:id", a.GetPisosByObra)
	piso.POST("/addObra", a.RegisterObraPiso)
	check.POST("/registrar", a.RegisterCheck)
	check.POST("/addForm", a.RegisterCheckForm)
	document.POST("/addDocument", a.AddFormToPlanControl)
	document.GET("/:id", a.GetDocumentsByObra)

}
