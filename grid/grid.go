package grid

import (
	"fmt"
	"strings"
)

type Grid struct {
	Width, Height int        // top left corner is (0,0)
	Tiles         [][]string // dimensions (height, width) (y,x)
	// getters and setters use (x,y)
}

type C struct { // Coordinates
	X, Y int
}

var (
	Up    = C{0, -1}
	Down  = C{0, 1}
	Left  = C{-1, 0}
	Right = C{1, 0}
)

func (m *Grid) GetTile(x, y int) (string, error) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return "", fmt.Errorf("coordinates %d, %d are out of range for grid with Width %d and Height %d", x, y, m.Width, m.Height)
	}
	return m.Tiles[y][x], nil
}

func (m *Grid) GetRow(rowNumber int) (row []string, err error) {
	if rowNumber < 0 || rowNumber >= m.Height {
		return row, fmt.Errorf("row %d is out of range for grid with Height %d", rowNumber, m.Height)
	}
	return m.Tiles[rowNumber], nil
}

func (m *Grid) GetColumn(colNumber int) (col []string, err error) {
	if colNumber < 0 || colNumber >= m.Width {
		return col, fmt.Errorf("column %d is out of range for grid with Width %d", colNumber, m.Width)
	}
	for _, row := range m.Tiles {
		col = append(col, row[colNumber])
	}
	return col, nil
}

func (m *Grid) IsTile(x, y int, tileCompare string) (bool, error) {
	tile, err := m.GetTile(x, y)
	if err != nil {
		return false, err
	}
	return tile == tileCompare, nil
}

func (m *Grid) SetTile(x, y int, tile string) (bool, error) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return false, fmt.Errorf("coordinates %d, %d are out of range for grid with Width %d and Height %d", x, y, m.Width, m.Height)
	}
	m.Tiles[y][x] = tile
	return true, nil
}

func (m *Grid) MoveTileTo(x1, y1, x2, y2 int, empty string) (bool, error) {
	e, err := m.GetTile(x1, y1)
	if err != nil {
		return false, err
	}
	m.Tiles[y2][x2] = e
	m.Tiles[y1][x1] = empty
	return true, nil
}

func (m *Grid) MoveTileBy(x, y, dx, dy int, empty string) (bool, error) {
	success, err := m.MoveTileTo(x, y, x+dx, y+dy, empty)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (m *Grid) CountTile(tile string) (count int) {
	for j := range m.Height {
		for _, gridTile := range m.Tiles[j] {
			if tile == gridTile {
				count++
			}
		}
	}
	return
}

func (m *Grid) FindAll(tile string) (locations []C, count int) {
	for j := range m.Height {
		for i, gridTile := range m.Tiles[j] {
			if tile == gridTile {
				locations = append(locations, C{i, j})
				count++
			}
		}
	}
	return locations, count
}

func ImportGrid(b []byte) (m Grid) {
	// Imports a grid given as 2D array of characters such as:
	// #####\n
	// #..@#\n
	// #^..#\n
	// #####EOF
	var tilesRow []string
	for _, tileByte := range b {
		if tileByte == '\n' {
			m.Height++
			m.Tiles = append(m.Tiles, tilesRow)
			// make a new empty slice of strings with capacity for grid width
			tilesRow = make([]string, 0, m.Width)
		} else {
			tilesRow = append(tilesRow, string(tileByte))
		}
		if m.Height == 0 {
			m.Width++
		}
	}
	m.Tiles = append(m.Tiles, tilesRow)
	m.Height++
	return m
}

func (m Grid) Copy() (newGrid Grid) {
	newGrid.Height = m.Height
	newGrid.Width = m.Width
	for _, col := range m.Tiles {
		newCol := make([]string, len(col))
		copy(newCol, col)
		newGrid.Tiles = append(newGrid.Tiles, newCol)
	}
	return
}

func (m Grid) String() string {
	var b strings.Builder
	for j := range m.Height {
		for i := range m.Width {
			b.WriteString(m.Tiles[j][i])
		}
		b.WriteString("\n")
	}
	return b.String()
}

func ImportDirections(encodedDirections string, directionsMap map[rune]C) (directions []C, err error) {
	for _, encodedDirection := range encodedDirections {
		direction, ok := directionsMap[encodedDirection]
		if !ok {
			var empty []C
			return empty, fmt.Errorf("Found direction not in grid, %#v", encodedDirection)
		}
		directions = append(directions, direction)
	}
	return directions, nil
}
