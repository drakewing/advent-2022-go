package main

import (
	"fmt"

	"github.com/drakewing/advent-2022-go/d01"
	"github.com/drakewing/advent-2022-go/d02"
	"github.com/drakewing/advent-2022-go/d03"
	"github.com/drakewing/advent-2022-go/d04"
	"github.com/drakewing/advent-2022-go/d05"
	"github.com/drakewing/advent-2022-go/d06"
	"github.com/drakewing/advent-2022-go/d09"
	"github.com/drakewing/advent-2022-go/helpers"
)

func main() {
	fmt.Printf("day 1, problem 1: %d\n", d01.P1(helpers.ReadLines("input/day01.txt")))
	fmt.Printf("day 1, problem 2: %d\n", d01.P2(helpers.ReadLines("input/day01.txt")))
	fmt.Printf("day 2, problem 1: %d\n", d02.P1(helpers.ReadLines("input/day02.txt")))
	fmt.Printf("day 2, problem 2: %d\n", d02.P2(helpers.ReadLines("input/day02.txt")))
	fmt.Printf("day 3, problem 1: %d\n", d03.P1(helpers.ReadLines("input/day03.txt")))
	fmt.Printf("day 3, problem 2: %d\n", d03.P2(helpers.ReadLines("input/day03.txt")))
	fmt.Printf("day 4, problem 1: %d\n", d04.P1(helpers.ReadLines("input/day04.txt")))
	fmt.Printf("day 4, problem 2: %d\n", d04.P2(helpers.ReadLines("input/day04.txt")))
	fmt.Printf("day 5, problem 1: %s\n", d05.P1(helpers.ReadLines("input/day05.txt")))
	fmt.Printf("day 5, problem 2: %s\n", d05.P2(helpers.ReadLines("input/day05.txt")))
	fmt.Printf("day 6, problem 1: %d\n", d06.P1(helpers.ReadLines("input/day06.txt")[0]))
	fmt.Printf("day 6, problem 2: %d\n", d06.P2(helpers.ReadLines("input/day06.txt")[0]))
	fmt.Printf("day 9, problem 1: %d\n", d09.P1(helpers.ReadLines("input/day09.txt")))
	// fmt.Printf("day 9, problem 2: %d\n", d09.P2(helpers.ReadLines("input/day09.txt")))
}
