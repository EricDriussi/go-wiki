package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"wiki/pkg/page"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(writeCmd)
}

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write a wiki page",
	Long:  "Write a new entry for the Wiki!!",
	Run: func(cmd *cobra.Command, args []string) {
		CreatePage()
	},
}

func CreatePage() {
	fmt.Println("Write a title for the new wiki page (no spaces allowed):")
	var title string
	fmt.Scanln(&title)

	fmt.Println("Write the content of the new wiki page:")
	body := basicReader()

	pageToWrite := page.Page{Title: title, Body: body}
	saveErr := pageToWrite.Save()
	if saveErr != nil {
		log.Fatal("[ERROR]: Couldn't save requested page")
	}
}

func basicReader() string {
	reader := bufio.NewReader(os.Stdin)
	body, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input", err)
		return ""
	}
	return strings.TrimSuffix(body, "\n")
}
