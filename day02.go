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

func d02p1() int {
	score := 0
	games := buildGames()

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

	switch r.determineWinner() {
	case Player2:
		score += 6
	case Draw:
		score += 3
	}

	return score
}

func buildGames() []RpsGame {
	var games []RpsGame

	f, _ := os.Open(inputFile)
	s := bufio.NewScanner(f)

	for s.Scan() {
		moves := strings.Split(s.Text(), " ")
		game := RpsGame{strToMove(moves[0]), strToMove(moves[1])}
		games = append(games, game)
	}

	return games
}

type RpsGame struct {
	p1Move RpsMove
	p2Move RpsMove
}

func (r RpsGame) determineWinner() RpsResult {
	switch r.p1Move {
	case Rock:
		switch r.p2Move {
		case Rock:
			return Draw
		case Paper:
			return Player2
		default:
			return Player1
		}
	case Paper:
		switch r.p2Move {
		case Rock:
			return Player1
		case Paper:
			return Draw
		default:
			return Player2
		}
	default:
		switch r.p2Move {
		case Rock:
			return Player2
		case Paper:
			return Player1
		default:
			return Draw
		}
	}
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
