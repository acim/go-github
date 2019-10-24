package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

const (
	owner    = "acim"
	repo     = "go-github"
	filepath = "counter.txt"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	f, _, _, err := client.Repositories.GetContents(context.TODO(), owner, repo, filepath, nil)
	if err != nil {
		panic(err)
	}

	c, err := f.GetContent()
	if err != nil {
		panic(err)
	}
	if c == "" {
		c = "0"
	}

	i, err := strconv.ParseInt(c, 10, 64)
	if err != nil {
		panic(err)
	}
	i++

	o := &github.RepositoryContentFileOptions{
		Content: []byte(strconv.FormatInt(i, 10)),
		SHA:     f.SHA,
		Message: github.String(fmt.Sprintf("increase counter to %d", i)),
	}
	_, _, err = client.Repositories.UpdateFile(context.TODO(), owner, repo, filepath, o)
	if err != nil {
		panic(err)
	}
}
