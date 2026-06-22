package engine

import (
	"image/color"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const PlayerSize = 2

func DrawMiniMap(screen *ebiten.Image, p *entity.Player, m *world.Map, rays []RayResult, scale, minimapOffsetX, minimapOffsetY float32) {
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
	lineLength := 6.0

	screenWidth := screen.Bounds().Dx()
	batchSize := screenWidth / 20

	for _, ray := range rays {
		if ray.X%batchSize != 0 {
			continue
		}

		finalX := ray.FinalX*float64(scale) + float64(minimapOffsetX)
		finalY := ray.FinalY*float64(scale) + float64(minimapOffsetY)

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

	vector.FillRect(
		screen,
		float32(playerX), float32(playerY),
		float32(PlayerSize), float32(PlayerSize),
		// TODO: not alocate color every render
		color.RGBA{255, 255, 0, 255}, false,
	)

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

func Draw3dWorld(screen *ebiten.Image, rays []RayResult, rectWidth int) {
	screenHeight := screen.Bounds().Dy()

	for _, ray := range rays {
		drawStart := (screenHeight - int(ray.WallHeight)) / 2

		var c color.Color
		if ray.Side == 0 {
			c = color.RGBA{180, 180, 180, 255}
		} else {
			c = color.RGBA{120, 120, 120, 255}
		}

		vector.FillRect(screen,
			float32(ray.X), float32(drawStart),
			float32(rectWidth), float32(ray.WallHeight),
			c,
			false,
		)
	}
}
