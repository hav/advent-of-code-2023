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
	count          int
}

var numbersRegex = regexp.MustCompile(`Card *(\d)+: *((?:[0-9]+ *)+) \| *((?:[0-9]+ *)+)`)

func createGame(inputLine string) *Game {
	numbers := numbersRegex.FindAllStringSubmatch(inputLine, -1)

	id := utils.ConvertToInt(numbers[0][1])
	game := Game{id: id, count: 1}

	for _, number := range strings.Fields(numbers[0][2]) {
		game.winningNumbers = append(game.winningNumbers, utils.ConvertToInt(number))
	}

	for _, number := range strings.Fields(numbers[0][3]) {
		game.numbers = append(game.numbers, utils.ConvertToInt(number))
	}
	return &game
}

func (g Game) countMatches() int {
	matches := 0

	for _, number := range g.numbers {
		if slices.Contains(g.winningNumbers, number) {
			matches += 1
		}
	}

	return matches
}

func (g Game) calculateWinnings() int {
	winnings := 0

	for i := 1; i <= g.countMatches(); i++ {
		if winnings == 0 {
			winnings += 1
		} else {
			winnings *= 2
		}
	}

	return winnings
}

func part1(input []string) int {
	winnings := 0
	games := []*Game{}
	for _, line := range input {
		games = append(games, createGame(line))
	}

	for _, game := range games {
		winnings += game.calculateWinnings()
	}

	return winnings
}

func part2(input []string) int {
	totalScratchcards := 0
	games := []*Game{}
	for _, line := range input {
		games = append(games, createGame(line))
	}

	for idx, game := range games {
		// fmt.Println("Adding", game.countMatches(), "for game", game.id, "to games after", idx)
		for i := 0; i < game.count; i++ {
			for i := 1; i <= game.countMatches(); i++ {
				index := idx + i
				if index > len(games)-1 {
					index = len(games) - 1
				}
				// fmt.Println("Adding", 1, "to game at index", index, games[index])

				games[index].count += 1
			}
		}
		// fmt.Println()
	}

	for _, game := range games {
		// fmt.Println(game)
		totalScratchcards += game.count
	}

	return totalScratchcards
}
