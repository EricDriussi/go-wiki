package cmd

import (
	"fmt"
	"os"
	"wiki/pkg/server"

	"github.com/spf13/cobra"
)

var version bool

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Version number")
}

var rootCmd = &cobra.Command{
	Use:   "wiki",
	Short: "Silly implementation of a wiki",
	Long: `
This isn't even a pet project, it's just me messing around with go.
The cli interface is built with cobra (https://github.com/spf13/cobra)!
`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println("Wiki Version Number: Dunno, 0.1?")
		} else {
			fmt.Println("Launching server on port 8080")
			SetupDefaultArticles()
			// TODO.index page?
			fmt.Println("Go to http://localhost:8080/wiki/view/Wombat to check it out")
			server.Run()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
