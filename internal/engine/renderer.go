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

func GetDeltaDistY(rayAngle float64) (float64, int) {
	var stepY int
	rayDir := math.Sin(rayAngle)

	if rayDir < 0 {
		stepY = -1
	} else {
		stepY = 1
	}

	deltaDistY := float64(stepY) / rayDir

	return deltaDistY, stepY
}

func GetDeltaDistX(rayAngle float64) (float64, int) {
	var stepX int
	rayDir := math.Cos(rayAngle)

	if rayDir < 0 {
		stepX = -1
	} else {
		stepX = 1
	}

	deltaDistX := float64(stepX) / rayDir

	return deltaDistX, stepX
}

func GetSideDistX(playerTileX float64, stepX int, deltaDistX float64) (float64, int) {
	mapX := int(playerTileX)
	var sideDistX float64

	if stepX < 0 {
		sideDistX = (playerTileX - float64(mapX)) * deltaDistX
	} else {
		sideDistX = (float64(mapX) + 1.0 - playerTileX) * deltaDistX
	}

	return sideDistX, mapX
}

func GetSideDistY(playerTileY float64, stepY int, deltaDistY float64) (float64, int) {
	mapY := int(playerTileY)
	var sideDistY float64

	if stepY < 0 {
		sideDistY = (playerTileY - float64(mapY)) * deltaDistY
	} else {
		sideDistY = (float64(mapY) + 1.0 - playerTileY) * deltaDistY
	}

	return sideDistY, mapY
}

func DrawMiniPlayer(screen *ebiten.Image, p *entity.Player, m *world.Map) {
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

	playerTileX := centerX / float64(TileSize)
	playerTileY := centerY / float64(TileSize)

	deltaDistX, stepX := GetDeltaDistX(p.Angle)
	deltaDistY, stepY := GetDeltaDistY(p.Angle)

	sideDistX, mapX := GetSideDistX(playerTileX, stepX, deltaDistX)
	sideDistY, mapY := GetSideDistY(playerTileY, stepY, deltaDistY)

	hit := false
	side := 0 // 0 = hit X, 1 = hit Y

	for !hit {
		if sideDistX < sideDistY {
			sideDistX += deltaDistX
			mapX += stepX
			side = 0
		} else {
			sideDistY += deltaDistY
			mapY += stepY
			side = 1
		}

		if mapX >= 0 && mapX < m.Width && mapY >= 0 && mapY < m.Height {
			mapIndex := (mapY * m.Width) + mapX
			if m.Grid[mapIndex] == 1 {
				hit = true
			}
		} else {
			// out of bounds
			hit = true
		}
	}

	var totalDistance float64
	if side == 0 {
		totalDistance = sideDistX - deltaDistX
	} else {
		totalDistance = sideDistY - deltaDistY
	}

	rayDirX := math.Cos(p.Angle)
	rayDirY := math.Sin(p.Angle)

	finalX := centerX + (rayDirX * totalDistance * float64(TileSize))
	finalY := centerY + (rayDirY * totalDistance * float64(TileSize))

	vector.StrokeLine(
		screen,
		float32(centerX), float32(centerY),
		float32(finalX), float32(finalY),
		2, color.RGBA{255, 0, 0, 255}, false,
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
