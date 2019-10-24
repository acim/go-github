package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		panic(err)
	}
	for _, r := range repos {
		fmt.Println(*r.Name)
	}

	a, _, _, err := client.Repositories.GetContents(context.Background(), "acim", "go-reflex", "Dockerfile", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(a.GetContent())

}
