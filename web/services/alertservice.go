package services

import (
	"context"

	"github.com/tandonraghav/go-logging/dto/request"
	"github.com/tandonraghav/go-logging/dto/response"
)

type AlertService interface {
	CreateAlert(ctx context.Context, alertReq *request.AlertRequest) (*response.CreateAlertResponse, error)
}
