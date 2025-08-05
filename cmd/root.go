package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh-coauthor [username]",
	Short: "GH CLI extension to add co-authors to your commits.",
	Long:  "A GitHub CLI extension for adding co-authors to your commits.\n\nInteractively shows repository authors to choose from if no argument is provided.",
	Args:  cobra.MaximumNArgs(1),
	Example: `gh coauthor
gh coauthor johndoe`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(args[0])
		} else {
			fmt.Println("hi world, this is the gh-coauthor extension!")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
