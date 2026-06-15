package world

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Map struct {
	Width    int
	Height   int
	TileSize int
	Grid     []byte
}

// TODO: mover pra UI
func (m *Map) DrawMap(screen *ebiten.Image) {
	var x, y, xo, yo int
	var c color.Color

	for y = 0; y < m.Height; y++ {
		for x = 0; x < m.Height; x++ {
			xo = x * m.TileSize
			yo = y * m.TileSize

			if m.Grid[y*m.Width+x] == 1 {
				c = color.White
			} else {
				c = color.Black
			}

			vector.FillRect(
				screen,
				float32(xo),
				float32(yo),
				float32(m.TileSize),
				float32(m.TileSize),
				c,
				false,
			)

			vector.StrokeRect(
				screen,
				float32(xo),
				float32(yo),
				float32(m.TileSize),
				float32(m.TileSize),
				2,
				color.RGBA{50, 50, 50, 255},
				false,
			)

		}
	}
}
