package cmd

import (
	"log"

	"github.com/idprm/go-telenor-pk/src/app"
	"github.com/idprm/go-telenor-pk/src/config"
	"github.com/idprm/go-telenor-pk/src/datasource/mongodb/mdb"
	"github.com/idprm/go-telenor-pk/src/datasource/pgsql/db"
	"github.com/idprm/go-telenor-pk/src/datasource/rabbitmq"
	"github.com/idprm/go-telenor-pk/src/logger"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Webserver CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * LOAD CONFIG
		 */
		cfg, err := config.LoadSecret("secret.yaml")
		if err != nil {
			panic(err)
		}

		/**
		 * SETUP PGSQL
		 */
		db := db.InitDB(cfg)

		/**
		 * SETUP MONGODB
		 */
		mdb := mdb.InitMongoDB(cfg)

		log.Println(mdb)
		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger(cfg)

		/**
		 * SETUP RMQ
		 */
		queue := rabbitmq.InitQueue(cfg)

		/**
		 * SETUP CHANNEL
		 */
		queue.SetUpChannel(RMQ_EXCHANGETYPE, true, RMQ_MOEXCHANGE, true, RMQ_MOQUEUE)
		queue.SetUpChannel(RMQ_EXCHANGETYPE, true, RMQ_DREXCHANGE, true, RMQ_DRQUEUE)

		router := app.StartApplication(cfg, db, mdb, logger, queue)
		log.Fatal(router.Listen(":" + cfg.App.Port))
	},
}
