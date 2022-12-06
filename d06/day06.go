package d06

func P1(input string) int {
	return soln(4, input)
}

func P2(input string) int {
	return soln(14, input)
}

func soln(length int, input string) int {
	for l, r := 0, length; r <= len(input); l, r = l+1, r+1 {
		if isSet(input[l:r]) {
			return r
		}
	}
	return -1
}

// are all characters in this string unique?
func isSet(input string) bool {
	set := make(map[rune]bool)

	for _, b := range input {
		if set[b] {
			return false
		}
		set[b] = true
	}

	return true
}
