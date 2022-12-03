package main

func sumIntSlice(l []int) int {
	total := 0

	for _, val := range l {
		total += val
	}

	return total
}
