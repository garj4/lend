package cmd

import (
	"fmt"

	"github.com/garj4/lend/db"

	"github.com/spf13/cobra"
)

// addPersonCmd represents the to command
var addPersonCmd = &cobra.Command{
	Use:   "addPerson",
	Short: "Adds a new person to the database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Printf("You must supply a first and last name for the new person.")
		}

		err := db.AddPerson(args[0], args[1])
		if err != nil {
			fmt.Printf("Failed to add new person: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addPersonCmd)
}
