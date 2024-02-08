package controllers

import (
	ae "goblastapi-example/endpoints/arithmetic"
	"goblastapi-example/models"
	"goblastapi-example/services"

	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/goblast"
)

type ArithmeticController struct {
	ArithmeticService services.IArithmeticService
}

func (a ArithmeticController) Register(echo *echo.Echo) {
	controllerRouter := goblast.ControllerRouter{
		MainRouter: echo,
		PathPrefix: "/api/arithmetic",
	}

	addEndpoint := &ae.ArithmeticAddEndpoint{
		ArithmeticService: a.ArithmeticService,
	}
	ctxfulAddEndpoint := &goblast.ContextfulReqEndpoint[models.ArithmeticRequest]{
		Endpoint: addEndpoint,
	}
	clearAuthEndpoint := &goblast.ClearAuthEndpoint{
		Endpoint: ctxfulAddEndpoint,
	}
	controllerRouter.AddEndpoint(clearAuthEndpoint)

	controllerRouter.Build()
}
