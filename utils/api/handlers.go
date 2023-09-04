package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"proyectoort/encryption"
	"proyectoort/utils/api/dtos"
	"proyectoort/utils/models"
	"proyectoort/utils/service"

	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Username, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "El usuario ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, responseMessage{Message: "usuario creado"})
}
func (a *API) RegisterFrom(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.DocumentAudit{}
	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterFrom(ctx, params.Nombre, params.Informacion, params.Version, params.ControlID, params.UsuarioID)
	if err != nil {
		if err == service.ErrFormAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "El formulario ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, responseMessage{Message: "Se creo el formulario"})
}
func (a *API) RegisterControl(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterControl{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterControl(ctx, params.Descripcion, params.Tipo)
	if err != nil {
		if err == service.ErrFormAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "El control ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, responseMessage{Message: "Se a creado el control"})
}
func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	u, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	roles, err := a.serv.GetUsersRole(ctx, u.ID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	token, err := encryption.SignedLoginToken(u)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	c.SetCookie(cookie)
	for _, r := range roles {
		if r.RoleID == 1 {
			return c.JSON(http.StatusOK, map[string]string{"redirect": "/admin"})
		} else if r.RoleID == 2 {
			return c.JSON(http.StatusOK, map[string]string{"redirect": "/user"})
		} else if r.RoleID == 3 {
			return c.JSON(http.StatusOK, map[string]string{"redirect": "/manager"})
		}
	}
	return c.JSON(http.StatusUnauthorized, responseMessage{Message: "No tiene permisos para iniciar sesión"})
}
func (a *API) GetUsers(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.Usuarios{}
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}
	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	u, err := a.serv.GetUsers(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error al obtener los usuarios"})
	}
	return c.JSON(http.StatusOK, u)
}
func (a *API) GetForms(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DocumentAudit{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	forms, err := a.serv.GetForms(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error al obtener los formularios"})
	}
	return c.JSON(http.StatusOK, forms)
}
func (a *API) GetContorls(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.RegisterControl{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}
	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	control, err := a.serv.GetControls(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error al obtener los controles"})
	}
	return c.JSON(http.StatusOK, control)
}
func (a *API) RegisterObra(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterObra{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterObra(ctx, params.Nombre)
	if err != nil {
		if err == service.ErrObraAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "La Obra ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, nil)
}
func (a *API) RegisterEtapa(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterEtapa{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterEtapa(ctx, params.Nombre)
	if err != nil {
		if err == service.ErrEtapaAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "La Etapa ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, nil)
}
func (a *API) RegisterPiso(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterPiso{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterPiso(ctx, params.ID, params.Numero)
	if err != nil {
		if err == service.ErrPisoAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "El piso ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, responseMessage{Message: "Se creo el piso"})
}
func (a *API) RegisterObraPiso(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.ConexionObraPiso{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.AddObraPiso(ctx, int64(params.ObraID), int64(params.PisoID))
	if err != nil {
		if err == service.ErrPisoObraAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "La conexion ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, nil)
}
func (a *API) RegisterCheck(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterCheck{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterCheck(ctx, params.Estado, params.Observaciones, params.Version, params.Fecha)
	if err != nil {
		if err == service.ErrFormAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "El check ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, responseMessage{Message: "Se creo el check"})
}
func (a *API) RegisterCheckForm(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.ConexionCheckForm{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.AddCheckForm(ctx, int64(params.CheckID), int64(params.FormularioID))
	if err != nil {
		if err == service.ErrCheckFormAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "La conexion ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, responseMessage{Message: "Se creo el checkqueo para el formulario"})
}
func (a *API) GetRoles(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.Roles{}
	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}
	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	r, err := a.serv.GetAllRoles(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error al obtener los roles"})
	}
	return c.JSON(http.StatusOK, r)
}
func (a *API) RegisterUserRol(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.UsuarioRol{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.AddUserRole(ctx, params.UserID, params.RoleID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error al asignar el rol"})
	}
	return c.JSON(http.StatusOK, responseMessage{Message: "Se asigna el rol para el usuario selecionado"})
}
func (a *API) GetUserRoles(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.UsuarioRol{}
	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}
	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	ru, err := a.serv.GetUsersRole(ctx, params.UserID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error al obtener los roles"})
	}
	return c.JSON(http.StatusOK, ru)
}
func (a *API) AddForm(c echo.Context) error {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	claims, err := encryption.ParseLoginJWT(cookie.Value)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	email := claims["email"].(string)

	ctx := c.Request().Context()
	params := dtos.DocumentAudit{}

	err = c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	f := models.Formulario{
		ID:          params.ID,
		Informacion: params.Informacion,
		Version:     params.Version,
		Nombre:      params.Nombre,
	}

	err = a.serv.AddForm(ctx, email, f)
	if err != nil {
		log.Println(err)

		if err == service.ErrInvalidPermissions {
			return c.JSON(http.StatusForbidden, responseMessage{Message: "Error de permisos"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusOK, nil)
}
func (a *API) GetPisos(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.RegisterPiso{}
	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}
	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	p, err := a.serv.GetPisos(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error al obtener los pisos"})
	}
	return c.JSON(http.StatusOK, p)
}
func (a *API) RegisterPhoto(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.FotoDTO{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterPhoto(ctx, params.Nombre, params.Notas, params.FormularioID)
	if err != nil {
		if err == service.ErrFotoAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "la foto ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, responseMessage{Message: "Se creo la foto para el formulario"})
}
func (a *API) GetFotosForm(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.FotoDTO{}
	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}
	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	fo, err := a.serv.GetPhotos(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error al obtener las fotos"})
	}
	return c.JSON(http.StatusOK, fo)
}
func (a *API) DownloadPhoto(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.FotoDTO{}

	// Lógica para obtener la ruta del archivo de la foto desde la base de datos usando el ID
	filePath, err := a.serv.GetPhotoFilePath(ctx, params.ID)
	if err != nil {
		// Manejar el error, por ejemplo, si la foto no existe
		return c.String(http.StatusNotFound, "Foto no encontrada")
	}

	// Leer el contenido del archivo de la foto
	fotoBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		// Manejar el error si no se puede leer el archivo
		return c.String(http.StatusInternalServerError, "Error al leer la foto")
	}

	// Configurar la respuesta HTTP
	c.Response().Header().Set("Content-Disposition", "attachment; filename=foto.jpg") // Cambia el nombre de archivo según la foto
	c.Response().Header().Set("Content-Type", "image/jpeg")                           // Cambia el tipo de contenido según la foto

	// Enviar la foto como respuesta
	return c.Blob(http.StatusOK, "image/jpeg", fotoBytes)
}
