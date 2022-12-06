package main

import (
	"testing"

	"github.com/drakewing/advent-2022-go/d01"
	"github.com/drakewing/advent-2022-go/d02"
	"github.com/drakewing/advent-2022-go/d03"
	"github.com/drakewing/advent-2022-go/d04"
	"github.com/drakewing/advent-2022-go/d05"
	"github.com/drakewing/advent-2022-go/d06"
	"github.com/drakewing/advent-2022-go/helpers"
)

func TestAnswers(t *testing.T) {
	assertInts(t, d01.P1(helpers.ReadLines("input/day01.txt")), 66616)
	assertInts(t, d01.P2(helpers.ReadLines("input/day01.txt")), 199172)
	assertInts(t, d02.P1(helpers.ReadLines("input/day02.txt")), 10624)
	assertInts(t, d02.P2(helpers.ReadLines("input/day02.txt")), 14060)
	assertInts(t, d03.P1(helpers.ReadLines("input/day03.txt")), 7863)
	assertInts(t, d03.P2(helpers.ReadLines("input/day03.txt")), 2488)
	assertInts(t, d04.P1(helpers.ReadLines("input/day04.txt")), 507)
	assertInts(t, d04.P2(helpers.ReadLines("input/day04.txt")), 897)
	assertStrings(t, d05.P1(helpers.ReadLines("input/day05.txt")), "SPFMVDTZT")
	assertStrings(t, d05.P2(helpers.ReadLines("input/day05.txt")), "ZFSJBPRFP")
	assertInts(t, d06.P1(helpers.ReadLines("input/day06.txt")[0]), 1275)
	assertInts(t, d06.P2(helpers.ReadLines("input/day06.txt")[0]), 3605)
}

func assertInts(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
