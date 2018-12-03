package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode-2018/day01"
	"github.com/rjelierse/adventofcode-2018/day02"
	"github.com/rjelierse/adventofcode-2018/day03"
)

func main() {
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.HelpCommand(), "")

	subcommands.Register(day01.Command(), "Puzzles")
	subcommands.Register(day02.Command(), "Puzzles")
	subcommands.Register(day03.Command(), "Puzzles")

	flag.Parse()

	ctx := context.Background()
	subcommands.Execute(ctx)
}
