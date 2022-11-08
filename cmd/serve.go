package cmd

import (
	"wiki/src/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve wiki on port 8080",
	Long:  "Serve locally stored wiki pages on http://localhost:8080/wiki/view",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}
