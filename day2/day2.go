// I am day 2 of AOC2025
package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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
	fmt.Println(string(example[0:10]), "\n", string(puzzle[0:10]))
	input := puzzle
	rangesRaw := strings.Split(string(input), ",")
	rangesInt := make([][]int, 0)
	rangesString := make([][]string, 0)
	invalidSum1 := 0
	invalidSum2 := 0
	for _, r := range rangesRaw {
		splitRange := strings.Split(r, "-")
		rangesString = append(rangesString, splitRange)
		from := toInt(splitRange[0])
		to := toInt(splitRange[1])
		rangesInt = append(rangesInt, []int{from, to})
		invalidSum1 += part1logic(from, to)
		invalidSum2 += part2logic(splitRange[0], splitRange[1])

	}
	fmt.Println("Part 1: ", invalidSum1)
	fmt.Println("Part 2: ", invalidSum2)

}

func part1logic(from, to int) (sum int) {
	for i := from; i <= to; i++ {
		productID := strconv.Itoa(i)
		if len(productID)%2 != 0 {
			continue
		}
		if productID[0:len(productID)/2] == productID[len(productID)/2:len(productID)] {
			sum += i
		}
	}
	return
}

func part2logic(from, to string) (sum int) {
	intFrom := toInt(from)
	intTo := toInt(to)
	for numDigits := len(from); numDigits <= len(to); numDigits++ {
		seen := make(map[int]bool)
		for segmentLen := 1; segmentLen <= numDigits/2; segmentLen++ {
			if numDigits%segmentLen != 0 {
				continue
			}
			repeat := numDigits / segmentLen
			for segment := range int(math.Pow10(segmentLen)) {
				var numberBuilder strings.Builder
				for range repeat {
					numberBuilder.WriteString(strconv.Itoa(segment))
				}
				check := toInt(numberBuilder.String())
				if _, ok := seen[check]; !ok && inRange(check, intFrom, intTo) {
					seen[check] = true
					sum += check
				}
			}
		}
	}

	return
}

func toInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("couldn't convert to int: ", s)
	}
	return value
}
func inRange(value, from, to int) bool {
	return value >= from && value <= to
}
