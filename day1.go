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
	backpacks := buildBackpacks()
	maxCals := -1

	for _, bp := range backpacks {
		calories := sumIntSlice(bp.food)
		if calories > maxCals {
			maxCals = calories
		}
	}

	return maxCals
}

func d01p2() int {
	backpacks := buildBackpacks()
	maxCals := [3]int{-1, -1, -1}

	for _, bp := range backpacks {
		calories := sumIntSlice(bp.food)
		if calories > maxCals[2] {
			insertIntoFixedSortedSlice(calories, maxCals[:])
		}
	}

	return sumIntSlice(maxCals[:])
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
