package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

// maxSubsequence returns the lexicographically largest subsequence of s
// with exactly k characters, preserving original order.
func maxSubsequence(s string, k int) (string, error) {
	n := len(s)
	if k <= 0 {
		return "", fmt.Errorf("k must be positive")
	}
	if n < k {
		return "", fmt.Errorf("input too short: need %d, have %d", k, n)
	}

	result := make([]byte, 0, k)
	pos := 0
	for picks := range k {
		// end is the last index we can choose for this pick
		// if we need to pick (k-picks) digits including this one,
		// we must leave space for remaining-1 digits after the chosen index.
		remaining := k - picks
		end := n - remaining // inclusive

		// find max digit in s[pos : end+1]
		bestIdx := pos
		best := s[pos]
		for j := pos + 1; j <= end; j++ {
			if s[j] > best {
				best = s[j]
				bestIdx = j
				// early exit if we find '9'
				if best == '9' {
					break
				}
			}
		}

		result = append(result, best)
		pos = bestIdx + 1
	}

	return string(result), nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	k := 12
	total := big.NewInt(0)

	for scanner.Scan() {
		line := scanner.Text()
		// trim possible spaces, but assume file lines are pure digit strings
		if len(line) == 0 {
			continue
		}

		best, err := maxSubsequence(line, k)
		if err != nil {
			// if line too short just skip or treat as zero; here we skip
			log.Printf("skipping line: %v", err)
			continue
		}

		fmt.Println(best)

		n := new(big.Int)
		if _, ok := n.SetString(best, 10); !ok {
			log.Printf("failed to parse %q, skipping", best)
			continue
		}
		total.Add(total, n)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum of all battery values is: %s\n", total.String())
}
