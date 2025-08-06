package gitutil

import (
	"fmt"

	"github.com/go-git/go-git/v6"
)

func AddCoauthorToLastCommit(authorInfo string) {
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
