// I am day 9 of AOC2025
package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/Thav/aoc2025/convert"
)

type point struct {
	x, y int
}

type rectangle struct {
	corners []point
	area    int
}

type edge struct {
	points []point
}

func main() {
	filename := "puzzle.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't read ", filename)
	}
	var points []point
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")
		p := point{
			x: convert.ToInt(numbers[0]),
			y: convert.ToInt(numbers[1])}
		points = append(points, p)
	}

	fmt.Println(points[0])

	rectangles := part1logic(points)
	fmt.Println("Part 1: ", rectangles[0].area)
	area := part2logic(points, rectangles)

	fmt.Println("Part 2: ", area)

}

func findArea(p1, p2 point) int {
	dx := p1.x - p2.x
	if dx < 0 {
		dx = -dx
	}
	dy := p1.y - p2.y
	if dy < 0 {
		dy = -dy
	}
	a := (dx + 1) * (dy + 1)
	if a > 0 {
		return a
	}
	return -a
}

func (r *rectangle) findArea() {
	if len(r.corners) != 2 {
		log.Fatalln("expected two corners on a rectangle", r)
	}
	r.area = findArea(r.corners[0], r.corners[1])
}

func part1logic(points []point) (rectangles []rectangle) {
	for i, p1 := range points {
		for _, p2 := range points[i:] {
			rect := rectangle{corners: []point{p1, p2}}
			rect.findArea()
			rectangles = append(rectangles, rect)
		}
	}
	slices.SortFunc(rectangles, func(a, b rectangle) int {
		return -cmp.Compare(a.area, b.area)
	})
	return rectangles
}

func buildEdges(points []point) (vertical, horizontal []edge) {
	var i int
	for i = range len(points) - 1 {
		p1 := points[i]
		p2 := points[i+1]
		if p1.x == p2.x {
			vertical = append(vertical, edge{points: []point{p1, p2}})
		} else if p1.y == p2.y {
			horizontal = append(horizontal, edge{points: []point{p1, p2}})
		} else {
			log.Fatalln("non orthogonal point pair", p1, p2)
		}
	}
	p1 := points[i+1]
	p2 := points[0]
	if p1.x == p2.x {
		vertical = append(vertical, edge{points: []point{p1, p2}})
	} else if p1.y == p2.y {
		horizontal = append(horizontal, edge{points: []point{p1, p2}})
	} else {
		log.Fatalln("non orthogonal point pair", p1, p2)
	}
	slices.SortFunc(horizontal, func(a, b edge) int { return cmp.Compare(a.points[0].y, b.points[0].y) })
	slices.SortFunc(vertical, func(a, b edge) int { return cmp.Compare(a.points[0].x, b.points[0].x) })
	return
}

func findPointsXRange(points []point, p1, p2 point, inclusive bool) (found []point) {
	var check1, check2 point
	check1 = p1
	check2 = p2
	if p1.x > p2.x {
		check1 = p2
		check2 = p1
	}
	if inclusive {
		check2.x++
	} else {
		check1.x++
	}
	indexFrom, _ := slices.BinarySearchFunc(points, check1, func(a, b point) int { return cmp.Compare(a.x, b.x) })
	// What if the point is beyond the range of points?
	if indexFrom >= len(points) {
		return
	}
	indexTo, _ := slices.BinarySearchFunc(points[indexFrom:], check2, func(a, b point) int {
		return cmp.Compare(a.x, b.x)
	})
	indexTo += indexFrom
	// // What if it's at the very end, could be, why not?
	// indexTo := indexFrom + 1
	// if indexFrom == len(points)-1 {
	// 	indexTo = indexFrom
	// } else {
	// 	for i, next := range points[indexFrom+1:] {
	// 		if next.x >= check2.x {
	// 			indexTo += i
	// 			break
	// 		}
	// 	}

	// }
	found = slices.Clone(points[indexFrom:indexTo])
	found = slices.DeleteFunc(found, func(a point) bool { return a == p1 || a == p2 })
	return
}
func findPointsX(points []point, p point) (found []point) {
	indexFrom, itemFound := slices.BinarySearchFunc(points, p, func(a, b point) int { return cmp.Compare(a.x, b.x) })
	if !itemFound {
		return
	}
	var indexTo int
	for i, next := range points[indexFrom:] {
		if next.x != p.x {
			indexTo = i
			break
		}
	}
	found = slices.Clone(points[indexFrom:indexTo])
	found = slices.DeleteFunc(found, func(a point) bool { return a == p })
	return
}

