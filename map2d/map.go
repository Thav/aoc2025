package map2d

import (
	"fmt"
	"strings"
)

type Map struct {
	Width, Height int        // top left corner is (0,0)
	Tiles         [][]string // dimensions (height, width) (y,x)
	// getters and setters use (x,y)
}

type C struct { // Coordinates
	x, y int
}

var (
	Up    = C{0, -1}
	Down  = C{0, 1}
	Left  = C{-1, 0}
	Right = C{1, 0}
)

func (m *Map) getTile(x, y int) (string, error) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return "", fmt.Errorf("coordinates %d, %d are out of range for map with Width %d and Height %d", x, y, m.Width, m.Height)
	}
	return m.Tiles[y][x], nil
}

func (m *Map) isTile(x, y int, tileCompare string) (bool, error) {
	tile, err := m.getTile(x, y)
	if err != nil {
		return false, err
	}
	return tile == tileCompare, nil
}

func (m *Map) setTile(x, y int, tile string) (bool, error) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return false, fmt.Errorf("coordinates %d, %d are out of range for map with Width %d and Height %d", x, y, m.Width, m.Height)
	}
	m.Tiles[x][y] = tile
	return true, nil
}

func (m *Map) moveTileTo(x1, y1, x2, y2 int, empty string) (bool, error) {
	e, err := m.getTile(x1, y1)
	if err != nil {
		return false, err
	}
	m.Tiles[y2][x2] = e
	m.Tiles[y1][x1] = empty
	return true, nil
}

func (m *Map) moveTileBy(x, y, dx, dy int, empty string) (bool, error) {
	success, err := m.moveTileTo(x, y, x+dx, y+dy, empty)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (m *Map) countTile(tile string) (count int) {
	for j := range m.Height {
		for _, mapTile := range m.Tiles[j] {
			if tile == mapTile {
				count++
			}
		}
	}
	return
}

func (m *Map) findAll(tile string) (locations []C, count int) {
	for j := range m.Height {
		for i, mapTile := range m.Tiles[j] {
			if tile == mapTile {
				locations = append(locations, C{i, j})
				count++
			}
		}
	}
	return locations, count
}

func ImportMap(b []byte) (m Map) {
	// Imports a map given as 2D array of characters such as:
	// #####\n
	// #..@#\n
	// #^..#\n
	// #####EOF
	var tilesRow []string
	for _, tileByte := range b {
		if tileByte == '\n' {
			m.Height++
			m.Tiles = append(m.Tiles, tilesRow)
			// make a new empty slice of strings with capacity for map width
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

func (m Map) String() string {
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
			return empty, fmt.Errorf("Found direction not in map, %#v", encodedDirection)
		}
		directions = append(directions, direction)
	}
	return directions, nil
}
