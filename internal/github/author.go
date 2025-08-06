package github

import (
	"fmt"
	"log"

	"github.com/cli/go-gh/v2/pkg/api"
)

func GetAuthorInfo(username string) string {
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
