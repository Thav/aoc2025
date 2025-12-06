// I am day 6 of AOC2025
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Thav/aoc2025/convert"
)

type problem struct {
	numbers  []int
	operator string
}

func main() {
	filename := "puzzle.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't read ", filename)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	problems1 := part1input(scanner)
	fmt.Println(problems1[0])
	p1 := part1logic(problems1)

	f, err = os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't read ", filename)
	}
	defer f.Close()
	scanner = bufio.NewScanner(f)

	problems2 := part2input(scanner)
	p2 := part1logic(problems2)
	fmt.Println(problems2)

	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)

}

func part1input(scanner *bufio.Scanner) (problems []problem) {
	for scanner.Scan() {
		input := scanner.Text()
		found := true
		for i := 0; found; i++ {
			input = strings.TrimLeft(input, " ")
			var chunk string
			chunk, input, found = strings.Cut(input, " ")
			if !found && chunk == "" {
				continue
			}
			if len(problems) == i {
				problems = append(problems, problem{})
			}
			if chunk == "*" || chunk == "+" {
				problems[i].operator = chunk
			} else {
				problems[i].numbers = append(problems[i].numbers, convert.ToInt(chunk))
			}
		}
	}
	return
}

func part2input(scanner *bufio.Scanner) (problems []problem) {
	var builders []strings.Builder
	first := true
	// Rotate the input -90*
	for scanner.Scan() {
		t := scanner.Text()
		if first {
			builders = make([]strings.Builder, len(t))
			first = false
		}
		for i := len(t) - 1; i > -1; i-- {
			builders[len(t)-i-1].WriteByte(t[i])
		}
	}
	var lines []string
	for _, b := range builders {
		lines = append(lines, b.String())
	}
	p := problem{}
	for _, line := range lines {
		if !strings.ContainsAny(line, "123456789") {
			problems = append(problems, p)
			p = problem{}
			continue
		}
		if strings.ContainsAny(line, "*+") {
			p.operator = string(line[len(line)-1])
			line = line[:len(line)-1]
		}
		number := strings.TrimSpace(line)
		p.numbers = append(p.numbers, convert.ToInt(number))
	}
	problems = append(problems, p)
	return
}

func part1logic(problems []problem) (sum int) {
	for _, p := range problems {
		partial := p.numbers[0]
		for i := range len(p.numbers) - 1 {
			n := p.numbers[i+1]
			if p.operator == "*" {
				partial *= n
			} else {
				partial += n
			}
		}
		sum += partial
	}
	return
}
