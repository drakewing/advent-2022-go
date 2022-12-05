package main

import (
	"strconv"
	"strings"
)

type crate byte
type stack []byte
type move struct {
	ct   int
	src  int
	dest int
}

func d05p1(input []string) string {
	indexRow := findIndexRow(input)
	stacks := buildStacks(input[:indexRow+1])
	moves := buildMoves(input[indexRow+2:])
	rearrange(stacks, moves)
	return getTopCrates(stacks)
}

func rearrange(stacks []stack, moves []move) {
	for _, m := range moves {
		for i := 0; i < m.ct; i++ {
			lastPos := len(stacks[m.src]) - 1
			stacks[m.dest] = append(stacks[m.dest], stacks[m.src][lastPos])
			stacks[m.src] = stacks[m.src][:lastPos]
		}
	}
}

func getTopCrates(stacks []stack) string {
	s := strings.Builder{}

	for _, st := range stacks {
		s.WriteByte(st[len(st)-1])
	}

	return s.String()
}

func buildMoves(input []string) []move {
	moves := make([]move, 0)

	for _, line := range input {
		substrs := strings.Split(line, " ")
		ct, _ := strconv.Atoi(substrs[1])
		src, _ := strconv.Atoi(substrs[3])
		dest, _ := strconv.Atoi(substrs[5])
		moves = append(moves, move{ct, src - 1, dest - 1}) // 0-index stacks
	}

	return moves
}

func buildStacks(input []string) []stack {
	stackCt := findStackCt(input[len(input)-1])
	stacks := make([]stack, stackCt)

	// build stacks of crates. slice elements are ordered starting at the "bottom" of the stack.
	for i := len(input) - 2; i >= 0; i-- {
		for j, pos := 0, 1; pos < len(input[i]); j, pos = j+1, pos+4 {
			if input[i][pos] != ' ' {
				stacks[j] = append(stacks[j], input[i][pos])
			}
		}
	}

	return stacks
}

// finds row with stack indices, and assumes there is one.
func findIndexRow(input []string) int {
	for i, line := range input {
		if line[1] == '1' {
			return i
		}
	}

	panic("bad input for day 5")
}

func findStackCt(input string) int {
	stackCt, pos := 0, 1
	for ; pos < len(input); pos += 4 {
		stackCt++
	}
	return stackCt
}
