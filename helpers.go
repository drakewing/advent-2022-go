package main

import (
	"bufio"
	"os"
)

func sumIntSlice(l []int) int {
	total := 0

	for _, val := range l {
		total += val
	}

	return total
}

func readLines(path string) []string {
	var output []string

	f, _ := os.Open(path)
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		test := s.Text()
		output = append(output, test)
	}

	return output
}

func intersectByte(a, b map[byte]bool) map[byte]bool {
	intersection := make(map[byte]bool)

	for k := range a {
		if b[k] {
			intersection[k] = true
		}
	}

	for k := range b {
		if a[k] {
			intersection[k] = true
		}
	}

	return intersection
}
