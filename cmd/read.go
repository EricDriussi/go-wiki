package cmd

import (
	"fmt"
	"log"
	"wiki/src/page"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a wiki page",
	Long:  "Print out a wiki page's body given a title",
	Run: func(cmd *cobra.Command, args []string) {
		PrintPage()
	},
}

func PrintPage() {
	fmt.Println("Write the title of the desired wiki page (no spaces allowed):")
	var title string
	fmt.Scanln(&title)

	pageToRead, LoadErr := page.Load(title)
	if LoadErr != nil {
		log.Fatal("[ERROR]: Couldn't load requested page")
	}
	fmt.Println("Page found:")
	fmt.Println(string(pageToRead.Body))
}
