package cli

import (
	"flag"
	"fmt"
)

type Flag struct {
	Language string
	Inspect  string
}

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
func ParseFlags() *Flag {
	var (
		langFl string
		inspFl string
	)
	flag.StringVar(&langFl, "lang", "", "specify repo's most used language")
	flag.StringVar(&inspFl, "inspect", "", "specify how deep you can inspect repo's files")
	flag.Parse()

	return &Flag{
		Language: langFl,
		Inspect:  inspFl,
	}
}
