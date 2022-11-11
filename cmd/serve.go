package cmd

import (
	"fmt"
	"wiki/pkg/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve wiki on port 8080",
	Long:  "Serve locally stored wiki pages on http://localhost:8080/wiki",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go to http://localhost:8080/wiki")
		server.Run()
	},
}
