package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

var (
	mode   string
	source string
	token  string
)

func init() {
	flag.StringVar(&mode, "mode", "public", "what type of list.[org|publicList|privateList]")
	flag.StringVar(&source, "source", "kotaoue", "datasource name")
	flag.StringVar(&token, "token", "", "GitHub Access Token. cf. https://github.com/settings/tokens")
	flag.Parse()
}

func main() {
	fmt.Printf("mode:%s\n", mode)

	client := newClient()
	switch mode {
	case "org":
		listByOrg(client)
	case "public":
		publicList(client)
	case "private":
		privateList(client)
	}
}

func newClient() *github.Client {
	switch mode {
	case "private":
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		return github.NewClient(oauth2.NewClient(ctx, ts))

	}

	return github.NewClient(nil)
}

func listByOrg(client *github.Client) {
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), source, opt)

	if err != nil {
		fmt.Println(err)
	}

	printRepository(repos)
}

func list(client *github.Client, opt *github.RepositoryListOptions) {
	repos, _, err := client.Repositories.List(context.Background(), source, opt)

	if err != nil {
		fmt.Println(err)
	}

	printRepository(repos)
}

func publicList(client *github.Client) {
	list(client, &github.RepositoryListOptions{Type: "public"})
}

func privateList(client *github.Client) {
	fmt.Println("----private----")
	fmt.Println("----When i don't have permittion, printing public repository only----")
	list(client, &github.RepositoryListOptions{Type: source})
}

func printRepository(repos []*github.Repository) {
	for _, v := range repos {
		fmt.Printf("Name:%s GitURL:%s\n", *v.Name, *v.GitURL)
	}
}
