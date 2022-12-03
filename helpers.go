package main

func sumIntSlice(l []int) int {
	total := 0

	for _, val := range l {
		total += val
	}

	return total
}

// slice is unchanged when val is smaller than all elements
func insertIntoFixedSortedSlice(newVal int, slice []int) {
	for i, curVal := range slice {
		if curVal >= newVal {
			continue
		}

		slice[i] = newVal
		newVal = curVal
	}
}
