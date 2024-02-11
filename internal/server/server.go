package server

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"github.com/humanbelnik/ggrep/internal/lib/pretify"
)

type Repositories struct {
	Repos  []github.Repository
	Len    int
	Client *github.Client
}

// BrowseRepositories calls GitHub API
// to browse repositories
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
		Repos:  repositories,
		Len:    counter,
		Client: client,
	}, nil
}

func (r *Repositories) GetInfoAll() {
	for i, r := range r.Repos {
		fmt.Printf("%d. [%s]\n", i+1, pretify.WithColor(pretify.Cyan, r.GetFullName()))
		fmt.Printf("%s %s\n", pretify.WithColor(pretify.LightGray, "Link:"), r.GetHTMLURL())
		fmt.Printf("%s %s\n\n", pretify.WithColor(pretify.LightGray, "Language:"), pretify.LanguageWithColor(r.GetLanguage()))
	}
	fmt.Println("Total:", r.Len)

}

func (r *Repositories) GetInfoWithLanguage(lang string) {
	i := 0
	for _, r := range r.Repos {
		if strings.ToLower(r.GetLanguage()) != lang {
			continue
		}
		fmt.Printf("%d. [%s]\n", i+1, pretify.WithColor(pretify.Cyan, r.GetFullName()))
		fmt.Printf("%s %s\n", pretify.WithColor(pretify.LightGray, "Link:"), r.GetHTMLURL())
		fmt.Printf("%s %s\n\n", pretify.WithColor(pretify.LightGray, "Language:"), pretify.LanguageWithColor(r.GetLanguage()))
		i++
	}
	fmt.Println("Total:", i)
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
