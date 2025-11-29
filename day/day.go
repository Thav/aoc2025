// I am day X of AOC2025
package main

import (
	"fmt"
	"log"

	"github.com/Thav/aoc2025/map2d"
)

var mapString = []byte("#####\n#..@#\n#^..#\n#####")
var directionsString = "<^<<v><v>^"
var directionsMap = map[rune]map2d.C{
	'<': map2d.Left,
	'^': map2d.Up,
	'>': map2d.Right,
	'v': map2d.Down,
}

func main() {
	levelMap := map2d.ImportMap(mapString)
	tile, err := levelMap.GetTile(3, 1)
	if err != nil {
		log.Fatalln("getTile failed", err)
	}
	if tile != "@" {
		log.Fatalf("Expected %v to be @\n", tile)
	}
	fmt.Println(levelMap)
}
