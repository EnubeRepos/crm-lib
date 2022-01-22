package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Short: "Application Worker",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logrus.Info("Version " + viper.GetString("version"))
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./configs/default.yaml)")
}

var configFile string

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("default")
		viper.AddConfigPath("./configs")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Fatal("Unable to read config from file")
	}
}

// Execute the application.
func Execute() {
	godotenv.Load()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
