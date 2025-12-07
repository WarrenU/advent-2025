package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isCellAccessible(stock [][]rune, r, c int) bool {
	if stock[r][c] != '@' {
		return false
	}

	rows := len(stock)
	cols := len(stock[0])

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	adj := 0

	for _, d := range dirs {
		nr := r + d[0]
		nc := c + d[1]

		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
			continue
		}

		if stock[nr][nc] == '@' {
			adj++
		}
	}

	return adj < 4
}

func CountAccessible(stock [][]rune) int {
	rows := len(stock)
	if rows == 0 {
		return 0
	}
	cols := len(stock[0])

	total := 0
	for r := range rows {
		for c := range cols {
			if isCellAccessible(stock, r, c) {
				total++
			}
		}
	}

	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stock := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		stockRow := []rune{}
		for _, r := range line {
			stockRow = append(stockRow, r)
		}
		stock = append(stock, stockRow)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	total := CountAccessible(stock)

	fmt.Printf("Total Rolls accessible is: %d\n", total)
}
