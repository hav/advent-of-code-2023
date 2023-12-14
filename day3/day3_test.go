package day3

import (
	"testing"

	"besharati.se/advent-of-code-2023/utils"
)

func TestPart1TestInput(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598.."}

	partsSum := part1(input)

	if partsSum != 4361 {
		t.Errorf("was %d", partsSum)
	}
}

func TestPart1Input(t *testing.T) {
	input := utils.ReadFileAsSlice("input.txt")
	partsSum := part1(input)

	if partsSum != 529618 {
		t.Errorf("was %d", partsSum)
	}
}
