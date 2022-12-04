package main

import "fmt"

func main() {
	fmt.Printf("day 1, problem 1: %d\n", d01p1())
	fmt.Printf("day 1, problem 2: %d\n", d01p2())
	fmt.Printf("day 2, problem 1: %d\n", d02p1())
	fmt.Printf("day 2, problem 2: %d\n", d02p2())
	fmt.Printf("day 3, problem 1: %d\n", d03p1(readLines("input/day03.txt")))
	fmt.Printf("day 3, problem 2: %d\n", d03p2(readLines("input/day03.txt")))
}
