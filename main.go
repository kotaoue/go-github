package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v41/github"
)

func main() {
	client := github.NewClient(nil)
	listByOrg(client)
	publicList(client)
	privateList(client)
}

func listByOrg(client *github.Client) {
	fmt.Println("----org----")
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)

	if err != nil {
		fmt.Println(err)
	}

	printRepository(repos)
}

func list(client *github.Client, opt *github.RepositoryListOptions) {
	repos, _, err := client.Repositories.List(context.Background(), "kotaoue", opt)

	if err != nil {
		fmt.Println(err)
	}

	printRepository(repos)
}

func publicList(client *github.Client) {
	fmt.Println("----public----")
	list(client, &github.RepositoryListOptions{Type: "public"})
}

func privateList(client *github.Client) {
	fmt.Println("----private----")
	fmt.Println("----When i don't have permittion,  printing public repository only----")
	list(client, &github.RepositoryListOptions{Type: "private"})
}

func printRepository(repos []*github.Repository) {
	for _, v := range repos {
		fmt.Printf("Name:%s GitURL:%s\n", *v.Name, *v.GitURL)
	}
}
