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
	form.PUT("/:id", a.UpdateForm)
	form.GET("/:id", a.GetControlsSinForm)
	form.DELETE("/eliminar/:id", a.DeleteFormulario)
	controls.GET("", a.GetContorls)
	controls.POST("/byForm", a.GetControlsByForm)
	controls.POST("/addForm", a.AddControlForm)
	controls.PUT("/:id", a.UpdateControl)
	controls.DELETE("/:control_id/formulario", a.DeleteControlForm)
	controls.DELETE("/eliminar/:id", a.DeleteControl)
	form.GET("", a.GetForms)
	obra.POST("/registrar", a.RegisterObra)
	obra.GET("", a.GetObras)
	obra.GET("/Byid", a.GetObra)
	obra.GET("/pisos", a.GetObrasPiso)
	obra.DELETE("/eliminar/:id", a.DeleteObra)
	obra.PUT("/:id", a.UpdateObra)
	etapa.POST("/registrar", a.RegisterEtapa)
	piso.POST("/registrar", a.RegisterPiso)
	piso.GET("", a.GetPisos)
	piso.GET("/:id", a.GetPisosByObra)
	piso.POST("/addObra", a.RegisterObraPiso)
	piso.PUT("/:id", a.UpdatePiso)
	piso.DELETE("/eliminar/:id", a.DeletePiso)
	check.POST("/registrar", a.RegisterCheck)
	check.POST("/addForm", a.RegisterCheckForm)
	check.PUT("/:id", a.UpdateCheck)
	check.GET("/document/:document_id", a.GetDocumentChecks)
	document.POST("/addDocument", a.AddFormToPlanControl)
	document.GET("/:id", a.GetDocumentsByObra)
	document.DELETE("/eliminar/:id", a.DeleteDocument)

	document.GET("/export/:id", a.ExportDocument)
	document.GET("/export/obra/:id", a.ExportDocumentsByObra)
	document.GET("/export", a.ExportDocument)
	foto := e.Group("/fotos")
	foto.POST("/registrar", a.RegisterPhoto)
	foto.GET("/formulario", a.GetFotosForm)
	foto.GET("download/:id", a.DownloadPhoto)

}
