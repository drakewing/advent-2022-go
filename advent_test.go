package main

import "testing"

func TestAnswers(t *testing.T) {
	assertInts(t, d01p1(), 66616)
	assertInts(t, d01p2(), 199172)
	assertInts(t, d02p1(), 10624)
	assertInts(t, d02p2(), 14060)
}

func assertInts(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
