package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v41/github"
)

func main() {
	client := github.NewClient(nil)

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range repos {
		fmt.Println(v)
	}
}
