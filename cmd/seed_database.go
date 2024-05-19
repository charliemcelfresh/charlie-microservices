package cmd

import (
	"github.com/charliemcelfresh/charlie-microservices/internal/seed_database"
	"github.com/spf13/cobra"
)

var seedDatabaseCmd = &cobra.Command{
	Use: "seed-database",
	Run: func(cmd *cobra.Command, args []string) {
		seed_database.SeedDatabase()
	},
}

func init() {
	rootCmd.AddCommand(seedDatabaseCmd)
}
