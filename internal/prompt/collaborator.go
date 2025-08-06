package prompt

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"

	"github.com/theerebuss/gh-coauthor/internal/github"
)

func promptForCollaborator(collaborators []github.Collaborator) github.Collaborator {
	var collaborator github.Collaborator
	options := make([]huh.Option[github.Collaborator], len(collaborators))
	for i, c := range collaborators {
		label := c.Login
		if c.Name != "" {
			label = fmt.Sprintf("%s (%s)", c.Name, c.Login)
		}
		options[i] = huh.NewOption(label, c)
	}

	selection := huh.NewSelect[github.Collaborator]().
		Title("Choose the co-author").
		Options(options...).Value(&collaborator)

	form := huh.NewForm(huh.NewGroup(selection))
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return collaborator
}

func GetCollaboratorFromUser() string {
	collaborators, err := github.GetCollaborators()
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
