package day4

import (
	"regexp"
	"slices"
	"strings"

	"besharati.se/advent-of-code-2023/utils"
)

type Game struct {
	id             int
	winningNumbers []int
	numbers        []int
}

var numbersRegex = regexp.MustCompile(`Card *(\d)+: *((?:[0-9]+ *)+) \| *((?:[0-9]+ *)+)`)

func createGame(inputLine string) Game {
	numbers := numbersRegex.FindAllStringSubmatch(inputLine, -1)

	id := utils.ConvertToInt(numbers[0][1])
	game := Game{id: id}

	for _, number := range strings.Fields(numbers[0][2]) {
		game.winningNumbers = append(game.winningNumbers, utils.ConvertToInt(number))
	}

	for _, number := range strings.Fields(numbers[0][3]) {
		game.numbers = append(game.numbers, utils.ConvertToInt(number))
	}
	return game
}

func (g Game) calculateWinnings() int {
	winnings := 0

	for _, number := range g.numbers {
		if slices.Contains(g.winningNumbers, number) {
			if winnings == 0 {
				winnings += 1
			} else {
				winnings *= 2
			}
		}
	}

	return winnings
}

func part1(input []string) int {
	winnings := 0
	games := []Game{}
	for _, line := range input {
		games = append(games, createGame(line))
	}

	for _, game := range games {
		winnings += game.calculateWinnings()
	}

	return winnings
}
