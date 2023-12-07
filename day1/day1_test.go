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

	// testInput := []string{"pqr3stu8vwx"}

	sumAll := calibrationValues2(testInput, false)

	if sumAll != 142 {
		t.Errorf("was %d", sumAll)
	}
}

func TestDay1Real(t *testing.T) {
	input := utils.ReadFileAsSlice("input.txt")

	// sumAll := calibrationValues(input, "[^0-9]", false)
	sumAll := calibrationValues2(input, false)

	if sumAll != 55621 {
		t.Errorf("was %d", sumAll)
	}
}

// func TestDay12(t *testing.T) {
// 	testInput := []string{"two1nine",
// 		"eightwothree",
// 		"abcone2threexyz",
// 		"xtwone3four",
// 		"4nineeightseven2",
// 		"zoneight234",
// 		"7pqrstsixteen",
// 		"vrrgrhnj1xgmcmd76two2oneightgqp"}

// 	var re = regexp.MustCompile("[0-9]")
// 	intIndexes := re.FindAllStringIndex(testInput[0], -1)

// 	firstIndex := 1000000
// 	lastIndex := 0
// 	lastWord := ""

// 	if firstIntIndex := intIndexes[0]; firstIntIndex != nil {
// 		firstIndex = intIndexes[0][0]
// 	}

// 	if lastIntIndex := intIndexes[len(intIndexes)-1]; lastIntIndex != nil {
// 		lastIndex = intIndexes[len(intIndexes)-1][0]
// 	}

// 	for _, v := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
// 		if index := strings.LastIndex(testInput[0], v); index > -1 {
// 			if index > lastIndex {
// 				lastIndex = index
// 				lastWord = v
// 			}
// 		}

// 		if index := strings.Index(testInput[0], v); index > -1 {
// 			if index < firstIndex {
// 				firstIndex = index
// 			}
// 		}
// 	}

// 	fmt.Printf("index first number: %d, index last number: %d\n", firstIndex, lastIndex)
// 	fmt.Printf("first number: %s, last number: %s\n", testInput[0][firstIndex:firstIndex+1], charToIntMap[lastWord])

// 	// sumAll := calibrationValues(testInput, "[^0-9]", true)

// 	// // if sumAll != 281 {
// 	// if sumAll != 18 {
// 	// 	t.Errorf("was %d", sumAll)
// 	// }

func TestDay1Part2Real(t *testing.T) {
	input := utils.ReadFileAsSlice("input.txt")

	sumAll := calibrationValues2(input, true)

	if sumAll != 53592 {
		t.Errorf("was %d", sumAll)
	}
}
