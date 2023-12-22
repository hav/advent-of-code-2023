package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ConvertToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func SumIntSlice(slice []int) int {
	total := 0
	for i := 0; i < len(slice); i++ {
		total += slice[i]
	}
	return total
}

func ReadFile(fileName string) string {
	input, _ := os.ReadFile(fileName)
	return string(input)
}

func ReadFileAsSlice(fileName string) []string {
	input, _ := os.Open(fileName)
	defer input.Close()
	lines := []string{}
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
