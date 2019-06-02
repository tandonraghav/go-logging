package mux

import (
	"github.com/tandonraghav/go-logging/config"
	"net/http"

	"github.com/go-chi/chi"
	logger "github.com/tandonraghav/go-logging/logging"
	"github.com/tandonraghav/go-logging/web/api"
	"github.com/tandonraghav/go-logging/web/middlewares"
)

func InitilizeMux() {
	r := chi.NewRouter()
	r.Use(middlewares.DefaultMiddleware)

	RegisterHandlers(r)

	port := config.GetPort()
	logger.GetLogger(nil).Info( "Server starting at port "+port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		logger.GetLogger(nil).Fatal( "Err from RegisterHandlers:"+err.Error())
	}
}

func RegisterHandlers(r *chi.Mux) {
	registerAlertHandler(r)
}

func registerAlertHandler(r *chi.Mux) {
	alertHandler := api.NewAlertHandler()
	r.Get("/test", alertHandler.CreateAlertHandler)
}
