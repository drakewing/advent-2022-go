package main

import (
	"testing"

	"github.com/drakewing/advent-2022-go/d01"
	"github.com/drakewing/advent-2022-go/d02"
	"github.com/drakewing/advent-2022-go/d03"
	"github.com/drakewing/advent-2022-go/d04"
	"github.com/drakewing/advent-2022-go/d05"
	"github.com/drakewing/advent-2022-go/d06"
	"github.com/drakewing/advent-2022-go/d07"
	"github.com/drakewing/advent-2022-go/d09"
	"github.com/drakewing/advent-2022-go/d10"

	// "github.com/drakewing/advent-2022-go/d12"
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
	assertInts(t, d07.P1(helpers.ReadLines("input/day07.txt")), 1118405)
	assertInts(t, d07.P2(helpers.ReadLines("input/day07.txt")), 12545514)
	assertInts(t, d09.P1(helpers.ReadLines("input/day09.txt")), 5695)
	assertInts(t, d09.P2(helpers.ReadLines("input/day09.txt")), 2434)
	assertInts(t, d10.P1(helpers.ReadLines("input/day10.txt")), 14620)
	// assertInts(t, d11.P1(helpers.ReadLines("input/day11.txt")), 78678)
	// assertInts(t, d11.P2(helpers.ReadLines("input/day11.txt")), 15333249714)
	// assertInts(t, d12.P1(helpers.ReadLines("input/day12.txt")), 408)
	// assertInts(t, d12.P2(helpers.ReadLines("input/day12.txt")), 399)
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
