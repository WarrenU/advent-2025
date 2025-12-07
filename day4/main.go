package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// isCellAccessible returns true when the cell at r,c is an '@' and has
// fewer than 4 adjacent '@' neighbors (8-directional).
func isCellAccessible(stock [][]rune, r, c int) bool {
	if stock[r][c] != '@' {
		return false
	}

	rows := len(stock)
	// 8 directions
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	adj := 0
	for _, d := range dirs {
		nr := r + d[0]
		nc := c + d[1]

		if nr < 0 || nr >= rows {
			continue
		}
		// guard for ragged rows
		if nc < 0 || nc >= len(stock[nr]) {
			continue
		}
		if stock[nr][nc] == '@' {
			adj++
			// if already 4, we can stop counting
			if adj == 4 {
				return false
			}
		}
	}
	return true
}

// markAccessibleOnce finds all accessible '@' in the current grid and
// returns their coordinates. It does not mutate the grid.
func markAccessibleOnce(stock [][]rune) [][2]int {
	var marked [][2]int
	for r := range len(stock) {
		for c := range len(stock[r]) {
			if isCellAccessible(stock, r, c) {
				marked = append(marked, [2]int{r, c})
			}
		}
	}
	return marked
}

// applyIncrementalRemoval runs rounds until no more accessible '@' remain.
// After each round it applies all replacements and prints the grid.
// It returns the total removed.
func applyIncrementalRemoval(stock [][]rune) int {
	totalRemoved := 0
	round := 0

	// fmt.Println("Initial state:")
	// printGrid(stock)
	// fmt.Println()

	for {
		round++
		toRemove := markAccessibleOnce(stock)
		if len(toRemove) == 0 {
			break
		}

		// Apply all removals simultaneously
		for _, rc := range toRemove {
			r, c := rc[0], rc[1]
			stock[r][c] = 'x'
		}

		// fmt.Printf("Remove %d rolls of paper:\n", len(toRemove))
		// printGrid(stock)
		// fmt.Println()

		totalRemoved += len(toRemove)
		// continue to next round
		_ = round
	}

	return totalRemoved
}

func printGrid(stock [][]rune) {
	for _, row := range stock {
		fmt.Println(string(row))
	}
}

func readGridFromFile(path string) ([][]rune, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var stock [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		stock = append(stock, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return stock, nil
}

func main() {
	stock, err := readGridFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := applyIncrementalRemoval(stock)
	fmt.Printf("Total removed: %d\n", total)
}
