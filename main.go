package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.HelpCommand(), "")

	flag.Parse()

	ctx := context.Background()
	subcommands.Execute(ctx)
}
