package services

import (
	"context"

	"github.com/go-logging/dto/request"
	"github.com/go-logging/dto/response"
)

type AlertService interface {
	CreateAlert(ctx context.Context, alertReq *request.AlertRequest) (*response.CreateAlertResponse, error)
}
