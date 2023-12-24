package day5

import (
	"fmt"
	"regexp"
)

// var parseMaps = regexp.MustCompile(`^(\S*)-to-(\S+) map:((?:\n.+)+)$`)
var mapping = regexp.MustCompile(`seeds: ((?:.+)+)\n|(\S*)-to-(\S+) map:\n*((?:\n.+)+)`)

type Range struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

type Almanac struct {
	seeds    []int
	mappings map[string]Range
}

func buildAlmanac() Almanac {
	almanac := Almanac{}
	almanac.seeds = []int{}

	return almanac
}

func part1(input string) int {
	minPlot := 0

	// f, _ := os.ReadFile("input.txt")
	// // fmt.Println(string(f))
	m := mapping.FindAllStringSubmatch(input, -1)

	for idx := range m {
		if idx == 0 {
			fmt.Println("== seeds", m[idx][1])
			fmt.Println()
			continue
		}
		fmt.Println("==", m[idx][2], "to", m[idx][3], m[idx][4])
		fmt.Println()
	}

	return minPlot
}
