package app

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-telenor-pk/src/config"
	"github.com/idprm/go-telenor-pk/src/logger"
	"github.com/wiliehidayat87/rmqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartApplication(cfg *config.Secret, db *sql.DB, mdb *mongo.Collection, logger *logger.Logger, rmpq rmqp.AMQP) *fiber.App {
	return mapUrls(cfg, db, mdb, logger, rmpq)
}
