package main

import (
	"bufio"
	"os"
	"strings"
)

const inputFile = "input/day02.txt"

type RpsMove int

const (
	Rock RpsMove = iota
	Paper
	Scissors
)

type RpsResult int

const (
	Player1 RpsResult = iota
	Player2
	Draw
)

type RpsGame struct {
	p1Move RpsMove
	p2Move RpsMove
}

func d02p1() int {
	games := buildGames(true)
	return scoreGames(games)
}

func d02p2() int {
	games := buildGames(false)
	return scoreGames(games)
}

func scoreGames(games []RpsGame) int {
	score := 0

	for _, game := range games {
		score += scoreGame(game)
	}

	return score
}

// scores a game from my perspective
func scoreGame(r RpsGame) int {
	score := 0

	switch r.p2Move {
	case Rock:
		score += 1
	case Paper:
		score += 2
	case Scissors:
		score += 3
	}

	switch winnerMap[r.p1Move][r.p2Move] {
	case Player2:
		score += 6
	case Draw:
		score += 3
	}

	return score
}

func buildGames(isPart1 bool) []RpsGame {
	var games []RpsGame

	f, _ := os.Open(inputFile)
	s := bufio.NewScanner(f)

	for s.Scan() {
		moves := strings.Split(s.Text(), " ")
		oppMove := strToMove(moves[0])

		var myMove RpsMove
		if isPart1 {
			myMove = strToMove(moves[1])
		} else {
			myMove = outcomeMap[oppMove][moves[1]]
		}
		game := RpsGame{oppMove, myMove}
		games = append(games, game)
	}

	return games
}

// maps p1 move -> p2 moves -> winner
var winnerMap = map[RpsMove]map[RpsMove]RpsResult{
	Rock: {
		Rock:     Draw,
		Paper:    Player2,
		Scissors: Player1,
	},
	Paper: {
		Rock:     Player1,
		Paper:    Draw,
		Scissors: Player2,
	},
	Scissors: {
		Rock:     Player2,
		Paper:    Player1,
		Scissors: Draw,
	},
}

// maps p1 move -> desired outcome -> required action
var outcomeMap = map[RpsMove]map[string]RpsMove{
	Rock: {
		"X": Scissors,
		"Y": Rock,
		"Z": Paper,
	},
	Paper: {
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	},
	Scissors: {
		"X": Paper,
		"Y": Scissors,
		"Z": Rock,
	},
}

// assumes valid input
func strToMove(s string) RpsMove {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	}

	return Scissors
}
