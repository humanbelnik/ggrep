package server

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

type Repositories struct {
	Repos []github.Repository
	Len   int
}

// BrowseRepositories calls GitHub API
// to concurrently browse repositories
// whose names contain all tokens.
// ! Function contains panic()
func BrowseRepositories(tokens []string) (*Repositories, error) {
	ctx := context.Background()
	client := github.NewClient(nil)

	var (
		repositories []github.Repository
		counter      int
	)

	tokenStr := getTokenStr(tokens)

	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 10000},
	}
	q := fmt.Sprintf("%s in:name", tokenStr)
	repos, _, err := client.Search.Repositories(ctx, q, opts)
	if err != nil {
		panic("search error")
	}

	for _, r := range repos.Repositories {
		name := r.GetName()
		if satisfyTokens(name, tokens) {
			repositories = append(repositories, r)
			counter++
		}
	}

	return &Repositories{
		Repos: repositories,
		Len:   counter,
	}, nil
}

func getTokenStr(tokens []string) string {
	var buf bytes.Buffer
	for _, t := range tokens {
		buf.WriteString(t)
		buf.WriteRune(' ')
	}

	return buf.String()
}

func satisfyTokens(n string, ts []string) bool {
	for _, t := range ts {
		if !strings.Contains(strings.ToLower(n), strings.ToLower(t)) {
			return false
		}
	}
	return true
}
