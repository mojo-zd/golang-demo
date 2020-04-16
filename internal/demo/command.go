package demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const unexpectExitCode = 1

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Example: `hugo serve`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hugo start...")
	},
}

// Execute start cobra command
func Execute() {
	addChildren()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(unexpectExitCode)
	}
}

func addChildren() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "start hugo server",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("execute serve cmd of command!")
			cli := http.Client{}
			res, err := cli.Get("")
			defer func() {
				if err = res.Body.Close(); err != nil {
					fmt.Println("body close failed")
				}
			}()
			if err == nil {
				return
			}
			if _, err = ioutil.ReadAll(res.Body); err != nil {
				return
			}
		},
	})
}
