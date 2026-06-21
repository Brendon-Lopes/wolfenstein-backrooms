package engine

import (
	"image/color"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const PlayerSize = 4

func DrawMiniMap(screen *ebiten.Image, p *entity.Player, m *world.Map, scale, minimapOffsetX, minimapOffsetY float32) {
	var c color.Color
	blockSize := float32(m.TileSize) * scale

	for y := range m.Height {
		for x := range m.Width {
			xOffset := float32(x)*blockSize + minimapOffsetX
			yOffset := float32(y)*blockSize + minimapOffsetY

			if m.Grid[y*m.Width+x] == 1 {
				c = color.White
			} else {
				c = color.Black
			}

			vector.FillRect(
				screen,
				float32(xOffset), float32(yOffset),
				float32(blockSize), float32(blockSize),
				c,
				false,
			)

			vector.StrokeRect(
				screen,
				float32(xOffset), float32(yOffset),
				float32(blockSize), float32(blockSize),
				1,
				color.RGBA{50, 50, 50, 255},
				false,
			)
		}
	}

	centerX := (p.X*float64(scale) + (float64(PlayerSize) / 2)) + float64(minimapOffsetX)
	centerY := (p.Y*float64(scale) + (float64(PlayerSize) / 2)) + float64(minimapOffsetY)
	lineLength := 10.0

	screenHeight := screen.Bounds().Dy()
	screenWidth := screen.Bounds().Dx()

	worldCenterX := p.X + (float64(PlayerSize) / 2)
	worldCenterY := p.Y + (float64(PlayerSize) / 2)

	batchSize := screenWidth / 20

	for x := range screenWidth {
		if x%batchSize != 0 {
			continue
		}

		rayX, rayY, _, _ := CalculateRay(x, screenWidth, screenHeight, worldCenterX, worldCenterY, m, p)

		finalX := rayX*float64(scale) + float64(minimapOffsetX)
		finalY := rayY*float64(scale) + float64(minimapOffsetY)

		vector.StrokeLine(
			screen,
			float32(centerX), float32(centerY),
			float32(finalX), float32(finalY),
			1,
			color.RGBA{255, 0, 0, 255},
			false,
		)
	}

	playerX := p.X*float64(scale) + float64(minimapOffsetX)
	playerY := p.Y*float64(scale) + float64(minimapOffsetY)

	// player square
	vector.FillRect(
		screen,
		float32(playerX), float32(playerY),
		float32(PlayerSize), float32(PlayerSize),
		// TODO: not alocate color every render
		color.RGBA{255, 255, 0, 255}, false,
	)

	// player direction line
	vector.StrokeLine(
		screen,
		float32(centerX), float32(centerY),
		float32(centerX+p.DirX*lineLength), float32(centerY+p.DirY*lineLength),
		// TODO: not alocate color every render
		1,
		color.RGBA{255, 255, 0, 255},
		false,
	)
}

func Draw3dWorld(screen *ebiten.Image, p *entity.Player, m *world.Map) {
	centerX := p.X + (float64(PlayerSize) / 2)
	centerY := p.Y + (float64(PlayerSize) / 2)

	screenHeight := screen.Bounds().Dy()
	screenWidth := screen.Bounds().Dx()
	rectWidth := 4

	for x := range screenWidth {
		if x%rectWidth != 0 {
			continue
		}
		_, _, wallHeight, side := CalculateRay(x, screenWidth, screenHeight, centerX, centerY, m, p)

		drawStart := (screenHeight - int(wallHeight)) / 2

		var c color.Color
		if side == 0 {
			c = color.RGBA{180, 180, 180, 255}
		} else {
			c = color.RGBA{120, 120, 120, 255}
		}

		vector.FillRect(screen,
			float32(x), float32(drawStart),
			float32(rectWidth), float32(wallHeight),
			c,
			false,
		)
	}
}
