package api

import (
	"log"
	"net/http"
	"proyectoort/encryption"
	"proyectoort/utils/api/dtos"
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

	return c.JSON(http.StatusCreated, nil)
}
func (a *API) RegisterFrom(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.DocumentAudit{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Solicitud no válida"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterFrom(ctx, params.Nombre, params.Informacion, params.Version, params.Fecha.Format("23/12/2023"), int(params.EtapaID), int(params.UsuarioID))
	if err != nil {
		if err == service.ErrFormAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "El formulairo ya existe"})
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
	return c.JSON(http.StatusOK, map[string]string{"usuario logueado": "true"})

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

	err = a.serv.RegisterPiso(ctx, int64(params.Numero))
	if err != nil {
		if err == service.ErrPisoAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "El piso ya existe"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error interno del servidor"})
	}

	return c.JSON(http.StatusCreated, nil)
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
