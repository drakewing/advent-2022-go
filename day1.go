package main

import (
	"bufio"
	"os"
	"strconv"
)

type backpack struct {
	food []int
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

func d1p1() int {
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
