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

	f, _, _, err := client.Repositories.GetContents(context.Background(), "acim", "go-github", "test.txt", nil)
	if err != nil {
		panic(err)
	}

	c, err := f.GetContent()
	if err != nil {
		panic(err)
	}

	c += "test\n"
	o := &github.RepositoryContentFileOptions{
		Content: []byte(c),
	}
	_, _, err = client.Repositories.UpdateFile(context.Background(), "acim", "go-github", "test.txt", o)
	if err != nil {
		panic(err)
	}
}
