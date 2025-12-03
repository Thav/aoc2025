// I am day 3 of AOC2025
package main

import (
	"fmt"
	"log"
	"os"
	"slices"

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
	for _, bank := range banks {
		bank = slices.Compact(bank)
		indices := make(map[int][]int, 0)
		for i, jattery := range bank {
			if jattIndices, ok := indices[jattery]; ok {
				indices[jattery] = append(jattIndices, i)
			} else {
				indices[jattery] = []int{i}
			}
		}

		j1 := part1logic(bank, indices)
		joltage1 += j1
	}
	fmt.Println("Part 1: ", joltage1)
	fmt.Println("Part 2: ", "")

}

func part1logic(bank []int, indices map[int][]int) (joltage int) {
	for j := 9; j > 0; j-- {
		jattIndices, ok := indices[j]
		if !ok || jattIndices[0] == (len(bank)-1) {
			continue
		}
		if len(jattIndices) > 1 {
			return j*10 + j
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

func part2logic(from, to string) (sum int) {

	return
}
