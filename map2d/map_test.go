package map2d

import (
	"fmt"
	"testing"
)

var mapString []byte = []byte("#####\n#..@#\n#^..#\n#####")
var directionsString string = "<^<<v><v>^"
var directionsMap map[rune]C = map[rune]C{
	'<': Left,
	'^': Up,
	'>': Right,
	'v': Down,
}

func TestImportMap(t *testing.T) {
	levelMap := ImportMap(mapString)
	tile, err := levelMap.getTile(3, 1)
	if err != nil {
		t.Error("getTile failed", err)
		return
	}
	if tile != "@" {
		t.Errorf("Expected %v to be @", tile)
		return
	}
	fmt.Println(levelMap)
}

func TestOthers(t *testing.T) {
	fmt.Println("Hello, World!")
	directionsImport, err := ImportDirections(directionsString, directionsMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(directionsImport)

}
