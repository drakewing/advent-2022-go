package d08

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func findVisibility(grid [][]int) [][]bool {
	visibilities := make([][]bool, 0)
	for _, row := range grid {
		visibilities = append(visibilities, findVisibilityHor(row))
	}

	// Iterate cols of the grid, scanning N->S and S->N for each col, calculating viz along the way.
	for col := 0; col < len(grid[0]); col++ {
		nmax, smax := -1, -1
		for n, s := 0, len(grid[col])-1; n < len(grid[col]); n, s = n+1, s-1 {
			visibilities[n][col] = visibilities[n][col] || grid[n][col] > nmax
			visibilities[s][col] = visibilities[s][col] || grid[s][col] > smax

			if grid[n][col] >= nmax {
				nmax = grid[n][col]
			}

			if grid[s][col] >= smax {
				smax = grid[s][col]
			}
		}
	}

	return visibilities
}

func findVisibilityHor(row []int) []bool {
	output := make([]bool, len(row))

	lmax, rmax := -1, -1
	for l, r := 0, len(row)-1; l < len(row); l, r = l+1, r-1 {
		output[l] = output[l] || row[l] > lmax
		output[r] = output[r] || row[r] > rmax

		if row[l] > lmax {
			lmax = row[l]
		}

		if row[r] > rmax {
			rmax = row[r]
		}
	}

	return output
}

func buildGrid(input []string) [][]int {
	grid := make([][]int, 0)

	for _, line := range input {
		cells := make([]int, 0)
		heights := strings.Split(line, "")

		for _, height := range heights {
			h, _ := strconv.Atoi(height)
			cells = append(cells, h)
		}

		grid = append(grid, cells)
	}

	return grid
}

// assume x & y are valid
func calcScenicScore(grid [][]int, x, y int) int {
	score := 1 // multiplicative identity

	// calculate north score
	acc := 0
	for i := y - 1; i >= 0; i-- {
		acc++
		if grid[i][x] >= grid[y][x] {
			break
		}
	}
	score *= acc

	// calculate south score
	acc = 0
	for i := y + 1; i < len(grid); i++ {
		acc++
		if grid[i][x] >= grid[y][x] {
			break
		}
	}
	score *= acc

	// calculate west score
	acc = 0
	for i := x - 1; i >= 0; i-- {
		acc++
		if grid[y][i] >= grid[y][x] {
			break
		}
	}
	score *= acc

	// calculate east score
	acc = 0
	for i := x + 1; i < len(grid[y]); i++ {
		acc++
		if grid[y][i] >= grid[y][x] {
			break
		}
	}
	score *= acc

	return score
}

func buildScenicGrid(grid [][]int) [][]int {
	scenicGrid := make([][]int, 0)

	for y := 0; y < len(grid); y++ {
		scenicGrid = append(scenicGrid, make([]int, 0))

		for x := 0; x < len(grid[y]); x++ {
			scenicGrid[y] = append(scenicGrid[y], calcScenicScore(grid, x, y))
		}
	}

	return scenicGrid
}

func findMaxCell(grid [][]int) int {
	largestCell := math.MinInt

	for _, row := range grid {
		for _, cell := range row {
			if cell > int(largestCell) {
				largestCell = cell
			}
		}
	}

	return largestCell
}

func countVis(grid [][]bool) int {
	sum := 0

	for _, row := range grid {
		for _, cell := range row {
			if cell {
				sum++
			}
		}
	}

	return sum
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func printVis(grid [][]bool) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func P1(input []string) int {
	grid := buildGrid(input)
	return countVis(findVisibility(grid))
}

func P2(input []string) int {
	grid := buildGrid(input)
	return findMaxCell(buildScenicGrid(grid))
}
