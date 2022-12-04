package main

import (
	"bufio"
	"os"
)

type Interval struct {
	lower int
	upper int
}

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

func intersectByteSets(a map[byte]bool, rest ...map[byte]bool) map[byte]bool {
	intersection := make(map[byte]bool)

	for k := range a {
		for i, m := range rest {
			if !m[k] {
				break
			}

			// k has been in all maps
			if i == len(rest)-1 {
				intersection[k] = true
			}
		}
	}

	return intersection
}

func getSliceFromKeys(m map[byte]bool) []byte {
	output := make([]byte, 0)

	for k := range m {
		output = append(output, k)
	}

	return output
}
