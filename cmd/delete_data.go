package cmd

import (
	"github.com/charliemcelfresh/charlie-microservices/internal/delete_data"
	"github.com/spf13/cobra"
)

var deleteAllDataCmd = &cobra.Command{
	Use: "delete-all-data",
	Run: func(cmd *cobra.Command, args []string) {
		delete_data.DeleteAll()
	},
}

func init() {
	rootCmd.AddCommand(deleteAllDataCmd)
}
