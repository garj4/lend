package cmd

import (
	"fmt"
	"os"

	"github.com/garj4/lend/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the transactions.",
	Run: func(cmd *cobra.Command, args []string) {
		peopleFlag, err := cmd.Flags().GetBool("people")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if peopleFlag {
			err = db.PrintPeople(os.Stdout)
		} else {
			err = db.PrintTransactions(os.Stdout)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("people", "p", false, "List people instead of transactions")
}
