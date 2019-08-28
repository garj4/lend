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
	Short: "Lists the transactions.",
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := db.GetRecords()
		if err != nil {
			fmt.Printf("Error when reading rows from DB: %s", err)
			os.Exit(1)
		}
		var id, person int
		var event, date string
		var amount float64
		for rows.Next() {
			err := rows.Scan(&id, &event, &amount, &date, &person)
			if err != nil {
				fmt.Printf("Error when reading rows from DB: %s", err)
				os.Exit(1)
			}
			fmt.Println(strconv.Itoa(id) + ": " + event + " on date: " + date + " from " + strconv.Itoa(person) + ": " + fmt.Sprintf("%f", amount))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
