package example

import (
	"fmt"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initConfig()
	}
}

func Execute() {
	if err := runCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		if s != nil {
			s.Close()
		}
	}()
}

func init() {
	// cobra.OnInitialize(initConfig, initStore)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	runCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.otpservice.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		log.Fatal("Config file is not found")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Config file is not found")
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

