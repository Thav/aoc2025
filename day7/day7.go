// I am day 7 of AOC2025
package main

import (
	"fmt"
	"log"
	"os"
	"slices"

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
	fmt.Println(string(example[0:10]), string(puzzle[0:10]))
	input := puzzle

	grid := grid.ImportGrid(input)

	// fmt.Println(grid)
	_, splits := part1logic(grid)
	// fmt.Println(grid)
	timeSplits := part2logic(grid)

	fmt.Println("Part 1: ", splits)
	fmt.Println("Part 2: ", timeSplits)

}

func part1logic(g grid.Grid) (beams, splits int) {
	// Find emitter "S"
	emitters, n := g.FindAll("S")
	if n != 1 {
		log.Fatal("Found more than one emitter")
	}
	// Find splitters
	_, n = g.FindAll("^")
	fmt.Println(n)
	for len(emitters) > 0 {
		emitter := emitters[0]
		// Abort early if beam already present, may not be necessary
		// given the checks in splitEmitters
		// below, err := g.GetTile(emitter.X, emitter.Y+1)
		// if err != nil {
		// 	log.Fatalln("Couldn't get tile to check for beam", err, g)
		// }
		// if below == "|" {
		// 	emitters = emitters[1:]
		// 	continue
		// }
		// Look for splitter or end in beam's path
		column, err := g.GetColumn(emitter.X)
		if err != nil {
			log.Fatal("Couldn't get the column", err, g)
		}
		beamIndex := slices.Index(column[emitter.Y:], "|")
		splitterIndex := slices.Index(column[emitter.Y:], "^")
		if beamIndex != -1 && beamIndex < splitterIndex {
			// Beam is hitting another beam before reaching
			// the next splitter. Draw to it but don't split
			drawBeam(g, emitter.X, emitter.Y+1, emitter.Y+1+beamIndex)
			emitters = emitters[1:]
			continue
		}
		if splitterIndex == -1 {
			// Reached the end!
			beams++
			drawBeam(g, emitter.X, emitter.Y+1, g.Height)
			emitters = emitters[1:]
			continue
		}
		// Encountered splitter
		// repeat the above with the coordinates up left and up right
		// of the splitter as new emitters
		splits++
		drawBeam(g, emitter.X, emitter.Y+1, emitter.Y+splitterIndex)
		emitters = emitters[1:]
		newEmitters := splitEmitters(g, grid.C{emitter.X, emitter.Y + splitterIndex}, true)
		emitters = append(emitters, newEmitters...)
	}

	return beams, splits
}

func splitEmitters(g grid.Grid, position grid.C, part1 bool) (newEmitters []grid.C) {
	if position.Y > 0 {
		if position.X > 0 {
			newEmitters = append(newEmitters, grid.C{position.X - 1, position.Y - 1})
		}
		if position.X < g.Width-1 {
			newEmitters = append(newEmitters, grid.C{position.X + 1, position.Y - 1})
		}
	}
	if part1 {
		return slices.DeleteFunc(newEmitters, func(p grid.C) bool {
			tile, err := g.GetTile(p.X, p.Y+1)
			if err != nil {
				log.Fatalln("Couldn't get tile", err)
			}
			return tile == "|"
		})
	}
	return newEmitters
}

func drawBeam(g grid.Grid, column, from, to int) {
	for i := range to - from {
		g.SetTile(column, from+i, "|")
	}
}

func part2logic(g grid.Grid) (timelines int) {
	// Find emitter "S"
	emitters, n := g.FindAll("S")
	if n != 1 {
		log.Fatal("Found more than one emitter")
	}
	// Find splitters
	// splitters, n = g.FindAll("^")
	// Want a
	splitterMap := make(map[grid.C]node)
	// Fill out the tree
	emitter := emitters[0]
	nextSplitter, end := beamNext(g, emitter)
	if end {
		return 0
	}
	// Locate first splitter
	// splitters := []grid.C{nextSplitter}
	// splitterMap[nextSplitter] = 0
	timelines = traverse(g, splitterMap, nextSplitter)
	return
}

func traverse(g grid.Grid, splitterMap map[grid.C]node, splitter grid.C) (timelines int) {
	// get left and right starting points, if they are on the grid
	newEmitters := splitEmitters(g, splitter, false)
	// for those beam starting points, see if they hit another splitter
	for _, emitter := range newEmitters {
		nextSplitter, end := beamNext(g, emitter)
		// If this child is a leaf, add a timeline
		if end {
			timelines++
			continue
		}
		// If we've seen this child before, add its timelines
		splitterNode, ok := splitterMap[nextSplitter]
		if ok {
			timelines += int(splitterNode)
			continue
		}
		// New child splitter, keep walking down
		timelines += traverse(g, splitterMap, nextSplitter)
	}
	splitterMap[splitter] = node(timelines)
	return
}

func beamNext(g grid.Grid, emitter grid.C) (nextEmitter grid.C, end bool) {
	column, err := g.GetColumn(emitter.X)
	if err != nil {
		log.Fatal("Couldn't get the column", err, g)
	}
	// not checking for beams in this version for part 2
	splitterIndex := slices.Index(column[emitter.Y:], "^")
	if splitterIndex == -1 {
		// Reached the end!
		return nextEmitter, true
	}
	// Encountered splitter
	return grid.C{emitter.X, emitter.Y + splitterIndex}, false
}

type node int

//struct {
// 	children []grid.C
// 	splits  int
// }
