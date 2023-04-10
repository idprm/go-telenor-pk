package cmd

import "github.com/spf13/cobra"

const (
	RMQ_AUTOACK            = false
	RMQ_EXCHANGETYPE       = "direct"
	RMQ_EXCHANGEDURABILITY = false
	RMQ_QUEUEDURABILITY    = true
	RMQ_DATATYPE           = "application/json"
	RMQ_MOEXCHANGE         = "E_MO"
	RMQ_MOQUEUE            = "Q_MO"
	RMQ_DREXCHANGE         = "E_DR"
	RMQ_DRQUEUE            = "Q_DR"
	MT_INSUFFICIENT        = "MT_INSUFFICIENT"
	MT_FIRSTPUSH           = "MT_FIRSTPUSH"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long:  `Cobra is a CLI library for Go that empowers applications.`,
	}
)

func init() {
	/**
	 * WEBSERVER SERVICE
	 */
	rootCmd.AddCommand(serverCmd)

	/**
	 * CONSUMER MO & DR
	 */
	rootCmd.AddCommand(consumerMOCmd)
	rootCmd.AddCommand(consumerDRCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
