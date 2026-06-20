package engine

import (
	"fmt"
	"image/color"
	"math"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const TileSize = 64

func CalculateYRay(screen *ebiten.Image, playerCenterX, playerCenterY float64, p *entity.Player) (float64, float64, float64) {
	var mapFinalX float64
	var mapFinalY float64
	var playerTilePositionY float64
	var targetY float64

	// 1. facing up
	if p.DeltaY < 0 {
		playerTilePositionY = playerCenterY / float64(TileSize) // Ex.: 100 / 64 = 1.56
		targetY = math.Floor(playerTilePositionY)               // Ex.: 1.56 -> 1
	}

	// facing down
	if p.DeltaY > 0 {
		playerTilePositionY = playerCenterY / float64(TileSize)
		targetY = math.Ceil(playerTilePositionY)
	}

	if p.DeltaY == 0 {
		mapFinalX = playerCenterX + 10000
		mapFinalY = playerCenterY
		return mapFinalX, mapFinalY, math.Inf(1)
	}

	mapFinalY = targetY * TileSize

	diffPlayerToY := targetY - playerTilePositionY // inverted subtraction to get negative value, as sine will also be negative, thus result positive

	// sideDistY (hypotenuse) -> SOH -> H = O/S
	sideDistY := diffPlayerToY / math.Sin(p.Angle)

	// 5. finalX = initialX + (direction * hypotenuse)
	playerTilePositionX := playerCenterX / float64(TileSize) // Ex.: 100 / 64 = 1.56
	finalX := playerTilePositionX + (math.Cos(p.Angle) * sideDistY)
	mapFinalX = finalX * TileSize

	return mapFinalX, mapFinalY, sideDistY
}

func CalculateXRay(screen *ebiten.Image, playerCenterX, playerCenterY float64, p *entity.Player) (float64, float64, float64) {
	var mapFinalX float64
	var mapFinalY float64
	var playerTilePositionX float64
	var targetX float64

	// looking right
	if p.DeltaX > 0 {
		playerTilePositionX = playerCenterX / TileSize
		targetX = math.Ceil(playerTilePositionX)
	}

	// looking left
	if p.DeltaX < 0 {
		playerTilePositionX = playerCenterX / TileSize
		targetX = math.Floor(playerTilePositionX)
	}

	if p.DeltaX == 0 {
		mapFinalX = playerCenterX
		mapFinalY = playerCenterY + 10000
		return mapFinalX, mapFinalY, math.Inf(1)
	}

	mapFinalX = targetX * TileSize

	diffPlayerToX := targetX - playerTilePositionX

	// CAH -> h = a / c
	sideDistX := diffPlayerToX / math.Cos(p.Angle)

	// finalY = initialY + (dir * hypotenuse)
	playerTilePositionY := playerCenterY / TileSize
	finalY := playerTilePositionY + (math.Sin(p.Angle) * sideDistX)
	mapFinalY = finalY * TileSize

	return mapFinalX, mapFinalY, sideDistX
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

	verticalRayX, verticalRayY, sideDistY := CalculateYRay(screen, centerX, centerY, p)
	horizontalRayX, horizontalRayY, sideDistX := CalculateXRay(screen, centerX, centerY, p)

	var finalRayX, finalRayY float64

	if sideDistY > sideDistX {
		finalRayX, finalRayY = horizontalRayX, horizontalRayY
	} else {
		finalRayX, finalRayY = verticalRayX, verticalRayY
	}

	fmt.Printf("X: %f, Y: %f", finalRayX, finalRayY)

	vector.StrokeLine(
		screen,
		float32(centerX),
		float32(centerY),
		float32(finalRayX),
		float32(finalRayY),
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
