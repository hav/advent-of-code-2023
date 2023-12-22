package day3

import (
	"testing"

	"besharati.se/advent-of-code-2023/utils"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "test",
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598.."},
			expected: 4361,
		},
		{
			desc:     "puzzleInput",
			input:    utils.ReadFileAsSlice("input.txt"),
			expected: 529618,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			partsSum := part1(tC.input)
			if partsSum != tC.expected {
				t.Errorf("Error, was %d but should be %d", partsSum, tC.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "testInput",
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598.."},
			expected: 467835,
		},
		{
			desc:     "puzzleInput",
			input:    utils.ReadFileAsSlice("input.txt"),
			expected: 77509019,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			partsSum := part2(tC.input)
			if partsSum != tC.expected {
				t.Errorf("Error, was %d but should be %d", partsSum, tC.expected)
			}

		})
	}
}
