package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows := [][]int{}
	arithmetic := []string{}
	maxCols := 0
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			// skip empty lines
			continue
		}

		fields := strings.Fields(line)
		row := make([]int, len(fields))
		for i, fld := range fields {
			if fld != "*" && fld != "+" {
				n, err := strconv.Atoi(fld)
				if err != nil {
					log.Fatalf("failed to parse %q: %s", fld, err)
				}
				row[i] = n
			} else {
				arithmetic = append(arithmetic, fld)
			}
		}

		if len(row) > maxCols {
			maxCols = len(row)
		}
		rows = append(rows, row)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	if len(rows) == 0 || maxCols == 0 {
		fmt.Println("no numeric rows found")
		return
	}

	results := make([]int, maxCols)
	totalRows := len(rows) - 1
	for c := 0; c < maxCols; c++ {
		// start with first row value if present, otherwise zero
		for i := range totalRows {
			switch arithmetic[c] {
			case "*":
				if results[c] == 0 {
					results[c] = 1
				}
				results[c] *= rows[i][c]
			case "+":
				results[c] += rows[i][c]
			}
		}
	}

	// print column-wise results, one per line
	totalSum := 0
	for c, v := range results {
		fmt.Printf("Column %d: %d\n", c, v)
		totalSum += v
	}
	fmt.Printf("Answer %d\n", totalSum)
}
