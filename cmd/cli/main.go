package main

import (
	"fmt"

	cli "github.com/humanbelnik/ggrep/internal/clients"
	"github.com/humanbelnik/ggrep/internal/server"
)

func main() {
	// Parse flags.
	flags := cli.ParseFlags()
	fmt.Println(flags)

	// Parse tokens from CLI arguments.
	tokens, err := cli.MustTokens()
	if err != nil {
		panic(fmt.Errorf("search tokens unspecified: %w", err))
	}

	repos, _ := server.BrowseRepositories(tokens)
	for _, r := range repos.Repos {
		fmt.Println(r.GetFullName())
	}
	fmt.Println("found: ", repos.Len, "repositories")
}
