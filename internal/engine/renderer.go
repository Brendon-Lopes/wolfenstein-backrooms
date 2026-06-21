package engine

import (
	"image/color"
	"math"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const PlayerSize = 8

func DrawMiniPlayer(screen *ebiten.Image, p *entity.Player, m *world.Map) {

	centerX := p.X + (float64(PlayerSize) / 2)
	centerY := p.Y + (float64(PlayerSize) / 2)
	lineLength := 25.0

	// player square
	vector.FillRect(
		screen,
		float32(p.X), float32(p.Y),
		float32(PlayerSize), float32(PlayerSize),
		// TODO: not alocate color every render
		color.RGBA{255, 255, 0, 255}, false,
	)

	// player direction line
	vector.StrokeLine(
		screen,
		float32(centerX), float32(centerY),
		float32(centerX+math.Cos(p.Angle)*lineLength), float32(centerY+math.Sin(p.Angle)*lineLength),
		// TODO: not alocate color every render
		1, color.RGBA{255, 255, 0, 255}, false,
	)

}

func DrawMiniMap(screen *ebiten.Image, m *world.Map) {
	var c color.Color

	for y := range m.Height {
		for x := range m.Width {
			xo := x * m.TileSize
			yo := y * m.TileSize

			if m.Grid[y*m.Width+x] == 1 {
				c = color.White
			} else {
				c = color.Black
			}

			vector.FillRect(
				screen,
				float32(xo), float32(yo),
				float32(m.TileSize), float32(m.TileSize),
				c, false,
			)

			vector.StrokeRect(
				screen,
				float32(xo), float32(yo),
				float32(m.TileSize), float32(m.TileSize),
				2,
				color.RGBA{50, 50, 50, 255},
				false,
			)

		}
	}
}

func Draw3dWorld(screen *ebiten.Image, p *entity.Player, m *world.Map) {
	centerX := p.X + (float64(PlayerSize) / 2)
	centerY := p.Y + (float64(PlayerSize) / 2)

	screenHeight := screen.Bounds().Dy()
	screenWidth := screen.Bounds().Dx()
	// xOffset := width

	for x := range screenWidth {
		// finalX, finalY, wallHeight, side := CalculateRay(x, screenWidth, screenHeight, centerX, centerY, m, p)
		_, _, wallHeight, side := CalculateRay(x, screenWidth, screenHeight, centerX, centerY, m, p)

		drawStart := (screenHeight - int(wallHeight)) / 2

		var c color.Color
		if side == 0 {
			c = color.RGBA{180, 180, 180, 255}
		} else {
			c = color.RGBA{120, 120, 120, 255}
		}

		vector.FillRect(screen,
			// float32(x+xOffset), float32(drawStart),
			float32(x), float32(drawStart),
			1, float32(wallHeight),
			c,
			false,
		)

		// vector.StrokeLine(
		// 	screen,
		// 	float32(centerX), float32(centerY),
		// 	float32(finalX), float32(finalY),
		// 	2,
		// 	color.RGBA{255, 0, 0, 255},
		// 	false,
		// )
	}
}
