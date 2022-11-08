package cmd

import (
	"fmt"
	"wiki/src/retriever"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up wiki pages",
	Long:  "Fetch three wiki pages from Wikipedia",
	Run: func(cmd *cobra.Command, args []string) {
		SetupWiki()
	},
}

var articlesToDownload = []string{"Wombat", "Platypus", "TempleOS"}

func SetupWiki() {
	fmt.Println("Downloading a bunch of pages from wikipedia...")
	retriever.DownloadArticles(articlesToDownload)
	fmt.Println("All done!")
}
