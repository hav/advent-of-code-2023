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
			desc:     "puzzle",
			input:    utils.ReadFileAsSlice("input.txt"),
			expected: 529618,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			partsSum := part1(tC.input)
			if partsSum != tC.expected {
				t.Errorf("was %d", partsSum)
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
			expected: 467835,
		},
		{
			desc:     "puzzle",
			input:    utils.ReadFileAsSlice("input.txt"),
			expected: 77509019,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			partsSum := part2(tC.input)
			if partsSum != tC.expected {
				t.Errorf("was %d", partsSum)
			}

		})
	}
}

// func TestPart2TestInput(t *testing.T) {
// 	input := []string{
// 		"467..114..",
// 		"...*......",
// 		"..35..633.",
// 		"......#...",
// 		"617*......",
// 		".....+.58.",
// 		"..592.....",
// 		"......755.",
// 		"...$.*....",
// 		".664.598.."}

// 	// input := []string{
// 	// 	"...*...",
// 	// 	"432.997"}

// 	partsSum := part2(input)

// 	if partsSum != 467835 {
// 		t.Errorf("was %d", partsSum)
// 	}
// }

// func TestPart2Input(t *testing.T) {
// 	input := utils.ReadFileAsSlice("input.txt")

// 	partsSum := part2(input)

// 	if partsSum != 70104392 {
// 		t.Errorf("was %d", partsSum)
// 	}
// }
