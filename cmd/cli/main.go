package main

import (
	"fmt"

	cli "github.com/humanbelnik/ggrep/internal/clients"
	"github.com/humanbelnik/ggrep/internal/server"
)

func main() {
	// Parse flags.
	flags := cli.ParseFlags()

	// Parse tokens from CLI arguments.
	tokens, err := cli.MustTokens()
	if err != nil {
		panic(fmt.Errorf("search tokens unspecified: %w", err))
	}

	// Get repositories.
	repos, _ := server.BrowseRepositories(tokens)

	if flags.Language != "" {
		repos.GetInfoWithLanguage(flags.Language)
	} else {
		repos.GetInfoAll()
	}
}
