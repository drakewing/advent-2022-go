package d01

import (
	"strconv"

	"github.com/drakewing/advent-2022-go/helpers"
)

type backpack struct {
	food []int
}

func P1(input []string) int {
	backpacks := buildBackpacks(input)
	top1 := findTopN(1, backpacks)
	return helpers.SumIntSlice(top1)
}

func P2(input []string) int {
	backpacks := buildBackpacks(input)
	top3 := findTopN(3, backpacks)
	return helpers.SumIntSlice(top3)
}

func findTopN(n int, backpacks []backpack) []int {
	maxCals := make([]int, n)

	for _, bp := range backpacks {
		calories := helpers.SumIntSlice(bp.food)
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
