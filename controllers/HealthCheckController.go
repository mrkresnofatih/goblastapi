package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/goblast"
)

type HealthCheckController struct {
}

func (h HealthCheckController) Register(echo *echo.Echo) {
	controllerRouter := goblast.ControllerRouter{
		MainRouter: echo,
		PathPrefix: "",
	}

	healthCheckEndpoint := &HealthCheckEndpoint{}
	controllerRouter.AddEndpoint(healthCheckEndpoint)

	controllerRouter.Build()
}

type HealthCheckEndpoint struct {
}

func (h HealthCheckEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	}
}

func (h HealthCheckEndpoint) GetPath() string {
	return "/healthcheck"
}

func (h HealthCheckEndpoint) Register(group *echo.Group) {
	group.Match([]string{http.MethodGet}, h.GetPath(), h.GetHandler())
}
