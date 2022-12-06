package d02

import (
	"strings"
)

type rpsMove int

const (
	Rock rpsMove = iota
	Paper
	Scissors
)

type rpsResult int

const (
	Player1 rpsResult = iota
	Player2
	Draw
)

type rpsGame struct {
	p1Move rpsMove
	p2Move rpsMove
}

func P1(input []string) int {
	games := buildGames(input, true)
	return scoreGames(games)
}

func P2(input []string) int {
	games := buildGames(input, false)
	return scoreGames(games)
}

func scoreGames(games []rpsGame) int {
	score := 0

	for _, game := range games {
		score += scoreGame(game)
	}

	return score
}

// scores a game from my perspective
func scoreGame(r rpsGame) int {
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

func buildGames(input []string, isPart1 bool) []rpsGame {
	var games []rpsGame

	for _, line := range input {
		moves := strings.Split(line, " ")
		oppMove := strToMove(moves[0])

		var myMove rpsMove
		if isPart1 {
			myMove = strToMove(moves[1])
		} else {
			myMove = outcomeMap[oppMove][moves[1]]
		}
		game := rpsGame{oppMove, myMove}
		games = append(games, game)
	}

	return games
}

// maps p1 move -> p2 moves -> winner
var winnerMap = map[rpsMove]map[rpsMove]rpsResult{
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
var outcomeMap = map[rpsMove]map[string]rpsMove{
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
func strToMove(s string) rpsMove {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	}

	return Scissors
}
