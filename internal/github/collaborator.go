package github

import (
	"encoding/json"
	"fmt"

	"github.com/cli/go-gh"
)

type Collaborator struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

func GetCollaborators() ([]Collaborator, error) {
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
