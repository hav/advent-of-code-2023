package day1

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var charToIntMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
var re = regexp.MustCompile("[0-9]|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)")

func replaceAllStrings(input string, re *regexp.Regexp) string {
	var tmpString = input
	splitString := strings.Split(input, "")
	matches := re.FindAllStringIndex(tmpString, -1)

	for _, match := range matches {
		stringToReplace := strings.Join(splitString[match[0]:match[1]], "")
		tmpString = strings.ReplaceAll(tmpString, stringToReplace, charToIntMap[stringToReplace])

	}
	return tmpString
}

// func findFirstAndLastInt(input string) (int int) {
// 	firstInt, lastInt := 0, 0
// 	re := regexp.MustCompile(`\d+`)
// 	matches := re.FindAllString(input, -1)
// 	if len(matches) > 0 {
// 		firstInt, lastInt := convertToInt(matches[len(matches)-1]), convertToInt(matches[0])
// 	}
// 	return firstInt, lastInt
// }

type WordMatch struct {
	index int
	word  string
}

func sortWordMatches(wordSet map[WordMatch]bool) []WordMatch {
	words := []WordMatch{}
	for k := range wordSet {
		words = append(words, k)
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].index < words[j].index
	})

	return words
}

func calibrationValues2(input []string, partTwo bool) int {
	re := regexp.MustCompile("[0-9]")
	keys := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	totalSum := 0
	for _, line := range input {
		firstIntIndex := 1000000
		lastIntIndex := 0

		if intIndexes := re.FindAllStringIndex(line, -1); intIndexes != nil {
			firstIntIndex = intIndexes[0][0]
			lastIntIndex = intIndexes[len(intIndexes)-1][0]
		}

		wordSet := map[WordMatch]bool{}

		for _, v := range keys {
			firstWordIndex := strings.Index(input[0], v)
			lastWordIndex := strings.LastIndex(input[0], v)
			if firstWordIndex != -1 {
				wordSet[WordMatch{firstWordIndex, v}] = true
			}

			if lastWordIndex != -1 {
				wordSet[WordMatch{lastWordIndex, v}] = true
			}
		}

		words := sortWordMatches(wordSet)

		firstIndex := firstIntIndex
		lastIndex := lastIntIndex

		var firstInt, lastInt int

		if len(words) != 0 && words[0].index < firstIntIndex {
			firstInt = convertToInt(charToIntMap[words[0].word])
		} else {
			firstInt = convertToInt(line[firstIndex : firstIndex+1])
		}

		if len(words) != 0 && words[len(words)-1].index > lastIntIndex {
			lastInt = convertToInt(charToIntMap[words[len(words)-1].word])
		} else {
			lastInt = convertToInt(line[lastIndex : lastIndex+1])
		}

		// fmt.Println(firstInt*10 + lastInt)
		totalSum += firstInt*10 + lastInt
		// fmt.Println(firstInt, lastInt)
		// fmt.Printf("index first number: %d, index last number: %d\n", firstIndex, lastIndex)
		// fmt.Printf("first number: %s, last number: %s\n", line[firstIndex:firstIndex+1], charToIntMap[lastWord])

		// totalSum += convertToInt(firstInt + lastInt)
	}

	return totalSum
}

func calibrationValues(input []string, regex string, partTwo bool) int {
	totalSum := 0
	var nonNumericPattern = regexp.MustCompile(regex)
	for _, line := range input {
		var newLine = line
		if partTwo {
			newLine = replaceAllStrings(line, re)
		}

		r1 := nonNumericPattern.ReplaceAllString(newLine, "")
		if len(r1) == 1 {
			r1 += r1
		}

		r := strings.Split(r1, "")
		sum := convertToInt(r[0] + r[len(r)-1])

		totalSum += sum
		// fmt.Println(line, newLine, r1, sum, totalSum)
		if partTwo {

			fmt.Printf("original: %s, replaced: %s, stripped: %s, to add: %d, total: %d\n", line, newLine, r1, sum, totalSum)
		}
	}
	return totalSum
}

func convertToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
