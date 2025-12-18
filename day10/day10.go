// I am day X of AOC2025
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Thav/aoc2025/convert"
	"gonum.org/v1/gonum/stat/combin"
)

type factoryMachine struct {
	lights   []int
	buttons  [][]int
	joltages []int
}

func pressesToState(presses []int) []int {
	state := make([]int, len(presses))
	for i, p := range presses {
		state[i] = p % 2
	}
	return state
}

func main() {
	filename := "example.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't read ", filename)
	}
	var machines []factoryMachine
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		fmt.Println(split)
		lightsString := split[0]
		buttonsStrings := split[1 : len(split)-2]
		joltageString := split[len(split)-1]
		if lightsString[0] != '[' || joltageString[0] != '{' {
			log.Fatalln("Malformed input,", lightsString, joltageString)
		}
		lights := make([]int, len(lightsString)-2)
		for i, l := range lightsString[1 : len(lightsString)-1] {
			if l == '#' {
				lights[i] = 1
			}
		}
		joltages := convert.SliceToInt(strings.Split(joltageString[1:len(joltageString)-1], ","))
		buttons := make([][]int, len(buttonsStrings))
		for i, buttonsString := range buttonsStrings {
			if buttonsString[0] != '(' {
				log.Fatalln("Malformed input,", buttonsString)
			}
			buttons[i] = make([]int, len(lights))
			buttonPositions := convert.SliceToInt(strings.Split(buttonsString[1:len(buttonsString)-1], ","))
			for _, b := range buttonPositions {
				buttons[i][b] = 1
			}
		}
		machines = append(machines, factoryMachine{lights, buttons, joltages})
	}

	fmt.Println(machines[0])

	pressesOn := part1logic(machines)
	part2logic()

	fmt.Println("Part 1: ", pressesOn)
	fmt.Println("Part 2: ", "")

}

func part1logic(machines []factoryMachine) (pressesOn int) {
	for _, machine := range machines {
		n := len(machine.buttons)
		gen := combin.NewCombinationGenerator(n, n)
		fewest := n
		for gen.Next() {
			comb := gen.Combination()
		}
	}
	return
}

func part2logic() {
	return
}
