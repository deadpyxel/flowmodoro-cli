package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "flowmodoro",
	Short: "Flowmodoro is a CLI tool for the Flowmodoro technique",
	Long: `Flowmodoro is a CLI application to implement the Flowmodoro technique.
  A time management technique were you focus first and then define your break intervals`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// add config flag
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/flowmodoro-cli/config.yaml)")
}

func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search for config file in XDG config directory with name "flowmodoro-cli" (without extension).
		viper.AddConfigPath(home + "/.config/flowmodoro-cli")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config.yaml")
	}

	// Set default value for statePath. Will use your HOME if not set.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	viper.SetDefault("statePath", home+"/state.json")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
