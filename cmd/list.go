package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the amount each person owes in the terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := sql.Open("sqlite3", path.Join(configDir, "sqlite.db"))
		if err != nil {
			os.Exit(1)
		}

		rows, _ := database.Query("SELECT id, firstname, lastname, amount FROM people")
		var id int
		var firstname string
		var lastname string
		var amount float64
		for rows.Next() {
			rows.Scan(&id, &firstname, &lastname, &amount)
			fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname + ": " + fmt.Sprintf("%f", amount))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
