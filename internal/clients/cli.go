package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/humanbelnik/ggrep/internal/model"
)

// MustTokens parses non-flag CLI argumaments and
// returns error is there're no arguments.
// ! Must call it after ParseFlags().
func MustTokens() ([]string, error) {
	args := flag.Args()
	if len(args) == 0 {
		return args, fmt.Errorf("tokens not found")
	}

	return args, nil
}

// ParseFlags parses flags from CLI.
func ParseFlags() model.Flags {
	var (
		langFl string
		inspFl string
	)
	flag.StringVar(&langFl, "lang", "", "specify repo's most used language")
	flag.StringVar(&inspFl, "inspect", "", "specify how deep you can inspect repo's files")
	flag.Parse()

	return model.Flags{
		Language: strings.ToLower(langFl),
		Inspect:  strings.ToLower(inspFl),
	}
}
