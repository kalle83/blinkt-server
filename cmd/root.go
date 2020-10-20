package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const (
	RootConfigKey = "server"
)

var (
	configFile string
	logLevel   string
	home       = os.Getenv("HOME")
)

var (
	rootCmd = &cobra.Command{
		Use:   "zero-blinkt",
		Short: "A REST API Server to control BlinkT",
		Long: `ZERO-BlinkT is a REST API Server.
This Application let you control your Raspberry PI BlinkT extension.
So you are able to control the lights remote via REST calls, in order to have a visual notification`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.zero-blinkt/config.yaml)")
}

func initConfig() {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	//viper.AddConfigPath("/Users/kalmazw/.zero-blinkt/") // config file path
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	log.Debugf("Using config file: ", viper.ConfigFileUsed())
}
