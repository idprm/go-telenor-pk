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

func moProcessor(cfg *config.Secret, db *sql.DB, logger *logger.Logger, wg *sync.WaitGroup, message []byte) {
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

	serviceRepo := repository.NewServiceRepository(db)
	serviceService := services.NewServiceService(serviceRepo)

	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	moHandler := handler.MOHandler(cfg, logger, serviceService)

	wg.Done()
}

func drProcessor(cfg *config.Secret, db *sql.DB, logger *logger.Logger, wg *sync.WaitGroup, message []byte) {

	serviceRepo := repository.NewServiceRepository(db)
	serviceService := services.NewServiceService(serviceRepo)

	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	wg.Done()
}
