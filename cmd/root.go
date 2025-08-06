package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/go-git/go-git/v6"
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
		if len(args) > 1 {
			return
		}

		var authorInfo string

		if len(args) == 0 {
			login := getCollaboratorFromUser()
			authorInfo = getAuthorInfo(login)
		} else {
			authorInfo = getAuthorInfo(args[0])
		}

		addCoauthorToLastCommit(authorInfo)
	},
}

type Collaborator struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

func getCollaborators() ([]Collaborator, error) {
	output, _, err := gh.Exec("repo", "view", "--json", "mentionableUsers", "--jq", ".mentionableUsers")
	if err != nil {
		return nil, fmt.Errorf("error getting collaborators: %w", err)
	}

	var collaborators []Collaborator
	err = json.Unmarshal(output.Bytes(), &collaborators)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling collaborators: %w", err)
	}

	return collaborators, nil
}

func promptForCollaborator(collaborators []Collaborator) Collaborator {
	var collaborator Collaborator
	options := make([]huh.Option[Collaborator], len(collaborators))
	for i, c := range collaborators {
		label := c.Login
		if c.Name != "" {
			label = fmt.Sprintf("%s (%s)", c.Name, c.Login)
		}
		options[i] = huh.NewOption(label, c)
	}

	selection := huh.NewSelect[Collaborator]().
		Title("Choose the co-author").
		Options(options...).Value(&collaborator)

	form := huh.NewForm(huh.NewGroup(selection))
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return collaborator
}

func getCollaboratorFromUser() string {
	collaborators, err := getCollaborators()
	if err != nil {
		fmt.Println("Error getting collaborators:", err)
		return ""
	}

	collaborator := promptForCollaborator(collaborators)
	if collaborator.Login == "" {
		fmt.Println("No collaborator selected.")
		return ""
	}

	return collaborator.Login
}

func getAuthorInfo(username string) string {
	client, err := api.DefaultRESTClient()
	if err != nil {
		log.Fatalf("failed to create API client: %w", err)
	}

	res := struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Login string `json:"login"`
		Email string `json:"email"`
	}{}
	err = client.Get("users/"+username, &res)
	if err != nil {
		log.Fatalf("failed to get user info: %w", err)
	}

	var name string
	if res.Name == "" {
		name = res.Login
	} else {
		name = res.Name
	}

	var email string
	if res.Email == "" {
		email = fmt.Sprintf("%d+%s@users.noreply.github.com", res.Id, res.Login)
	} else {
		email = res.Email
	}

	return fmt.Sprintf("%s <%s>", name, email)
}

func addCoauthorToLastCommit(authorInfo string) {
	r, err := git.PlainOpen(".")
	if err != nil {
		fmt.Println("Error opening git repository:", err)
		return
	}

	w, err := r.Worktree()
	if err != nil {
		fmt.Println("Error getting worktree:", err)
		return
	}
	h, err := r.Head()
	if err != nil {
		fmt.Println("Error getting repository head:", err)
		return
	}

	c, err := r.CommitObject(h.Hash())
	if err != nil {
		fmt.Println("Error getting latest commit:", err)
		return
	}

	newMessage := fmt.Sprintf("%s\n\nCo-authored-by: %s", c.Message, authorInfo)
	w.Commit(newMessage, &git.CommitOptions{
		Amend: true,
	})

	fmt.Printf("Added co-author: %s\n", authorInfo)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
