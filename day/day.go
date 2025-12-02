// I am day X of AOC2025
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	example, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatalln("Couldn't read example.txt")
	}
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Couldn't read input.txt")
	}
	fmt.Println(string(example[0:10]), string(input[0:10]))
}
