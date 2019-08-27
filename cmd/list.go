package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/garj4/lend/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the amount each person owes in the terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := db.GetRecords()
		if err != nil {
			fmt.Printf("Error when reading rows from DB: %s", err)
			os.Exit(1)
		}
		var id int
		var event string
		var amount float64
		for rows.Next() {
			rows.Scan(&id, &event, &amount)
			fmt.Println(strconv.Itoa(id) + ": " + event + ": " + fmt.Sprintf("%f", amount))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
