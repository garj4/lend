package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
