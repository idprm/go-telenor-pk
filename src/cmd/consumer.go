package cmd

import (
	"fmt"
	"sync"

	"github.com/idprm/go-telenor-pk/src/config"
	"github.com/idprm/go-telenor-pk/src/datasource/pgsql/db"
	"github.com/idprm/go-telenor-pk/src/datasource/rabbitmq"
	"github.com/idprm/go-telenor-pk/src/logger"
	"github.com/spf13/cobra"
)

var consumerMOCmd = &cobra.Command{
	Use:   "consumer_mo",
	Short: "Consumer MO Service CLI",
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

		messagesData := queue.Subscribe(1, false, RMQ_MOQUEUE, RMQ_MOEXCHANGE, RMQ_MOQUEUE)

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		processor := NewProcessor(cfg, db, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {
				wg.Add(1)
				processor.MO(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerDRCmd = &cobra.Command{
	Use:   "consumer_dr",
	Short: "Consumer DR Service CLI",
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
		queue.SetUpChannel(RMQ_EXCHANGETYPE, true, RMQ_DREXCHANGE, true, RMQ_DRQUEUE)

		messagesData := queue.Subscribe(1, false, RMQ_DRQUEUE, RMQ_DREXCHANGE, RMQ_DRQUEUE)

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		processor := NewProcessor(cfg, db, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {
				wg.Add(1)
				processor.DR(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)
			}
		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}
