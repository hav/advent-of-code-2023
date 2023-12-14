package day2

import (
	"fmt"
	"regexp"
	"strings"

	"besharati.se/advent-of-code-2023/utils"
)

type Set struct {
	red   int
	green int
	blue  int
}

func (s Set) String() string {
	return fmt.Sprintf("Set: %d red, %d green, %d blue", s.red, s.green, s.blue)
}

type Game struct {
	id   int
	sets []Set
}

func (g *Game) String() string {
	return fmt.Sprintf("Game: %d, sets: %v", g.id, g.sets)
}

func (g Game) isPossible(limitSet Set) bool {
	isPossible := true
	for _, set := range g.sets {
		if set.red > limitSet.red || set.green > limitSet.green || set.blue > limitSet.blue {
			isPossible = false
		}
	}

	return isPossible
}

func (g Game) maxSet() Set {
	maxSet := Set{}
	for _, set := range g.sets {
		if set.red > maxSet.red {
			maxSet.red = set.red
		}
		if set.green > maxSet.green {
			maxSet.green = set.green
		}
		if set.blue > maxSet.blue {
			maxSet.blue = set.blue
		}
	}
	return maxSet
}

func (g Game) power() int {
	maxSet := g.maxSet()
	return maxSet.red * maxSet.green * maxSet.blue
}

var gameRe = regexp.MustCompile(`Game ([0-9]+): (.+)`)
var setRe = regexp.MustCompile(`(([0-9]+) ([bdeglnru]+))`)

func parseGames(gamesInput []string) []Game {
	var games []Game
	for _, gameInput := range gamesInput {
		split := gameRe.FindAllStringSubmatch(gameInput, -1)
		for i := range split {
			game := Game{id: utils.ConvertToInt(split[i][1])}

			sets := strings.Split(split[i][2], ";")
			for _, singleSet := range sets {
				set := parseSet(singleSet)
				game.sets = append(game.sets, set)
			}
			games = append(games, game)
		}
	}

	return games
}

func parseSet(singleSet string) Set {
	dices := setRe.FindAllStringSubmatch(singleSet, -1)
	set := Set{}
	for _, diceCountColor := range dices {
		switch diceCountColor[3] {
		case "red":
			set.red = utils.ConvertToInt(diceCountColor[2])
		case "green":
			set.green = utils.ConvertToInt(diceCountColor[2])
		case "blue":
			set.blue = utils.ConvertToInt(diceCountColor[2])
		}
	}
	return set
}

func part1(input []string, limit Set) int {
	games := parseGames(input)
	sumAll := 0
	for _, v := range games {
		if v.isPossible(limit) {
			sumAll += v.id
		}
	}
	return sumAll
}

func part2(input []string) int {
	games := parseGames(input)
	power := 0
	for _, v := range games {
		power += v.power()
	}

	return power
}
