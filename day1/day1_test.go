package day1

import (
	"testing"

	"besharati.se/advent-of-code-2023/utils"
)

func TestDay1(t *testing.T) {
	testInput := []string{"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"}

	sumAll := calibrationValues(testInput, "[^0-9]", false)

	if sumAll != 142 {
		t.Errorf("was %d", sumAll)
	}
}

func TestDay1Real(t *testing.T) {
	input := utils.ReadFileAsSlice("input.txt")

	sumAll := calibrationValues(input, "[^0-9]", false)

	if sumAll != 55621 {
		t.Errorf("was %d", sumAll)
	}
}

func TestDay12(t *testing.T) {
	testInput := []string{"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
		"eight"}

	sumAll := calibrationValues(testInput, "[^0-9]", true)

	if sumAll != 281 {
		t.Errorf("was %d", sumAll)
	}
}

func TestDay1Part2Real(t *testing.T) {
	input := utils.ReadFileAsSlice("input.txt")

	sumAll := calibrationValues(input, "[^0-9]", true)

	if sumAll != 53592 {
		t.Errorf("was %d", sumAll)
	}
}
