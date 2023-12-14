package day2

import (
	"testing"

	"besharati.se/advent-of-code-2023/utils"
)

func TestPart1TestInput(t *testing.T) {
	testInput := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}

	sumAll := part1(testInput, Set{red: 12, green: 13, blue: 14})

	if sumAll != 8 {
		t.Errorf("was %d", sumAll)
	}
}

func TestPart1Input(t *testing.T) {
	input := utils.ReadFileAsSlice("input.txt")

	sumAll := diceGame(input, Set{red: 12, green: 13, blue: 14})

	if sumAll != 0 {
		t.Errorf("was %d", sumAll)
	}
}

func TestPart2TestInput(t *testing.T) {
	
}
