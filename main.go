package main

import (
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gh-coauthor",
		Short: "A GitHub CLI extension for co-authors",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hi world, this is the gh-coauthor extension!")
			client, err := api.DefaultRESTClient()
			if err != nil {
				fmt.Println(err)
				return
			}
			response := struct{ Login string }{}
			err = client.Get("user", &response)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("running as %s\n", response.Login)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
