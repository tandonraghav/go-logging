package app

import (
	"github.com/tandonraghav/go-logging/web/config"
	"net/http"

	"github.com/tandonraghav/go-logging/web/api"
	logger "github.com/tandonraghav/go-logging/web/logging"
	"github.com/tandonraghav/go-logging/web/middlewares"
	"github.com/go-chi/chi"
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
