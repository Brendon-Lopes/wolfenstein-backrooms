package world

type Map struct {
	Width    int
	Height   int
	TileSize int
	Grid     [64]byte
}

const TileSize = 64

func NewMap() *Map {
	mapGrid := [...]byte{
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 1, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
	}

	return &Map{
		Width: 8, Height: 8,
		TileSize: TileSize,
		Grid:     mapGrid,
	}
}
