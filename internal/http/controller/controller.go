package controller

import (
	"github.com/ptaas-tool/gateway/internal/config"
	"github.com/ptaas-tool/gateway/internal/http/controller/handler"
	"github.com/ptaas-tool/gateway/internal/utils/jwt"
	"github.com/ptaas-tool/gateway/pkg/client"

	"github.com/ptaas-tool/base-api/pkg/models"
)

type (
	Controller struct {
		Config           config.Config
		JWTAuthenticator jwt.Authenticator
		Models           *models.Interface
		ErrHandler       handler.ErrorHandler
		Client           client.HTTPClient
		Metrics          *Metrics
	}

	Metrics struct {
		SuccessfulRequests int `json:"successful_requests"`
		FailedRequests     int `json:"failed_requests"`
		TotalExecutes      int `json:"total_executes"`
		TotalDownloads     int `json:"total_downloads"`
	}
)
