package d10

import (
	"fmt"
	"strconv"
	"strings"
)

type opcode int

const (
	noop opcode = iota
	addx
)

type instruction struct {
	op      opcode
	operand int
}

func P1(input []string) int {
	instructions := buildInstructions(input)

	regx, sum := 1, 0
	cycle, wait, pc := 1, -1, 0

	for {
		// do we need to:
		// 		a) stop?
		if pc == len(instructions) {
			break
		}

		// 		b) wait?
		if instructions[pc].op == addx && wait == -1 {
			wait = 1
		}

		// check signal strength & advance cycle
		if cycle == 20 || (cycle-20)%40 == 0 {
			signal := regx * cycle
			sum += signal
		}
		cycle++

		if wait > 0 {
			wait--
			continue
		}

		// execute op & advance "pc"
		if instructions[pc].op == addx {
			regx += instructions[pc].operand
		}
		pc++
		wait = -1
	}

	return sum
}

func P2(input []string) int {
	instructions := buildInstructions(input)

	scanlines := make([][]rune, 0)
	regx, wait, pc := 1, -1, 0

	for i := 0; i < 6; i++ {
		scanline := make([]rune, 40)

		for j := 0; j < 40; j++ {
			if pixelInRange(j, regx) {
				scanline[j] = '#'
			} else {
				scanline[j] = '.'
			}

			if instructions[pc].op == addx && wait == -1 {
				wait = 1
			}

			if wait > 0 {
				wait--
				continue
			}

			// execute op & advance "pc"
			if instructions[pc].op == addx {
				regx += instructions[pc].operand
			}
			pc++
			wait = -1
		}

		scanlines = append(scanlines, scanline)
	}

	for _, scanline := range scanlines {
		fmt.Printf("%q\n", scanline)
	}
	return 0
}

func pixelInRange(pos, regx int) bool {
	return pos >= regx-1 && pos <= regx+1
}

func buildInstructions(input []string) []instruction {
	instructions := make([]instruction, 0)
	for _, line := range input {
		parts := strings.Split(line, " ")

		if parts[0] == "noop" {
			instructions = append(instructions, instruction{noop, 0})
		}

		if parts[0] == "addx" {
			operand, _ := strconv.Atoi(parts[1])
			instructions = append(instructions, instruction{addx, operand})
		}
	}

	return instructions
}
