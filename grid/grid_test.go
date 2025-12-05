package grid

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// #####
// #..@#
// #^..#
// #####
var gridString = []byte("#####\n#..@#\n#^..#\n#####")
var directionsString = "<^<<v><v>^"
var directionsMap = map[rune]C{
	'<': Left,
	'^': Up,
	'>': Right,
	'v': Down,
}
var levelGrid Grid

func TestMain(m *testing.M) {
	levelGrid = ImportGrid(gridString)
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestImportGrid(t *testing.T) {
	levelGrid := ImportGrid(gridString)
	tile, err := levelGrid.GetTile(3, 1)
	if err != nil {
		t.Error("getTile failed", err)
		return
	}
	if tile != "@" {
		t.Errorf("Expected %v to be @", tile)
		return
	}
	fmt.Println(levelGrid)
}

func TestGetColumn(t *testing.T) {
	col, err := levelGrid.GetColumn(3)
	if err != nil {
		t.Error("GetColumn failed", err)
		return
	}
	var colExpected = []string{"#", "@", ".", "#"}
	if len(col) != len(colExpected) {
		t.Errorf("Returned column length wrong size. Got %d, expected %d", len(col), len(colExpected))
		return
	}
	for i := range len(col) {
		if col[i] != colExpected[i] {
			t.Errorf("Returned column has a mismatched character on row %d. Got %d, expected %d", i, len(col), len(colExpected))
		}
	}
	fmt.Println(col)
	col, err = levelGrid.GetColumn(-1)
	if err == nil {
		t.Error("Expected error when indexing column -1")
	}
	col, err = levelGrid.GetColumn(5)
	if err == nil {
		t.Error("Expected error when indexing column 5")
	}
	// if col != "@" {
	// 	t.Errorf("Expected %v to be @", tile)
	// 	return
	// }
}

func TestOthers(t *testing.T) {
	fmt.Println("Hello, World!")
	directionsImport, err := ImportDirections(directionsString, directionsMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(directionsImport)

}
