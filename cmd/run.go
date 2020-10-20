package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"vault.fritz.box/PI-ZERO-BLINKT-SERVER/internal/controller/fake"
	"vault.fritz.box/PI-ZERO-BLINKT-SERVER/internal/controller/rpi"
	"vault.fritz.box/PI-ZERO-BLINKT-SERVER/internal/delivery/http"
)

const (
	portConfigKey     = RootConfigKey + ".port"
	prodConfigKey     = RootConfigKey + ".prod"
	defaultServerPort = 8080
)

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"serve", "start"},
	Short:   "Starts the Server",
	Run: func(cmd *cobra.Command, args []string) {

		port := viper.GetInt(portConfigKey)
		prod := viper.GetBool(prodConfigKey)

		if prod {
			log.Infof("Start Production server at: %d\n", port)
			http.NewRouter(rpi.NewAlarmController(),port)
			return
		}

		log.Infof("Start DEV server at: %d\n", port)
		http.NewRouter(fake.NewAlarmController(),port)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().IntP("port", "p", defaultServerPort, "Port to run Application server on")
	runCmd.Flags().Bool("prod", false, "Defines if the RaspberryPi libraries are used, because this are just for the ARM arch available.")

	viper.BindPFlag(portConfigKey, runCmd.Flags().Lookup("port"))
	viper.BindPFlag(prodConfigKey, runCmd.Flags().Lookup("prod"))
}
