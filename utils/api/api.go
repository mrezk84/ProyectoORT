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
		AllowOrigins:     []string{"127.0.0.1:5000"},
		AllowMethods:     []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderContentType},
		AllowCredentials: true,
	}))

	e.Use(middleware.Static("static"))

	return e.Start(address)
}
