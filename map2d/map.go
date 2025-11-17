package map2d

import "fmt"

type Map[E comparable] struct {
	// The Tiles type should be agit pull
	Width, Height int   // top left corner is (0,0)
	Tiles         [][]E // dimensions (width,height) (x,y)
}

type C struct { // Coordinates
	x, y int
}

var (
	up    = C{0, -1}
	down  = C{0, 1}
	left  = C{-1, 0}
	right = C{1, 0}
)

func (m *Map[E]) getTile(x, y int) (E, error) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		var zero E
		return zero, fmt.Errorf("coordinates %d, %d are out of range for map with Width %d and Height %d", x, y, m.Width, m.Height)
	}
	return m.Tiles[x][y], nil
}

func (m *Map[E]) isTile(x, y int, e E) (bool, error) {
	tile, err := m.getTile(x, y)
	if err != nil {
		return false, err
	}
	return tile == e, nil
}

func (m *Map[E]) setTile(x, y int, e E) (bool, error) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return false, fmt.Errorf("coordinates %d, %d are out of range for map with Width %d and Height %d", x, y, m.Width, m.Height)
	}
	m.Tiles[x][y] = e
	return true, nil
}

func (m *Map[E]) moveTileTo(x1, y1, x2, y2 int, empty E) (bool, error) {
	e, err := m.getTile(x1, y1)
	if err != nil {
		return false, err
	}
	m.Tiles[x2][y2] = e
	m.Tiles[x1][y1] = empty
	return true, nil
}

func (m *Map[E]) moveTileBy(x, y, dx, dy int, empty E) (bool, error) {
	success, err := m.moveTileTo(x, y, x+dx, y+dy, empty)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (m *Map[E]) countTile(e E) (count int) {
	for i := range m.Width {
		for j := range m.Height {
			if e == m.Tiles[i][j] {
				count++
			}
		}
	}
	return
}

func (m *Map[E]) findAll(e E) (locations []C, count int) {
	for i := range m.Width {
		for j := range m.Height {
			if e == m.Tiles[i][j] {
				locations = append(locations, C{i, j})
				count++
			}
		}
	}
	return locations, count
}

func ImportMap[E comparable](b []byte) (m Map[E]) {
	// Imports a map given as 2D array of characters such as:
	// #####\n
	// #..@#\n
	// #^..#\n
	// #####EOF
}

// print
