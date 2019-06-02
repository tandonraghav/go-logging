package impl

import (
	"context"
	"github.com/tandonraghav/go-logging/dto/request"
	"github.com/tandonraghav/go-logging/dto/response"
	logger "github.com/tandonraghav/go-logging/logging"
	"github.com/tandonraghav/go-logging/web/services"
)

type alertServiceImpl struct {
	db string
}

func NewAlertService() services.AlertService {
	return &alertServiceImpl{
		db: "test",
	}
}

func (a *alertServiceImpl) CreateAlert(ctx context.Context, alertReq *request.AlertRequest) (*response.CreateAlertResponse, error) {
	logger.GetLogger(ctx).Info( "Inside Create Alert")
	resp := response.CreateAlertResponse{
		AlertID: 123,
	}
	//i1:=1
	//i2:=0
	//_ =i1/i2


	return &resp, nil
}
