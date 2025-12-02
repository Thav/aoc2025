// I am day 1 of AOC2025
package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
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

	var position int = 50
	var count, count2 int
	pattern := regexp.MustCompile(`([LR])(\d{1,3})`)
	matches := pattern.FindAllSubmatch(input, -1)
	for i, line := range matches {
		var direction = 1
		if string(line[1]) == "L" {
			direction = -1
		}
		distance, err := strconv.Atoi(string(line[2]))
		if err != nil {
			log.Fatalln("got some nonsense when trying to get a distance on ", string(line[2]))
		}
		previousPosition := position
		previousCount2 := count2
		position += distance * direction
		if position < 0 {
			count2 += int(math.Floor(float64(-1*position-1) / 100))
			if previousPosition > 0 {
				count2++
			}
			position = (position % 100) + 100
		}
		if position > 99 {
			count2 += int(math.Floor(float64(position-1) / 100))
			position = position % 100
		}
		if position == 0 {
			count++
			count2++
		}
		if count2-previousCount2 > 1 {
			fmt.Printf("%d: turned %d by %d, went from position %d to %d, count increased from %d to %d\n", i, direction, distance, previousPosition, position, previousCount2, count2)
		}
	}
	fmt.Println("Got this many for part 1:", count)
	fmt.Println("Got this many for part 2:", count2)
}
