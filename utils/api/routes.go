package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {

	//usuarios
	users := e.Group("/usuarios")
	users.GET("", a.GetUsers)
	users.POST("/registrar", a.RegisterUser)
	users.POST("/login", a.LoginUser)
	users.POST("/rol", a.RegisterUserRol)
	users.GET("/rol:id", a.GetUserRoles)

	//roles
	roles := e.Group("/roles")
	roles.GET("", a.GetRoles)

	// controles
	controls := e.Group("/controles")
	controls.GET("", a.GetContorls)
	controls.POST("/registrar", a.RegisterControl)
	//obras
	obra := e.Group("/obras")
	obra.POST("/registrar", a.RegisterObra)
	// etapas
	etapa := e.Group("/etapas")
	etapa.POST("/registrar", a.RegisterEtapa)
	//pisos
	piso := e.Group("/pisos")
	piso.POST("/registrar", a.RegisterPiso)
	piso.POST("/addObra", a.RegisterObraPiso)
	piso.GET("", a.GetPisos)

	//cheks
	check := e.Group("/checks")
	check.POST("/registrar", a.RegisterCheck)
	check.POST("/addForm", a.RegisterCheckForm)

	//formularios
	form := e.Group("/formularios")
	form.POST("/registrar", a.RegisterFrom)
	form.GET("", a.GetForms)

	//fotos
	foto := e.Group("/fotos")
	foto.POST("/registrar", a.RegisterPhoto)
	foto.GET("", a.GetFotosForm)
	foto.GET("download/:id", a.DownloadPhoto)
}
