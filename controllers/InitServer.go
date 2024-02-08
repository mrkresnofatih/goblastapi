package controllers

import (
	"goblastapi-example/services"
	"net/http"
	"sync"

	"github.com/mrkresnofatih/goblast"
)

func InitServer(runState *sync.WaitGroup) {
	appServer := goblast.ApplicationServer{
		RunState: runState,
		CorsConfig: goblast.ApplicationServerCorsConfiguration{
			AllowMethods:     []string{http.MethodPost, http.MethodGet},
			AllowHeaders:     []string{"*"},
			AllowCredentials: true,
			AllowOrigins:     []string{"*"},
		},
		Port: "1323",
	}

	appServer.AddController(&ArithmeticController{
		ArithmeticService: &services.ArithmeticService{},
	})

	appServer.AddController(&HealthCheckController{})

	appServer.Initialize()
}
