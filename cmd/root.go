package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/garj4/lend/helpers"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lend",
	Short: "Track your outstanding balances from lending money to friends",
	Long: `Lender allows you to easily record the amount of money that changes hands among friends.
  Not everyone has the ability to make an immediate reimbursement, and sometimes money changes
  hands so frequently that it's simply not worth it to constantly make bank transfers.`,
}

// Execute is called from main.go in cobra and runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	helpers.ConfigDir = path.Join(home, ".lend")

	viper.AddConfigPath(helpers.ConfigDir)
	viper.SetConfigName("config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
