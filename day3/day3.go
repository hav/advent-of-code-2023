package day3

import (
	"fmt"
	"regexp"

	"besharati.se/advent-of-code-2023/utils"
)

var numberRegex = regexp.MustCompile(`([0-9]+)`)
var symbolRegex = regexp.MustCompile(`([^\.0-9])`)

type Part struct {
	part       string
	firstIndex int
	lastIndex  int
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
			// fmt.Printf("Trying to identify if %s is a part number\n", pp.part)
			inputIdxsToCheck := []int{inputIdx - 1, inputIdx, inputIdx + 1}

			// figure out which lines to look for symbols in
			switch inputIdx {
			case 0:
				inputIdxsToCheck = inputIdxsToCheck[1:]
			case len(input) - 1:
				inputIdxsToCheck = inputIdxsToCheck[:len(inputIdxsToCheck)-1]
			}

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

		}
	}

	return partsSum
}
