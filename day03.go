package main

type rucksack struct {
	compartments []compartment
}
type compartment map[byte]bool

func d03p1(input []string) int {
	rucksacks := buildRucks(input)
	intersections := findCompartmentIntersections(rucksacks)
	sumPriorities := 0

	for _, intersection := range intersections {
		for _, b := range intersection {
			sumPriorities += itemToPriority(b)
		}
	}

	return sumPriorities
}

func findCompartmentIntersections(rucks []rucksack) [][]byte {
	output := make([][]byte, 0)

	for i, ruck := range rucks {
		output = append(output, make([]byte, 0))
		intersection := intersectByte(ruck.compartments[0], ruck.compartments[1])
		for k := range intersection {
			output[i] = append(output[i], k)
		}
	}

	return output
}

func buildRucks(input []string) []rucksack {
	var rucksacks []rucksack

	for _, line := range input {
		newRuck := rucksack{}
		midpoint := len(line) / 2
		c1 := buildCompartment(line[:midpoint])
		c2 := buildCompartment(line[midpoint:])
		newRuck.compartments = append(newRuck.compartments, c1, c2)
		rucksacks = append(rucksacks, newRuck)
	}

	return rucksacks
}

func buildCompartment(input string) compartment {
	output := make(compartment)

	for i := 0; i < len(input); i++ {
		output[input[i]] = true
	}

	return output
}

func itemToPriority(item byte) int {
	if item >= 97 && item <= 122 {
		return int(item - 96)
	}

	if item >= 65 && item <= 90 {
		return int(item - 64 + 26)
	}

	panic("bad item value")
}
