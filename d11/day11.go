package d11

import (
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items       []int
	operation   func(old int) int
	test        func(old int) bool
	testDivisor int
	trueTarget  int
	falseTarget int
	inspected   int
}

func buildMonkey(input []string) Monkey {
	m := Monkey{}

	items := strings.Split(input[0], " ")[4:]
	for _, item := range items {
		if item[len(item)-1] == ',' {
			item = item[:len(item)-1]
		}

		i, _ := strconv.Atoi(item)
		m.items = append(m.items, i)
	}

	funcParts := strings.Split(input[1], " ")[5:]
	m.operation = buildFuncFromParts(funcParts)

	testParts := strings.Split(input[2], " ")
	divisor, _ := strconv.Atoi(testParts[len(testParts)-1])
	m.test = func(dividend int) bool { return dividend%divisor == 0 }
	m.testDivisor = divisor

	trueTargetParts := strings.Split(input[3], " ")
	trueTarget, _ := strconv.Atoi(trueTargetParts[len(trueTargetParts)-1])
	m.trueTarget = trueTarget

	falseTargetParts := strings.Split(input[4], " ")
	falseTarget, _ := strconv.Atoi(falseTargetParts[len(falseTargetParts)-1])
	m.falseTarget = falseTarget

	return m
}

func buildMonkeys(input []string) []Monkey {
	output := make([]Monkey, 0)

	for i := 1; i < len(input); i += 7 {
		output = append(output, buildMonkey(input[i:i+5]))
	}

	return output
}

func buildFuncFromParts(parts []string) func(old int) int {
	if parts[2] == "old" {
		return func(old int) int {
			return old * old
		}
	}

	a, _ := strconv.Atoi(parts[2])
	switch parts[1] {
	case "+":
		return func(old int) int { return old + a }
	default:
		return func(old int) int { return old * a }
	}
}

func processRound(monkeys []Monkey, update func(old int) int) {
	for i := 0; i < len(monkeys); i++ {
		for len(monkeys[i].items) > 0 {
			item := monkeys[i].items[0]
			monkeys[i].items = monkeys[i].items[1:]
			monkeys[i].inspected++

			item = monkeys[i].operation(item)
			item = update(item)

			target := monkeys[i].falseTarget
			if monkeys[i].test(item) {
				target = monkeys[i].trueTarget
			}
			monkeys[target].items = append(monkeys[target].items, item)
		}
	}
}

func lcmOfPrimes(nums []int) int {
	lcm := 1
	for _, n := range nums {
		lcm *= n
	}
	return lcm
}

func P1(input []string) int {
	monkeys := buildMonkeys(input)

	for i := 0; i < 20; i++ {
		processRound(monkeys, func(item int) int { return item / 3 })
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})

	return monkeys[0].inspected * monkeys[1].inspected
}

func P2(input []string) int {
	monkeys := buildMonkeys(input)
	divisors := make([]int, 0)

	for _, monkey := range monkeys {
		divisors = append(divisors, monkey.testDivisor)
	}

	lcm := lcmOfPrimes(divisors)

	for i := 0; i < 10000; i++ {
		processRound(monkeys, func(item int) int { return item % lcm })
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})

	return monkeys[0].inspected * monkeys[1].inspected
}
