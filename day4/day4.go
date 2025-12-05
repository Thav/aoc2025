// I am day 4 of AOC2025
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Thav/aoc2025/grid"
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
	grid := grid.ImportGrid(input)
	accessible := 0
	removed := 0
	// var b strings.Builder
	for y := range grid.Height {
		for x := range grid.Width {
			a1 := part1logic(grid, x, y)
			// b.WriteString(fmt.Sprintf("%d: %v \n %d\n", i, bank, j2))
			accessible += a1
		}
	}
	r, newGrid := part2logic(grid)
	removed += r
	for r > 0 {
		r, newGrid = part2logic(newGrid)
		removed += r
	}
	fmt.Println("Part 1: ", accessible)
	fmt.Println("Part 2: ", removed)
	// os.WriteFile("output.txt", []byte(b.String()), 0666)

}

func part1logic(grid grid.Grid, x, y int) (accessible int) {
	// If a roll (@), check for less than four other rolls around
	tile, err := grid.GetTile(x, y)
	if err != nil || tile != "@" {
		return 0
	}
	surrounding := 0
	for x2 := x - 1; x2 <= x+1; x2++ {
		for y2 := y - 1; y2 <= y+1; y2++ {
			tile, err = grid.GetTile(x2, y2)
			if err != nil ||
				tile != "@" ||
				(x == x2 && y == y2) {
				continue
			}
			surrounding++
		}
	}
	if surrounding >= 4 {
		return 0
	}
	return 1
}

func part2logic(grid grid.Grid) (removed int, newGrid grid.Grid) {
	// Remove accessible rolls
	newGrid = grid.Copy()
	for y := range grid.Height {
		for x := range grid.Width {
			r := part1logic(grid, x, y)
			// b.WriteString(fmt.Sprintf("%d: %v \n %d\n", i, bank, j2))
			if r > 0 {
				removed += r
				success, _ := newGrid.SetTile(x, y, ".")
				if !success {
					log.Fatalln("something broke", x, y, grid)
				}
			}
		}
	}
	return
}
