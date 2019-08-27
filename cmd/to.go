package cmd

import (
	"fmt"

	"github.com/garj4/lend/db"

	"github.com/spf13/cobra"
)

// toCmd represents the to command
var toCmd = &cobra.Command{
	Use:   "to",
	Short: "The most basic command to record a transaction.",
	Long: `This command takes two arguments: a person's name and an amount.
  For example, "lend to Garrett 5" would record that $5 have been lent to Garrett`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.AddRecord("event", "name", -1.0)
		if err != nil {
			fmt.Printf("Failed to add record: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(toCmd)
}
