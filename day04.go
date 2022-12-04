package main

import (
	"strconv"
	"strings"
)

type Pair struct {
	left  Interval
	right Interval
}

type d4cond func(Pair) bool

func d04p1(input []string) int {
	pairs := convLinesToPairs(input)
	return d4soln(pairs, p1Condition)
}

func d04p2(input []string) int {
	pairs := convLinesToPairs(input)
	return d4soln(pairs, p2Condition)
}

func d4soln(pairs []Pair, cond d4cond) int {
	acc := 0
	for _, pair := range pairs {
		if cond(pair) {
			acc++
		}
	}
	return acc
}

func p1Condition(pair Pair) bool {
	return aContainsB(pair.left, pair.right) || aContainsB(pair.right, pair.left)
}

func p2Condition(pair Pair) bool {
	return hasOverlap(pair.left, pair.right)
}

func hasOverlap(a, b Interval) bool {
	return isNumInInterval(a.lower, b) ||
		isNumInInterval(a.upper, b) ||
		isNumInInterval(b.lower, a) ||
		isNumInInterval(b.upper, a)
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
