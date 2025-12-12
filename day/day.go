// I am day X of AOC2025
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	filename := "example.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't read ", filename)
	}
	var input []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(string(input[0]))

	part1logic()
	part2logic()

	fmt.Println("Part 1: ", "")
	fmt.Println("Part 2: ", "")

}

func part1logic() {
	return
}

func part2logic() {
	return
}
