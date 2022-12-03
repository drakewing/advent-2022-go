package main

import (
	"bufio"
	"os"
	"strconv"
)

type backpack struct {
	food []int
}

func d01p1() int {
	top1 := findTopN(1)
	return sumIntSlice(top1)
}

func d01p2() int {
	top3 := findTopN(3)
	return sumIntSlice(top3)
}

func findTopN(n int) []int {
	backpacks := buildBackpacks()
	maxCals := make([]int, n)

	for _, bp := range backpacks {
		calories := sumIntSlice(bp.food)
		if calories > maxCals[len(maxCals)-1] {
			insertIntoFixedSortedSlice(calories, maxCals)
		}
	}

	return maxCals
}

func buildBackpacks() []backpack {
	f, _ := os.Open("input/day01.txt")
	s := bufio.NewScanner(f)

	var backpacks []backpack
	curBackpack := backpack{}

	for s.Scan() {
		line := s.Text()

		if line == "" {
			backpacks = append(backpacks, curBackpack)
			curBackpack = backpack{}
		} else {
			calories, _ := strconv.Atoi(line)
			curBackpack.food = append(curBackpack.food, calories)
		}
	}

	return backpacks
}
