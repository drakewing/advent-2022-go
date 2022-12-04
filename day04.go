package main

import (
	"strconv"
	"strings"
)

type Pair struct {
	left  Interval
	right Interval
}

func d04p1(input []string) int {
	pairs := convLinesToPairs(input)

	contained := 0
	for _, pair := range pairs {
		if aContainsB(pair.left, pair.right) || aContainsB(pair.right, pair.left) {
			contained++
		}
	}

	return contained
}

func aContainsB(a, b Interval) bool {
	return a.lower <= b.lower && a.upper >= b.upper
}

func convLinesToPairs(input []string) []Pair {
	output := make([]Pair, 0)

	for _, line := range input {
		subs := strings.Split(line, ",")
		l := convStrToRange(subs[0])
		r := convStrToRange(subs[1])
		output = append(output, Pair{l, r})
	}

	return output
}

func convStrToRange(input string) Interval {
	subs := strings.Split(input, "-")
	l, _ := strconv.Atoi(subs[0])
	r, _ := strconv.Atoi(subs[1])
	return Interval{l, r}
}
