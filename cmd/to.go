package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/garj4/lend/db"

	"github.com/spf13/cobra"
)

// toCmd represents the to command
var toCmd = &cobra.Command{
	Use:   "to",
	Short: "The most basic command to record a transaction.",
	Long: `This command takes two arguments: a person's name and an amount.
  For example, "lend to Garrett 5 Food" would record that $5 have been lent to Garrett for food.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Printf("Too few arguments. Use the form `lend to Garrett 5 Food`.\n")
			os.Exit(1)
		}

		amount, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Printf("Please provide a numerical amount.\n")
			os.Exit(1)
		}

		err = db.AddTransaction(args[2], args[0], amount)
		if err != nil {
			fmt.Printf("Failed to add record: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(toCmd)
}
