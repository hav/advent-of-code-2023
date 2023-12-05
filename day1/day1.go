package day1

import (
	"fmt"
	"regexp"
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
var re = regexp.MustCompile("(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)")

func replaceAllStrings(input string, re *regexp.Regexp) string {
	var tmpString = input
	splitString := strings.Split(input, "")
	matches := re.FindAllStringIndex(tmpString, 10)

	for _, match := range matches {
		stringToReplace := strings.Join(splitString[match[0]:match[1]], "")
		tmpString = strings.ReplaceAll(tmpString, stringToReplace, charToIntMap[stringToReplace])

	}
	return tmpString
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
