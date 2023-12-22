package day3

import (
	"fmt"
	"regexp"

	"besharati.se/advent-of-code-2023/utils"
)

var numberRegex = regexp.MustCompile(`([0-9]+)`)
var symbolRegex = regexp.MustCompile(`([^\.0-9])`)

func figureOutIndexes(inputIdx, inputLen int) []int {
	// figure out which lines to look for symbols in
	inputIdxsToCheck := []int{inputIdx - 1, inputIdx, inputIdx + 1}

	switch inputIdx {
	case 0:
		inputIdxsToCheck = inputIdxsToCheck[1:]
	case inputLen - 1:
		inputIdxsToCheck = inputIdxsToCheck[:len(inputIdxsToCheck)-1]
	}

	return inputIdxsToCheck
}

type Part struct {
	part       string
	firstIndex int
	lastIndex  int
}

func (p Part) String() string {
	return p.part
}

func (p Part) isPart(line string) (bool, string, []int) {
	if symbol := symbolRegex.FindAllStringIndex(line, -1); symbol != nil {
		for _, s := range symbol {
			firstIndex := p.firstIndex - 1
			if p.firstIndex <= 0 {
				firstIndex = 0
			}

			for i := firstIndex; i <= p.lastIndex; i++ {
				if s[0] == i || s[1]-1 == i {
					return true, line[s[0]:s[1]], s
				}
			}
		}
	}
	return false, "", nil
}

func part1(input []string) int {
	partsSum := 0

	for inputIdx, line := range input {
		potentialPartIdxs := numberRegex.FindAllStringIndex(line, -1)

		for _, partIdxs := range potentialPartIdxs {
			pp := Part{part: line[partIdxs[0]:partIdxs[1]], firstIndex: partIdxs[0], lastIndex: partIdxs[1]}
			inputIdxsToCheck := figureOutIndexes(inputIdx, len(input))

			for _, inputLineIndex := range inputIdxsToCheck {
				if ok, symbol, symbolIdx := pp.isPart(input[inputLineIndex]); ok {
					fmt.Printf("%s is a part number because we found %s at %v on line %d\n", pp.part, symbol, symbolIdx, inputLineIndex)
					partsSum += utils.ConvertToInt(pp.part)
				}
			}
		}
	}

	return partsSum
}

type Gear struct {
	firstIndex int
	lastIndex  int
	parts      map[Part]bool
}

func (g Gear) calculateRatio() int {
	if len(g.parts) == 2 {
		keys := make([]Part, 0, len(g.parts))
		for p := range g.parts {
			keys = append(keys, p)
		}

		ratio := utils.ConvertToInt(keys[0].part) * utils.ConvertToInt(keys[1].part)
		// fmt.Println("Gear", g, "has ratio", ratio)
		return ratio
	}

	return 0
}

func createGear(gear []int, line string) Gear {
	startIndex := gear[0] - 1
	lastIndex := gear[1]
	if gear[0] == 0 {
		startIndex = 0
	}

	if gear[1] == len(line)-1 {
		lastIndex = len(line)
	}
	return Gear{firstIndex: startIndex, lastIndex: lastIndex, parts: map[Part]bool{}}
}

func (g Gear) isGearPart(firstIndex, lastIndex int) bool {
	for i := g.firstIndex; i <= g.lastIndex; i++ {
		if firstIndex == i || lastIndex-1 == i {
			return true
		}
	}
	return false
}

func part2(input []string) int {
	symbolRegex := regexp.MustCompile(`([*])`)
	totalRatio := 0
	for inputIdx, line := range input {
		gears := symbolRegex.FindAllStringIndex(line, -1)
		for _, gear := range gears {
			g := createGear(gear, line)
			inputIdxsToCheck := figureOutIndexes(inputIdx, len(input))

			for _, inputLine := range inputIdxsToCheck {
				potentialParts := numberRegex.FindAllStringIndex(input[inputLine], -1)
				for _, part := range potentialParts {
					if g.isGearPart(part[0], part[1]) {
						pp := Part{part: input[inputLine][part[0]:part[1]], firstIndex: part[0], lastIndex: part[1]}
						g.parts[pp] = true
					}
				}
			}

			totalRatio += g.calculateRatio()
		}
	}
	return totalRatio
}
