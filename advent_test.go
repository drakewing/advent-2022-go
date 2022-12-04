package main

import "testing"

func TestAnswers(t *testing.T) {
	assertInts(t, d01p1(readLines("input/day01.txt")), 66616)
	assertInts(t, d01p2(readLines("input/day01.txt")), 199172)
	assertInts(t, d02p1(readLines("input/day02.txt")), 10624)
	assertInts(t, d02p2(readLines("input/day02.txt")), 14060)
	assertInts(t, d03p1(readLines("input/day03.txt")), 7863)
	assertInts(t, d03p2(readLines("input/day03.txt")), 2488)
}

func assertInts(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
