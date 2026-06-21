package engine

import (
	"image/color"
	"math"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const TileSize = 64

func GetDeltaDistY(centerX, centerY float64, p *entity.Player) (float64, float64, float64) {
	var stepY float64

	var mapFinalY float64
	var mapFinalX float64

	if p.DeltaY < 0 {
		stepY = -1
	}
	if p.DeltaY > 0 {
		stepY = 1
	}
	if p.DeltaY == 0 {
		return math.Inf(1), 10000, stepY
	}

	playerTilePositionY := centerY / TileSize // Ex.: 100 / 64 = 1.56
	mapFinalY = (playerTilePositionY + stepY) * TileSize

	deltaDistY := stepY / math.Sin(p.Angle)

	playerTilePositionX := centerX / TileSize
	finalX := playerTilePositionX + (math.Cos(p.Angle) * deltaDistY)
	mapFinalX = finalX * TileSize

	return deltaDistY, mapFinalX, mapFinalY
}

func GetDeltaDistX(centerX, centerY float64, p *entity.Player) (float64, float64, float64) {
	var stepX float64

	var mapFinalY float64
	var mapFinalX float64

	if p.DeltaX < 0 {
		stepX = -1
	}
	if p.DeltaX > 0 {
		stepX = 1
	}
	if p.DeltaX == 0 {
		return math.Inf(1), stepX, 10000
	}

	playerTilePositionX := centerX / TileSize // Ex.: 100 / 64 = 1.56
	mapFinalX = (playerTilePositionX + stepX) * TileSize

	deltaDistX := stepX / math.Cos(p.Angle)

	playerTilePositionY := centerY / TileSize
	finalY := playerTilePositionY + (math.Sin(p.Angle) * deltaDistX)
	mapFinalY = finalY * TileSize

	return deltaDistX, mapFinalX, mapFinalY
}

func GetSideDistX(centerX, centerY, deltaDistX float64, p *entity.Player) (float64, float64, float64) {
	var targetX float64

	playerTilePositionX := centerX / TileSize

	if p.DeltaX < 0 {
		targetX = math.Floor(playerTilePositionX)
	}
	if p.DeltaX > 0 {
		targetX = math.Ceil(playerTilePositionX)
	}

	diffToSideX := math.Abs(targetX - playerTilePositionX)
	sideDistX := diffToSideX * deltaDistX
	mapFinalX := targetX * TileSize

	playerTilePositionY := centerY / TileSize
	sideDistY := playerTilePositionY + (math.Sin(p.Angle) * sideDistX)
	mapFinalY := sideDistY * TileSize

	return sideDistX, mapFinalX, mapFinalY
}

func GetSideDistY(centerX, centerY, deltaDistY float64, p *entity.Player) (float64, float64, float64) {
	var targetY float64

	playerTilePositionY := centerY / TileSize

	if p.DeltaY < 0 {
		targetY = math.Floor(playerTilePositionY)
	}
	if p.DeltaY > 0 {
		targetY = math.Ceil(playerTilePositionY)
	}

	diffToSideY := math.Abs(targetY - playerTilePositionY)
	sideDistY := diffToSideY * deltaDistY
	mapFinalY := targetY * TileSize

	playerTilePositionX := centerX / TileSize
	sideDistX := playerTilePositionX + (math.Cos(p.Angle) * sideDistY)
	mapFinalX := sideDistX * TileSize

	return sideDistY, mapFinalX, mapFinalY
}

func DrawMiniPlayer(screen *ebiten.Image, p *entity.Player) {
	playerSize := 8

	centerX := p.X + (float64(playerSize) / 2)
	centerY := p.Y + (float64(playerSize) / 2)
	lineLength := 25.0

	// player square
	vector.FillRect(
		screen,
		float32(p.X),
		float32(p.Y),
		float32(playerSize),
		float32(playerSize),
		// TODO: not alocate color every render
		color.RGBA{255, 255, 0, 255},
		false,
	)

	// player direction line
	vector.StrokeLine(
		screen,
		float32(centerX),
		float32(centerY),
		float32(centerX+math.Cos(p.Angle)*lineLength),
		float32(centerY+math.Sin(p.Angle)*lineLength),
		1,
		// TODO: not alocate color every render
		color.RGBA{255, 255, 0, 255},
		false,
	)

	var finalX, finalY float64

	deltaDistX, _, _ := GetDeltaDistX(centerX, centerY, p)
	deltaDistY, _, _ := GetDeltaDistY(centerX, centerY, p)

	sideDistX, xFinalX, xFinalY := GetSideDistX(centerX, centerY, deltaDistX, p)
	sideDistY, yFinalX, yFinalY := GetSideDistY(centerX, centerY, deltaDistY, p)

	if sideDistX > sideDistY {
		finalX = yFinalX
		finalY = yFinalY
	}
	if sideDistX < sideDistY {
		finalX = xFinalX
		finalY = xFinalY
	}

	vector.StrokeLine(
		screen,
		float32(centerX),
		float32(centerY),
		float32(finalX),
		float32(finalY),
		2,
		color.RGBA{255, 0, 0, 255},
		false,
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
