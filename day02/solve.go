package day02

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode-2018/input"
	"github.com/rjelierse/adventofcode-2018/puzzle"
)

func Command() subcommands.Command {
	return puzzle.NewPuzzle("day02", func(path string) error {
		boxes := input.ReadAsStrings(path)
		fmt.Println("Part 1:", part1(boxes))
		fmt.Println("Part 2:", part2(boxes))
		return nil
	})
}
