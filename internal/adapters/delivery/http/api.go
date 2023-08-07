package http

import (
	"github.com/labstack/echo/v4"
	"os"
	"reservify/internal/adapters/delivery/http/middlewares"
	"reservify/internal/adapters/delivery/http/routes"
)

type API interface {
	Serve()
	loadRoutes()
}

type Options struct {
}

type api struct {
	options      *Options
	group        *echo.Group
	echoInstance *echo.Echo
}

// NewAPI
// @title Reservify API
// @version 1.0
// @description Reserva de quartos para Hotéis
// @contact.name Eduardo Melo
// @contact.email deveduardomelo@gmail.com
// @BasePath /api
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization

func NewAPI(options *Options) API {
	echoInstance := echo.New()

	return &api{
		options:      options,
		group:        echoInstance.Group("/api"),
		echoInstance: echoInstance,
	}
}

func (instance *api) Serve() {
	instance.loadRoutes()
	instance.echoInstance.Use(middlewares.GuardMiddleware)
	port := os.Getenv("SERVER_PORT")

	instance.echoInstance.Logger.Fatal(instance.echoInstance.Start(":" + port))
}

func (instance *api) loadRoutes() {
	router := routes.New()

	router.Load(instance.group)
}
