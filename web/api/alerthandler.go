package api

import (
	"encoding/json"
	"net/http"

	logger "github.com/tandonraghav/go-logging/logging"
	"github.com/tandonraghav/go-logging/web/services"
	"github.com/tandonraghav/go-logging/web/services/impl"
)

type alertHandler struct {
	alertService services.AlertService
}

func NewAlertHandler() *alertHandler {
	return &alertHandler{
		alertService: impl.NewAlertService(),
	}
}

func (ah *alertHandler) CreateAlertHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger.GetLogger(ctx).Info( "Inside Create Alert Handler")
	resp, _ := ah.alertService.CreateAlert(ctx, nil)
	ar, _ := json.Marshal(resp)
	w.Write(ar)
	logger.GetLogger(ctx).Info( "Exit Create Alert Handler")
}