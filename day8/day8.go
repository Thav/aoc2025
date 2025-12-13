// I am day 8 of AOC2025
package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/Thav/aoc2025/convert"
)

type point struct {
	x, y, z int
}

// I guess this is a disjoint tree set union data structure
type circuit struct {
	position point
	parent   *circuit
	size     int
}

type circuits []*circuit

type edge struct {
	distance int
	c1, c2   *circuit
}
type distances []edge

func findSet(c *circuit) (root *circuit) {
	root = c.parent
	for root != root.parent {
		root = root.parent
	}

	// flatten, as the wikipedia page for DSU describes
	// may not be necessary for this task, but let's go
	for c.parent != root {
		parent := c.parent
		c.parent = root
		c = parent
		// it feels weird to not update size here, but
		// it seems like this data structure is fine with
		// only the root nodes having a correct size
	}
	return
}

func connect(c1, c2 *circuit) {
	// get to the roots
	c1 = findSet(c1)
	c2 = findSet(c2)
	if c1 == c2 {
		return
	}
	if c1.size >= c2.size {
		c2.parent = c1
		c1.size += c2.size
		return
	}
	c1.parent = c2
	c2.size += c1.size
}

func main() {
	filename := "puzzle.txt"
	n := 10
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't read ", filename)
	}
	var circs circuits
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")
		circ := circuit{position: point{
			x: convert.ToInt(numbers[0]),
			y: convert.ToInt(numbers[1]),
			z: convert.ToInt(numbers[2])},
			size: 1}
		circ.parent = &circ
		circs = append(circs, &circ)
	}
	circuitSize := part1logic(circs, n)
	cordLength := part2logic(circs)

	fmt.Println("Part 1: ", circuitSize)
	// for _, c := range circs {
	// 	fmt.Println(c)
	// }
	fmt.Println("Part 2: ", cordLength)

}

func distance(p1, p2 point) float64 {
	return math.Sqrt(squareDiff(p1.x, p2.x) + squareDiff(p1.y, p2.y) + squareDiff(p1.z, p2.z))
}

func squareDiff(p1, p2 int) float64 {
	return math.Pow(float64(p1-p2), 2)
}

func part1logic(circs circuits, n int) (circuitSize int) {
	dist := getDistances(circs)
	// make n connections
	for i := range n {
		e := dist[i]
		connect(e.c1, e.c2)
	}

	// sort circuit list by size
	slices.SortFunc(circs, func(a, b *circuit) int { return cmp.Compare(b.size, a.size) })

	return int(circs[0].size * circs[1].size * circs[2].size)
}

func part2logic(circs circuits) (cordLength int) {
	dist := getDistances(circs)
	lastConnectionSize := 0
	var e edge
	// make connections until all boxes are in the same circuit
	for i := 0; lastConnectionSize < len(circs); i++ {
		e = dist[i]
		connect(e.c1, e.c2)
		lastConnectionSize = findSet(e.c1).size
		// fmt.Println(lastConnectionSize, e.c1, e.c2)
	}
	return int(e.c1.position.x * e.c2.position.x)
}

func getDistances(circs circuits) (dist distances) {
	for i, c1 := range circs {
		for _, c2 := range circs[i:] {
			if c1 == c2 {
				continue
			}
			d := distance(c1.position, c2.position)
			dInt := int(d * 100) // I don't trust using floats as keys
			dist = append(dist, edge{distance: dInt, c1: c1, c2: c2})
		}
	}
	slices.SortFunc(dist, func(a edge, b edge) int {
		return cmp.Compare(a.distance, b.distance)
	})
	return
}

func (circs circuits) String() string {
	var b strings.Builder
	var groups map[point][]point
	for _, c := range circs {
		root := findSet(c)
		group, ok := groups[root.position]
		if !ok {
			groups[root.position] = make([]point, 0)
		}
		groups[root.position] = append(group, c.position)
	}
	for _, v := range groups {
		b.WriteString(fmt.Sprintln(v))
	}
	return b.String()
}
