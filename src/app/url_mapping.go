package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	access_logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
	"github.com/idprm/go-telenor-pk/src/config"
	"github.com/idprm/go-telenor-pk/src/domain/repository"
	"github.com/idprm/go-telenor-pk/src/handler"
	"github.com/idprm/go-telenor-pk/src/logger"
	"github.com/idprm/go-telenor-pk/src/services"
	"github.com/wiliehidayat87/rmqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func mapUrls(cfg *config.Secret, db *sql.DB, mdb *mongo.Collection, logger *logger.Logger, rmpq rmqp.AMQP) *fiber.App {
	engine := html.New("./src/presenter/views", ".html")

	/**
	 * Init Fiber
	 */
	router := fiber.New(fiber.Config{
		Views: engine,
	})

	/**
	 * Access log on browser
	 */
	router.Use("/logs", filesystem.New(filesystem.Config{
		Root:         http.Dir(cfg.Log.Path),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "404.html",
		MaxAge:       3600,
	}))

	/**
	 * Write access logger
	 */
	file, err := os.OpenFile(cfg.Log.Path+"/access_log/log-"+time.Now().Format("2006-01-02")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	router.Use(requestid.New())
	router.Use(access_logger.New(access_logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   cfg.App.TimeZone,
		Output:     file,
	}))

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	router.Static("/static", path+"/public")

	campaignRepo := repository.NewCampaignRepository(db)
	campaignService := services.NewCampaignService(campaignRepo)

	blacklistRepo := repository.NewBlacklistRepository(db)
	blacklistService := services.NewBlacklistService(blacklistRepo)

	vipRepo := repository.NewVIPRepository(db)
	vipService := services.NewVIPService(vipRepo)

	serviceRepo := repository.NewServiceRepository(db)
	serviceService := services.NewServiceService(serviceRepo)

	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	incomingHandler := handler.NewIncomingHandler(cfg, logger, rmpq, campaignService, blacklistService, vipService, serviceService, subscriptionService)

	/**
	 * Routes Landing Page SUB & UNSUB
	 */
	router.Get("dreamleague", incomingHandler.DreamleagueSubPage)
	router.Get("dreamleague/unsub", incomingHandler.DreamleagueUnsubPage)
	router.Post("report/arpu", incomingHandler.ArpuTool)

	v1 := router.Group("v1")

	v1.Post("mo", incomingHandler.MessageOriginated)
	v1.Post("dr", incomingHandler.DeliveryReport)

	return router
}
