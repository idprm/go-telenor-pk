package cmd

import (
	"database/sql"
	"sync"

	"github.com/idprm/go-telenor-pk/src/config"
	"github.com/idprm/go-telenor-pk/src/domain/repository"
	"github.com/idprm/go-telenor-pk/src/handler"
	"github.com/idprm/go-telenor-pk/src/logger"
	"github.com/idprm/go-telenor-pk/src/services"
)

type Processor struct {
	cfg    *config.Secret
	db     *sql.DB
	logger *logger.Logger
}

func NewProcessor(
	cfg *config.Secret,
	db *sql.DB,
	logger *logger.Logger,
) *Processor {
	return &Processor{
		cfg:    cfg,
		db:     db,
		logger: logger,
	}
}

func (p *Processor) MO(wg *sync.WaitGroup, message []byte) {
	/**
	 * -. Check Valid Prefix
	 * -. Check VIP Prefix
	 * -. Check Blacklist
	 * -. Check Active Sub
	 * -. Save Sub
	 * -. MT API
	 * -. Update Sub
	 * -. Save Transaction
	 * -. Save History
	 */

	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)

	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	moHandler := handler.NewMOHandler(p.cfg, p.logger, serviceService, subscriptionService)

	moHandler.IsActive()

	wg.Done()
}

func (p *Processor) DR(wg *sync.WaitGroup, message []byte) {

	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)

	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)

	handler.NewDRHandler(p.cfg, p.logger, serviceService, subscriptionService, transactionService)

	wg.Done()
}
