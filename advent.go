package main

import (
	"fmt"

	"github.com/drakewing/advent-2022-go/d01"
	"github.com/drakewing/advent-2022-go/d02"
	"github.com/drakewing/advent-2022-go/d03"
	"github.com/drakewing/advent-2022-go/d04"
	"github.com/drakewing/advent-2022-go/d05"
	"github.com/drakewing/advent-2022-go/d06"
	"github.com/drakewing/advent-2022-go/d07"
	"github.com/drakewing/advent-2022-go/d08"
	"github.com/drakewing/advent-2022-go/d09"
	"github.com/drakewing/advent-2022-go/d10"
	"github.com/drakewing/advent-2022-go/d11"
	"github.com/drakewing/advent-2022-go/helpers"
)

func main() {
	fmt.Printf("day 01, problem 1: %d\n", d01.P1(helpers.ReadLines("input/day01.txt")))
	fmt.Printf("day 01, problem 2: %d\n", d01.P2(helpers.ReadLines("input/day01.txt")))
	fmt.Printf("day 02, problem 1: %d\n", d02.P1(helpers.ReadLines("input/day02.txt")))
	fmt.Printf("day 02, problem 2: %d\n", d02.P2(helpers.ReadLines("input/day02.txt")))
	fmt.Printf("day 03, problem 1: %d\n", d03.P1(helpers.ReadLines("input/day03.txt")))
	fmt.Printf("day 03, problem 2: %d\n", d03.P2(helpers.ReadLines("input/day03.txt")))
	fmt.Printf("day 04, problem 1: %d\n", d04.P1(helpers.ReadLines("input/day04.txt")))
	fmt.Printf("day 04, problem 2: %d\n", d04.P2(helpers.ReadLines("input/day04.txt")))
	fmt.Printf("day 05, problem 1: %s\n", d05.P1(helpers.ReadLines("input/day05.txt")))
	fmt.Printf("day 05, problem 2: %s\n", d05.P2(helpers.ReadLines("input/day05.txt")))
	fmt.Printf("day 06, problem 1: %d\n", d06.P1(helpers.ReadLines("input/day06.txt")[0]))
	fmt.Printf("day 06, problem 2: %d\n", d06.P2(helpers.ReadLines("input/day06.txt")[0]))
	fmt.Printf("day 07, problem 1: %d\n", d07.P1(helpers.ReadLines("input/day07.txt")))
	fmt.Printf("day 07, problem 2: %d\n", d07.P2(helpers.ReadLines("input/day07.txt")))
	fmt.Printf("day 08, problem 1: %d\n", d08.P1(helpers.ReadLines("input/day08.txt")))
	fmt.Printf("day 08, problem 2: %d\n", d08.P2(helpers.ReadLines("input/day08.txt")))
	fmt.Printf("day 09, problem 1: %d\n", d09.P1(helpers.ReadLines("input/day09.txt")))
	fmt.Printf("day 09, problem 2: %d\n", d09.P2(helpers.ReadLines("input/day09.txt")))
	fmt.Printf("day 10, problem 1: %d\n", d10.P1(helpers.ReadLines("input/day10.txt")))
	fmt.Printf("day 10, problem 2:\n")
	d10.P2(helpers.ReadLines("input/day10.txt"))
	fmt.Printf("day 11, problem 1: %d\n", d11.P1(helpers.ReadLines("input/day11.txt")))
	fmt.Printf("day 11, problem 2: %d\n", d11.P2(helpers.ReadLines("input/day11.txt")))
	// fmt.Printf("day 12, problem 1: %d\n", d12.P1(helpers.ReadLines("input/day12.txt")))
	// fmt.Printf("day 12, problem 2: %d\n", d12.P2(helpers.ReadLines("input/day12.txt")))
}
