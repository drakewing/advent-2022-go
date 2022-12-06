package d04

import (
	"strconv"
	"strings"

	"github.com/drakewing/advent-2022-go/helpers"
)

type intervalPair struct {
	left  helpers.Interval
	right helpers.Interval
}

type d4cond func(intervalPair) bool

func P1(input []string) int {
	pairs := convLinesToPairs(input)
	return soln(pairs, p1Condition)
}

func P2(input []string) int {
	pairs := convLinesToPairs(input)
	return soln(pairs, p2Condition)
}

func soln(pairs []intervalPair, cond d4cond) int {
	acc := 0
	for _, pair := range pairs {
		if cond(pair) {
			acc++
		}
	}
	return acc
}

func p1Condition(pair intervalPair) bool {
	return aContainsB(pair.left, pair.right) || aContainsB(pair.right, pair.left)
}

func p2Condition(pair intervalPair) bool {
	return hasOverlap(pair.left, pair.right)
}

func hasOverlap(a, b helpers.Interval) bool {
	return helpers.IsNumInInterval(a.Lower, b) ||
		helpers.IsNumInInterval(a.Upper, b) ||
		helpers.IsNumInInterval(b.Lower, a) ||
		helpers.IsNumInInterval(b.Upper, a)
}

func aContainsB(a, b helpers.Interval) bool {
	return a.Lower <= b.Lower && a.Upper >= b.Upper
}

func convLinesToPairs(input []string) []intervalPair {
	output := make([]intervalPair, 0)

	for _, line := range input {
		subs := strings.Split(line, ",")
		l := convStrToRange(subs[0])
		r := convStrToRange(subs[1])
		output = append(output, intervalPair{l, r})
	}

	return output
}

func convStrToRange(input string) helpers.Interval {
	subs := strings.Split(input, "-")
	l, _ := strconv.Atoi(subs[0])
	r, _ := strconv.Atoi(subs[1])
	return helpers.Interval{Lower: l, Upper: r}
}
