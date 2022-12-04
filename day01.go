package main

import (
	"strconv"
)

type backpack struct {
	food []int
}

func d01p1(input []string) int {
	backpacks := buildBackpacks(input)
	top1 := findTopN(1, backpacks)
	return sumIntSlice(top1)
}

func d01p2(input []string) int {
	backpacks := buildBackpacks(input)
	top3 := findTopN(3, backpacks)
	return sumIntSlice(top3)
}

func findTopN(n int, backpacks []backpack) []int {
	maxCals := make([]int, n)

	for _, bp := range backpacks {
		calories := sumIntSlice(bp.food)
		if calories > maxCals[len(maxCals)-1] {
			insertIfLarger(calories, maxCals)
		}
	}

	return maxCals
}

func buildBackpacks(input []string) []backpack {
	var backpacks []backpack
	curBackpack := backpack{}

	for _, line := range input {
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

// slice is unchanged when val is smaller than all elements
func insertIfLarger(newVal int, slice []int) {
	for i, curVal := range slice {
		if curVal >= newVal {
			continue
		}

		slice[i] = newVal
		newVal = curVal
	}
}
