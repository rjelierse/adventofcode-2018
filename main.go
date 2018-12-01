package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode-2018/day01"
)

func main() {
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.HelpCommand(), "")

	subcommands.Register(day01.Command(), "Puzzles")

	flag.Parse()

	ctx := context.Background()
	subcommands.Execute(ctx)
}
