package d09

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type visitedGrid [][]byte
type direction byte

const (
	up    direction = 'U'
	right direction = 'R'
	down  direction = 'D'
	left  direction = 'L'
)

var charToDir = map[byte]direction{
	'U': up,
	'R': right,
	'D': down,
	'L': left,
}

type square byte

const (
	visited  square = '#'
	unvisted square = '.'
)

type motion struct {
	dir direction
	mag int
}

type position struct {
	x int
	y int
}

type vector struct {
	x int
	y int
}

func soln(motions []motion, ropeLen int) int {
	row := []byte{byte(visited)}
	grid := [][]byte{row}

	knots := make([]position, ropeLen)
	for _, m := range motions {
		grid, knots = processMotion(grid, knots, m)
	}

	return calcScores(grid)
}

func P1(input []string) int {
	motions := buildMotions(input)
	return soln(motions, 2)
}

func P2(input []string) int {
	motions := buildMotions(input)
	return soln(motions, 10)
}

func calcScores(grid visitedGrid) int {
	score := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == '#' {
				score++
			}
		}
	}
	return score
}

func processMotion(grid visitedGrid, knots []position, move motion) (visitedGrid, []position) {
	grid, shift := expandIfNecessary(grid, knots[0], move) // may need to shift head
	newKnots := make([]position, len(knots))

	for i, knot := range knots {
		newKnots[i] = position{knot.x + shift.x, knot.y + shift.y}
	}

	for steps := move.mag; steps > 0; steps-- {
		newKnots[0] = moveHead(newKnots[0], move)

		for i := 1; i < len(newKnots); i++ {
			newKnots[i] = moveTail(newKnots[i-1], newKnots[i])
		}

		tail := newKnots[len(newKnots)-1]
		grid[tail.y][tail.x] = byte(visited)
	}

	return grid, newKnots
}

func moveTail(head, tail position) position {
	length := 1

	// move tail (if necessary) and mark square
	if withinOneSquare(head, tail) {
		return tail
	}

	// move right?
	if (head.x-tail.x) > length && head.y == tail.y {
		return position{head.x - length, tail.y}
	}

	// move left?
	if (tail.x-head.x) > length && head.y == tail.y {
		return position{head.x + length, tail.y}
	}

	// move up?
	if (tail.y-head.y) > length && head.x == tail.x {
		return position{head.x, head.y + length}
	}

	// move down?
	if (head.y-tail.y) > length && head.x == tail.x {
		return position{head.x, head.y - length}
	}

	// tail needs to move diagonally
	// move NE?
	if head.x > tail.x && head.y < tail.y {
		return position{tail.x + length, tail.y - length}
	}

	// move SE?
	if head.x > tail.x && head.y > tail.y {
		return position{tail.x + length, tail.y + length}
	}

	// move NW?
	if head.x < tail.x && head.y < tail.y {
		return position{tail.x - length, tail.y - length}
	}

	// move SW?
	if head.x < tail.x && head.y > tail.y {
		return position{tail.x - length, tail.y + length}
	}

	return tail
}

func moveHead(head position, move motion) position {
	newHead := position{head.x, head.y}

	// move head
	if move.dir == up {
		newHead.y--
	} else if move.dir == down {
		newHead.y++
	} else if move.dir == right {
		newHead.x++
	} else {
		newHead.x--
	}

	return newHead
}

func withinOneSquare(a, b position) bool {
	hor := math.Abs(float64(a.x) - float64(b.x))
	ver := math.Abs(float64(a.y) - float64(b.y))
	return hor <= 1.0 && ver <= 1.0
}

func printGrid(grid visitedGrid) {
	for _, row := range grid {
		fmt.Printf("%q\n", row)
	}
}

func buildMotions(input []string) []motion {
	output := make([]motion, len(input))

	for i, line := range input {
		substrs := strings.Split(line, " ")
		dir := charToDir[substrs[0][0]]
		mag, _ := strconv.Atoi(substrs[1])
		output[i] = motion{dir, mag}
	}

	return output
}

func expandIfNecessary(grid visitedGrid, head position, m motion) (visitedGrid, vector) {
	vec := vector{0, 0}
	target := position{head.x, head.y}

	if m.dir == up {
		target.y = head.y - m.mag
	} else if m.dir == down {
		target.y = head.y + m.mag
	} else if m.dir == left {
		target.x = head.x - m.mag
	} else {
		target.x = head.x + m.mag
	}

	for target.y < 0 {
		var ct int
		grid, ct = expandGridVertical(grid, false)
		target.y += ct
		vec.y += ct
	}

	for target.y >= len(grid) {
		grid, _ = expandGridVertical(grid, true)
	}

	for target.x < 0 {
		var ct int
		grid, ct = expandGridHorizontal(grid, false)
		target.x += ct
		vec.x += ct
	}

	for target.x >= len(grid[0]) {
		grid, _ = expandGridHorizontal(grid, true)
	}

	return grid, vec
}

// south = "should the new rows be appended south of existing rows?"
// returns # of new rows
func expandGridVertical(grid visitedGrid, south bool) (visitedGrid, int) {
	newRows := visitedGrid{}

	for i := 0; i < len(grid); i++ {
		row := make([]byte, len(grid[0]))
		for j := range row {
			row[j] = byte(unvisted)
		}

		newRows = append(newRows, row)
	}

	if south {
		return append(grid, newRows...), len(newRows)
	}
	return append(newRows, grid...), len(newRows)
}

// east = "should the new cols be appended east of existing cols?"
func expandGridHorizontal(grid visitedGrid, east bool) (visitedGrid, int) {
	newColCt := len(grid[0])

	for i := range grid {
		newCol := make([]byte, len(grid[i]))
		for j := range newCol {
			newCol[j] = byte(unvisted)
		}

		if east {
			grid[i] = append(grid[i], newCol...)
		} else {
			grid[i] = append(newCol, grid[i]...)
		}
	}

	return grid, newColCt
}
