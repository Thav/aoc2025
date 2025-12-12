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
func TestCopy(t *testing.T) {
	newGrid := levelGrid.Copy()
	if newGrid.Height != levelGrid.Height ||
		newGrid.Width != levelGrid.Width ||
		len(newGrid.Tiles) != len(levelGrid.Tiles) ||
		len(newGrid.Tiles[0]) != len(levelGrid.Tiles[0]) {
		t.Error("Copy failed to give matching height and width", newGrid, levelGrid)
		return
	}
	newCol, err := newGrid.GetColumn(3)
	if err != nil {
		t.Error("Couldn't get expected column from newGrid", newGrid)
		return
	}
	levelCol, err := levelGrid.GetColumn(3)
	if err != nil {
		t.Error("Couldn't get expected column from levelGrid", levelGrid)
		return
	}
	if newCol[1] != levelCol[1] {
		t.Error("Grids ended up with different values", newGrid, levelGrid)
	}
	if &newGrid.Tiles == &levelGrid.Tiles ||
		&newGrid.Tiles[0] == &levelGrid.Tiles[0] {
		t.Error("Somehow copied slice directly", &newGrid.Tiles, &levelGrid.Tiles, &newGrid.Tiles[0], &levelGrid.Tiles[0])
		return
	}
	fmt.Println(newGrid)
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

func TestSetTile(t *testing.T) {
	grid := levelGrid.Copy()
	a, err := grid.GetTile(3, 1)
	if err != nil || a != "@" {
		t.Error("failed on the GetTile", grid.Tiles)
	}
	success, err := grid.SetTile(3, 1, "U")
	if !success {
		t.Error("failed on SetTile", grid)
	} else {
		b, err := grid.GetTile(3, 1)
		if err != nil || b != "U" {
			t.Error("SetTile is said to have succeeded but didn't", grid)
		}
	}
	fmt.Println(grid)
}
func TestFindAll(t *testing.T) {
	coords, n := levelGrid.FindAll("@")
	if n != 1 {
		t.Fatal("Should have found 1 @, got", n)
	}
	x := coords[0].x
	y := coords[0].y
	if x != 3 || y != 1 {
		t.Fatalf("Wrong coordinates returns, expected (3,1), got (%d,%d)", x, y)
	}
	coords, n = levelGrid.FindAll("#")
	if n != 14 {
		t.Fatal("Should have found 14 #, got", n)
	}
}
