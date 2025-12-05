// I am day 3 of AOC2025
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Thav/aoc2025/lists"
)

func main() {
	example, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatalln("Couldn't read example.txt")
	}
	puzzle, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatalln("Couldn't read puzzle.txt")
	}
	fmt.Println(string(example[0:10]))
	fmt.Println(string(puzzle[0:10]))
	input := puzzle
	banks := lists.ImportRowListsInt(input, "")
	joltage1 := 0
	joltage2 := 0
	// var b strings.Builder
	for _, bank := range banks {
		indices := make(map[int][]int, 0)
		for i, jattery := range bank {
			if jattIndices, ok := indices[jattery]; ok {
				indices[jattery] = append(jattIndices, i)
			} else {
				indices[jattery] = []int{i}
			}
		}
		j1 := part1logic(bank, indices)
		j2 := part2logic(bank, indices)
		// b.WriteString(fmt.Sprintf("%d: %v \n %d\n", i, bank, j2))
		joltage1 += j1
		joltage2 += j2
	}
	fmt.Println("Part 1: ", joltage1)
	fmt.Println("Part 2: ", joltage2)
	// os.WriteFile("output.txt", []byte(b.String()), 0666)

}

func part1logic(bank []int, indices map[int][]int) (joltage int) {
	for j := 9; j > 0; j-- {
		jattIndices, ok := indices[j]
		if !ok || jattIndices[0] == (len(bank)-1) {
			continue
		}
		for k := 9; k > 0; k-- {
			kjattIndices, ok := indices[k]
			if !ok {
				continue
			}
			if kjattIndices[len(kjattIndices)-1] > jattIndices[0] {
				return j*10 + k
			}
		}

	}
	return 0
}

func part2logic(bank []int, indices map[int][]int) (sum int) {
	i := 0
	lastIndex := 0
	// Find the highest first digit with room for remaining digits behind it
	// Remove that index from list
	// Add that digit to the sum
	// Increment digit counter
	// track position of last digit
	// Find next best digit following the first, remove it
	// Find next best digit following the second ...
DigitLoop:
	for i < 12 {
		for j := 9; j > 0; j-- {
			jattIndices, ok := indices[j]
			if !ok || len(jattIndices) == 0 {
				continue
				// Don't have any 9, 8, 7...
			}
			for len(jattIndices) > 0 && jattIndices[0] < lastIndex {
				jattIndices = jattIndices[1:]
				// Remove leading index when past the last digit added
			}
			atMost := len(bank) - 11 + i
			if len(jattIndices) > 0 && jattIndices[0] >= atMost {
				continue
				// Remaining locations of this digit can't be next, but might work later
			}
			// now we have a good digit
			if len(jattIndices) > 0 {
				sum = sum*10 + j
				i++
				lastIndex = jattIndices[0]
				indices[j] = jattIndices[1:]
				jattIndices, ok = indices[j]
				continue DigitLoop
			}
		}
	}
	return
}