func findPointsYRange(points []point, p1, p2 point, inclusive bool) (found []point) {
	var check1, check2 point
	check1 = p1
	check2 = p2
	if p1.y > p2.y {
		check1 = p2
		check2 = p1
	}
	if inclusive {
		check2.y++
	} else {
		check1.y++
	}
	indexFrom, _ := slices.BinarySearchFunc(points, check1, func(a, b point) int {
		return cmp.Compare(a.y, b.y)
	})
	// What if the point is beyond the range of points?
	if indexFrom >= len(points) {
		return
	}
	// What if it's at the very end, could be, why not?
	indexTo, _ := slices.BinarySearchFunc(points[indexFrom:], check2, func(a, b point) int {
		return cmp.Compare(a.y, b.y)
	})
	indexTo += indexFrom
	// if indexFrom == len(points)-1 {
	// 	indexTo = indexFrom
	// } else {
	// 	for i, next := range points[indexFrom+1:] {
	// 		if next.y >= check2.y {
	// 			indexTo += i
	// 			break
	// 		}
	// 	}
	// }
	found = slices.Clone(points[indexFrom:indexTo])
	found = slices.DeleteFunc(found, func(a point) bool { return a == p1 || a == p2 })
	return
}
func findPointsY(points []point, p point) (found []point) {
	indexFrom, itemFound := slices.BinarySearchFunc(points, p, func(a, b point) int { return cmp.Compare(a.y, b.y) })
	if !itemFound {
		return
	}
	var indexTo int
	for i, next := range points[indexFrom:] {
		if next.y != p.y {
			indexTo = i
			break
		}
	}
	found = slices.Clone(points[indexFrom:indexTo])
	found = slices.DeleteFunc(found, func(a point) bool { return a == p })
	return
}

func findPointsInRectangle(points []point, r rectangle, inclusive bool) (found []point) {
	p1 := r.corners[0]
	p2 := r.corners[1]
	pointsInXRange := findPointsXRange(points, p1, p2, inclusive)
	slices.SortFunc(pointsInXRange, func(a, b point) int { return cmp.Compare(a.y, b.y) })
	found = findPointsYRange(pointsInXRange, p1, p2, inclusive)
	// if len(found) > 0 {
	// 	s := "inclusive"
	// 	if !inclusive {
	// 		s = "exclusive"
	// 	}
	// 	fmt.Printf("Points in rectangle %v, %v:\n", r, s)
	// 	fmt.Println(found)
	// }
	return
}

func isRectangleFilled(vertical []edge, r rectangle) bool {
	// define center before edge true, odd true
	p1 := r.corners[0]
	p2 := r.corners[1]
	center := point{
		x: (p1.x + p2.x) / 2,
		y: (p1.y + p2.y) / 2}
	// find and count crossing edges
	edgesCrossed := 0
	remaining := slices.Clone(vertical)
	index := 0
	for {
		remaining = remaining[index:]
		index = slices.IndexFunc(remaining, func(e edge) bool {
			ys := []int{e.points[0].y, e.points[1].y}
			slices.Sort(ys)
			return center.y >= ys[0] && center.y < ys[1]
		})
		if index == -1 {
			break
		}
		edgesCrossed++
		e := remaining[index]
		// If the next edge found is at or past our rectangle, stop.
		// no need to check points[1] since it's vertical
		if e.points[0].x >= p1.x && e.points[0].x >= p2.x {
			break
		}
		index++
	}
	if edgesCrossed == 0 {
		log.Fatalln("No edges crossed", r)
	}
	if index == -1 {
		log.Fatalln("how could this beeeeee", r)
	}
	odd := true
	if edgesCrossed%2 == 0 {
		odd = false
	}
	before := true
	if center.x > remaining[index].points[0].x {
		before = false
	}
	return before != odd
}

