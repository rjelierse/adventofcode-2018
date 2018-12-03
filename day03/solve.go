package day03

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode-2018/input"
	"github.com/rjelierse/adventofcode-2018/puzzle"
	"regexp"
	"strconv"
)

type Coordinate struct {
	x int
	y int
}

var claimExpression = regexp.MustCompile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")

func parseClaim(definition string) *Claim {
	matches := claimExpression.FindStringSubmatch(definition)
	if matches == nil {
		fmt.Println("Invalid input", definition)
		return nil
	}
	id, _ := strconv.Atoi(matches[1])
	top, _ := strconv.Atoi(matches[3])
	left, _ := strconv.Atoi(matches[2])
	width, _ := strconv.Atoi(matches[4])
	height, _ := strconv.Atoi(matches[5])
	return &Claim{
		id: id,
		top: top,
		left: left,
		width: width,
		height: height,
	}
}

func parseInput(definitions []string) (claims []*Claim) {
	for _, definition := range definitions {
		claim := parseClaim(definition)
		claims = append(claims, claim)
	}
	return
}

func Command() subcommands.Command {
	return puzzle.NewPuzzle("day03", func(path string) error {
		inputs := input.ReadAsStrings(path)
		claims := parseInput(inputs)
		cloth := NewCloth()
		for _, claim := range claims {
			cloth.AddClaim(claim)
		}
		fmt.Println("Part 1:", cloth.CountOverlaps())
		for _, claim := range claims {
			if cloth.IsUnique(claim) {
				fmt.Println("Part 2:", claim.id)
				break
			}
		}
		return nil
	})
}
