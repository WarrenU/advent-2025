package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Range struct {
	Start *big.Int
	End   *big.Int
}

var numRe = regexp.MustCompile(`\d+`)

func cloneBig(x *big.Int) *big.Int { return new(big.Int).Set(x) }

func sanitizeRange(line string) (Range, error) {
	parts := numRe.FindAllString(line, -1)
	if len(parts) < 2 {
		return Range{}, errors.New("could not find two numbers")
	}
	start := new(big.Int)
	end := new(big.Int)
	if _, ok := start.SetString(parts[0], 10); !ok {
		return Range{}, fmt.Errorf("invalid number: %q", parts[0])
	}
	if _, ok := end.SetString(parts[1], 10); !ok {
		return Range{}, fmt.Errorf("invalid number: %q", parts[1])
	}
	if start.Cmp(end) > 0 {
		start, end = end, start
	}
	return Range{Start: start, End: end}, nil
}

// readFileSections: ranges until blank line, then ids as big.Int
func readFileSections(path string) ([]Range, []*big.Int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var ranges []Range
	var ids []*big.Int
	readingRanges := true

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			readingRanges = false
			continue
		}
		if readingRanges {
			r, err := sanitizeRange(line)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse range line %q: %w", line, err)
			}
			ranges = append(ranges, r)
		} else {
			n := new(big.Int)
			if _, ok := n.SetString(line, 10); !ok {
				// skip invalid id lines rather than failing
				continue
			}
			ids = append(ids, n)
		}
	}

	if err := sc.Err(); err != nil {
		return nil, nil, err
	}
	return ranges, ids, nil
}

// consolidateRanges merges touching or overlapping ranges globally
func consolidateRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start.Cmp(ranges[j].Start) < 0
	})

	out := make([]Range, 0, len(ranges))
	curr := Range{Start: cloneBig(ranges[0].Start), End: cloneBig(ranges[0].End)}

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]
		oneAfterCurrEnd := new(big.Int).Add(curr.End, big.NewInt(1))
		// if next.Start <= curr.End+1 => merge
		if next.Start.Cmp(oneAfterCurrEnd) <= 0 {
			if next.End.Cmp(curr.End) > 0 {
				curr.End = cloneBig(next.End)
			}
		} else {
			out = append(out, curr)
			curr = Range{Start: cloneBig(next.Start), End: cloneBig(next.End)}
		}
	}
	out = append(out, curr)
	return out
}

// findRangeIndex finds index i such that consolidated[i].Start > id,
// so candidate is i-1 (if >=0). This uses sort.Search and big.Int cmp.
func findPotentialIndex(consolidated []Range, id *big.Int) int {
	return sort.Search(len(consolidated), func(i int) bool {
		return consolidated[i].Start.Cmp(id) > 0
	})
}

func main() {
	const inputPath = "input.txt"

	ranges, ids, err := readFileSections(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	consolidated := consolidateRanges(ranges)

	fmt.Println("Consolidated ranges:")
	for _, r := range consolidated {
		fmt.Printf("  %s-%s\n", r.Start.String(), r.End.String())
	}
	fmt.Println()

	// Lookup each id via binary search
	matched := make([]*big.Int, 0, len(ids))
	for _, id := range ids {
		idx := findPotentialIndex(consolidated, id)
		if idx == 0 {
			// no range start is greater than id, candidate would be -1
			continue
		}
		candidate := consolidated[idx-1]
		if id.Cmp(candidate.Start) >= 0 && id.Cmp(candidate.End) <= 0 {
			matched = append(matched, id)
		}
	}

	fmt.Printf("Matched ids (%d/%d):\n", len(matched), len(ids))
	// for _, m := range matched {
	// 	fmt.Println(m.String())
	// }
}