func isRectangleCrossed(vertical, horizontal []edge, r rectangle) bool {
	p1 := r.corners[0]
	p2 := r.corners[1]
	xs := []int{p1.x, p2.x}
	slices.Sort(xs)
	ys := []int{p1.y, p2.y}
	slices.Sort(ys)
	center := point{
		x: (p1.x + p2.x) / 2,
		y: (p1.y + p2.y) / 2}
	// find vertical edges within rectangle's x boundaries
	indexFrom := slices.IndexFunc(vertical, func(e edge) bool {
		return e.points[0].x > xs[0]
	})
	indexTo := slices.IndexFunc(vertical[indexFrom:], func(e edge) bool {
		return e.points[0].x >= xs[1]
	})
	indexTo += indexFrom
	// does a horizontal ray through center cross any of these edges?
	for _, e := range vertical[indexFrom:indexTo] {
		eys := []int{e.points[0].y, e.points[1].y}
		slices.Sort(eys)
		if eys[0] < center.y && eys[1] > center.y {
			return true
		}
	}

	// find horizontal edges within rectangle's y boundaries
	indexFrom = slices.IndexFunc(horizontal, func(e edge) bool {
		return e.points[0].y > ys[0]
	})
	indexTo = slices.IndexFunc(horizontal[indexFrom:], func(e edge) bool {
		return e.points[0].y >= ys[1]
	})
	indexTo += indexFrom
	// does a vertical ray through center cross any of these edges?
	for _, e := range horizontal[indexFrom:indexTo] {
		exs := []int{e.points[0].x, e.points[1].x}
		slices.Sort(exs)
		if exs[0] < center.x && exs[1] > center.x {
			return true
		}
	}

	return false
}

func part2logic(points []point, rectangles []rectangle) (area int) {
	pointsSortX := slices.Clone(points)
	slices.SortFunc(pointsSortX, func(a, b point) int { return cmp.Compare(a.x, b.x) })
	pointsSortY := slices.Clone(points)
	slices.SortFunc(pointsSortY, func(a, b point) int { return cmp.Compare(a.y, b.y) })
	vertical, horizontal := buildEdges(points)
	// Loop through rectangles, starting with greatest area
	for _, r := range rectangles {
		// Run through conditions
		// 1. If any other points are inside the rectangle, discard
		//    because some tiles must be non-red or non-green
		foundPoints := findPointsInRectangle(pointsSortX, r, false)
		if len(foundPoints) > 0 {
			continue
		}
		// 2. Unlikely, but if the largest area is the degenerate
		//    case of a width/height = 1 rectangle, with no points inside
		//    its rectangle,it is ALWAYS valid.
		//    Good to avoid the weird ways the checks below might fail.
		if r.corners[0].x == r.corners[1].x || r.corners[0].y == r.corners[1].y {
			fmt.Println("Step 2:", r)
			return r.area
		}
		// X. If any other points are not on the edges of the rectangle,
		//    all tiles in the rectangle beside the points of the rectangle
		//    are green, you have a winner! This step is important because
		//    these are harder to detect with #3
		// Forgot a case there there's a cove, so didn't work
		// foundPoints = findPointsInRectangle(pointsSortX, r, true)
		// if len(foundPoints) == 0 {
		// 	fmt.Println("Step 3:", r)
		// 	return r.area
		// }
		// 3. Draw a ray through both axes intersecting the center of our rectangle,
		//    and if an edge occurs in the range of the rectangle, discard
		if isRectangleCrossed(vertical, horizontal, r) {
			continue
		}
		// 4. Draw a ray from an axis intersecting the center of our rectangle,
		//    find and count the edges it crosses
		//    through until you reach one that belongs to our rectangle. Track
		//    if the center is before or after the found edge as you continue
		//    along the ray, and if the edge encountered was even or odd, XOR them
		//    e.g. center is after edge, edge is odd = center is in, winner!
		//         center is before edge, edge is odd = center is out, discard
		if !isRectangleFilled(vertical, r) {
			continue
		}
		fmt.Println("Passed:", r)
		return r.area
	}
	return
}
