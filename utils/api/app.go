package api

import (
	"proyectoort/utils/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	serv          service.Service
	dataValidator *validator.Validate
}

func New(serv service.Service) *API {
	return &API{
		serv:          serv,
		dataValidator: validator.New(),
	}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{":23.22.45.156:5000"},
		AllowMethods:     []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderContentType},
		AllowCredentials: true,
	}))

	return e.Start(address)
}
