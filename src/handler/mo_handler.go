package handler

import (
	"github.com/idprm/go-telenor-pk/src/config"
	"github.com/idprm/go-telenor-pk/src/logger"
	"github.com/idprm/go-telenor-pk/src/services"
)

type MOHandler struct {
	cfg                 *config.Secret
	logger              *logger.Logger
	serviceService      services.IServiceService
	subscriptionService services.ISubscriptionService
}

func NewMOHandler(
	cfg *config.Secret,
	logger *logger.Logger,
	serviceService services.IServiceService,
	subscriptionService services.ISubscriptionService,
) *MOHandler {

	return &MOHandler{
		cfg:                 cfg,
		logger:              logger,
		serviceService:      serviceService,
		subscriptionService: subscriptionService,
	}
}

func (h *MOHandler) IsRetry() {

}

func (h *MOHandler) IsActive() {

}
