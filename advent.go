package main

import "fmt"

func main() {
	fmt.Printf("day 1, problem 1: %d\n", d01p1(readLines("input/day01.txt")))
	fmt.Printf("day 1, problem 2: %d\n", d01p2(readLines("input/day01.txt")))
	fmt.Printf("day 2, problem 1: %d\n", d02p1(readLines("input/day02.txt")))
	fmt.Printf("day 2, problem 2: %d\n", d02p2(readLines("input/day02.txt")))
	fmt.Printf("day 3, problem 1: %d\n", d03p1(readLines("input/day03.txt")))
	fmt.Printf("day 3, problem 2: %d\n", d03p2(readLines("input/day03.txt")))
	fmt.Printf("day 4, problem 1: %d\n", d04p1(readLines("input/day04.txt")))
	// fmt.Printf("day 4, problem 2: %d\n", d03p2(readLines("input/day04.txt")))
}
