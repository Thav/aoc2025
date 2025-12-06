// I am day 5 of AOC2025
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/Thav/aoc2025/convert"
)

func main() {
	filename := "puzzle.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't read ", filename)
	}
	var ranges [][]int
	var ids []int
	scanner := bufio.NewScanner(f)
	rangesRe := regexp.MustCompile(`(\d+)-(\d+)`)
	for scanner.Scan() {
		t := scanner.Text()
		if rangesRe.MatchString(t) {
			split := strings.Split(t, "-")
			ranges = append(ranges, []int{convert.ToInt(split[0]), convert.ToInt(split[1])})
		} else if t != "" {
			ids = append(ids, convert.ToInt(t))
		}
	}

	fmt.Println(ranges[0])
	fmt.Println(ids[0])

	// combinedRanges := combineRanges(ranges)
	p1 := part1logic(ranges, ids)
	p2 := part2logic(ranges)

	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)

}

// func combineRanges(ranges [][]int) (combined [][]int) {
// 	for r := range ranges {

// 	}
// }

func part1logic(ranges [][]int, ids []int) (foundFresh int) {
	for _, id := range ids {
		fresh := false
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				fresh = true
				break
			}
		}
		if fresh {
			// fmt.Printf("Product %d is fresh!\n", id)
			foundFresh++
		} else {
			// fmt.Printf("Product %d is rotten! :(\n", id)
		}
	}
	return
}

func part2logic(ranges [][]int) (validCount int) {
	// Sort ranges
	slices.SortFunc(ranges, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] > b[0] {
			return 1
		}
		return 0
	})
	// fmt.Println(ranges)
	// Combine ranges
	var combinedRanges [][]int
	combinedRanges = append(combinedRanges, ranges[0])
	ranges = ranges[1:]
	for len(ranges) > 0 {
		lastIndex := len(combinedRanges) - 1
		if combinedRanges[lastIndex][1] >= ranges[0][0] {
			combinedRanges[lastIndex][1] = max(combinedRanges[lastIndex][1], ranges[0][1])
		} else {
			combinedRanges = append(combinedRanges, ranges[0])
		}
		ranges = ranges[1:]
	}
	// fmt.Println(combinedRanges)
	for _, r := range combinedRanges {
		diff := r[1] - r[0] + 1
		// fmt.Println(r, diff)
		validCount += diff
	}
	// Count up
	return
}
