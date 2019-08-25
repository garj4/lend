package cmd

import (
	"database/sql"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// toCmd represents the to command
var toCmd = &cobra.Command{
	Use:   "to",
	Short: "The most basic command to record a transaction.",
	Long: `This command takes two arguments: a person's name and an amount.
  For example, "lend to Garrett 5" would record that $5 have been lent to Garrett`,
	Run: func(cmd *cobra.Command, args []string) {
		database, err := sql.Open("sqlite3", path.Join(configDir, "sqlite.db"))
		if err != nil {
			os.Exit(1)
		}

		statement, _ := database.Prepare("INSERT INTO people (firstname, lastname, amount) VALUES (?, ?, ?)")
		statement.Exec("Sample", "Person", "5.0")
	},
}

func init() {
	rootCmd.AddCommand(toCmd)
}
