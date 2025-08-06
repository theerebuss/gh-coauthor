package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/theerebuss/gh-coauthor/internal/github"
	"github.com/theerebuss/gh-coauthor/internal/gitutil"
	"github.com/theerebuss/gh-coauthor/internal/prompt"
)

var rootCmd = &cobra.Command{
	Use:     "gh-coauthor [username]",
	Short:   "GH CLI extension to add co-authors to your commits.",
	Long:    "A GitHub CLI extension for adding co-authors to your commits.\n\nInteractively shows repository authors to choose from if no argument is provided.",
	Args:    cobra.MaximumNArgs(1),
	Example: `gh coauthor\ngh coauthor johndoe`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			return
		}

		var authorInfo string

		if len(args) == 0 {
			login := prompt.GetCollaboratorFromUser()
			authorInfo = github.GetAuthorInfo(login)
		} else {
			authorInfo = github.GetAuthorInfo(args[0])
		}

		gitutil.AddCoauthorToLastCommit(authorInfo)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
