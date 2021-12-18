package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v41/github"
)

func main() {
	client := github.NewClient(nil)
	listByOrg(client)
	list(client)
	privateList(client)
}

func listByOrg(client *github.Client) {
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)

	if err != nil {
		fmt.Println(err)
	}

	printRepository(repos)
}

func list(client *github.Client) {
	opt := &github.RepositoryListOptions{Type: "public"}
	repos, _, err := client.Repositories.List(context.Background(), "kotaoue", opt)

	if err != nil {
		fmt.Println(err)
	}

	printRepository(repos)
}

func privateList(client *github.Client) {
	opt := &github.RepositoryListOptions{Type: "private"}
	repos, _, err := client.Repositories.List(context.Background(), "kotaoue", opt)

	if err != nil {
		fmt.Println(err)
	}

	printRepository(repos)
}

func printRepository(repos []*github.Repository) {
	for _, v := range repos {
		fmt.Printf("Name:%s GitURL:%s\n", *v.Name, *v.GitURL)
	}
}
