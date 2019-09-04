package cmd

import (
	"fmt"
	"os"

	"github.com/garj4/lend/db"
	"github.com/spf13/cobra"
)

// totalCmd represents the total command
var totalCmd = &cobra.Command{
	Use:   "total",
	Short: "Lists the amount the specified person owes in the terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Printf("Use the form `lend total <first name>`\n")
			os.Exit(1)
		}

		amountOwed, err := db.GetAmountOwed(args[0])
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		fmt.Printf("%.2f\n", amountOwed)
	},
}

func init() {
	rootCmd.AddCommand(totalCmd)
}
