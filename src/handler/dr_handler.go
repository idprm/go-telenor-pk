package handler

import (
	"github.com/idprm/go-telenor-pk/src/config"
	"github.com/idprm/go-telenor-pk/src/logger"
	"github.com/idprm/go-telenor-pk/src/services"
)

type DRHandler struct {
	cfg                 *config.Secret
	logger              *logger.Logger
	serviceService      services.IServiceService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
}

func NewDRHandler(
	cfg *config.Secret,
	logger *logger.Logger,
	serviceService services.IServiceService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService) *DRHandler {

	return &DRHandler{
		cfg:                 cfg,
		logger:              logger,
		serviceService:      serviceService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
	}
}
