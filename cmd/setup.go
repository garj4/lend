package cmd

import (
	"database/sql"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the config for the application.",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := sql.Open("sqlite3", path.Join(configDir, "sqlite.db"))
		if err != nil {
			os.Exit(1)
		}

		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, amount FLOAT)")
		_, err = statement.Exec()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
