package day01

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode-2018/input"
	"github.com/rjelierse/adventofcode-2018/puzzle"
)

func part1(changes []int) int {
	frequency := 0
	for _, change := range changes {
		frequency = frequency + change
	}
	return frequency
}

func part2(changes []int) int {
	frequency := 0
	history := make(map[int]bool)
	for {
		for _, change := range changes {
			frequency = frequency + change
			if history[frequency] {
				return frequency
			}
			history[frequency] = true
		}
	}
}

func Command() subcommands.Command {
	p := puzzle.NewPuzzle("day01", func(path string) error {
		changes := input.ReadAsInts(path)
		fmt.Println("Part 1:", part1(changes))
		fmt.Println("Part 2:", part2(changes))
		return nil
	})

	return p
}
