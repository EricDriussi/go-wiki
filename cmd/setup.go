package cmd

import (
	"fmt"
	"wiki/internal/wikipedia"

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
		if len(args) > 0 {
			SetupCustomArticles(args)
		} else {
			SetupDefaultArticles()
		}
	},
}

var defaultArticles = []string{"Wombat", "Platypus", "TempleOS"}

func SetupDefaultArticles() {
	fmt.Println("Downloading a bunch of pages from wikipedia...")
	wikipedia.DownloadArticles(defaultArticles)
	fmt.Println("All done!")
}

func SetupCustomArticles(art []string) {
	fmt.Println("Downloading the pages you requested from wikipedia...")
	wikipedia.DownloadArticles(art)
	fmt.Println("All done!")
}
