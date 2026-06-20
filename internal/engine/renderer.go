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

func DrawRayY(screen *ebiten.Image, playerCenterX, playerCenterY float64, p *entity.Player) {
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
	}

	// 2. find first target Y
	mapFinalY = targetY * TileSize

	// 3. diff from player to Y
	diffPlayerToY := targetY - playerTilePositionY // inverted subtraction to get negative value, as sine will also be negative, thus result positive

	// 4. sideDistY (hypotenuse) -> SOH -> H = O/S
	sideDistY := diffPlayerToY / math.Sin(p.Angle)

	// 5. finalX = initialX + (direction * hypotenuse)
	playerTilePositionX := playerCenterX / float64(TileSize) // Ex.: 100 / 64 = 1.56
	finalX := playerTilePositionX + (math.Cos(p.Angle) * sideDistY)
	mapFinalX = finalX * TileSize

	for range 1 {
		vector.StrokeLine(
			screen,
			float32(playerCenterX),
			float32(playerCenterY),
			float32(mapFinalX),
			float32(mapFinalY),
			2,
			color.RGBA{255, 0, 0, 255},
			false,
		)
	}
}

func DrawRayX(screen *ebiten.Image, playerCenterX, playerCenterY float64, p *entity.Player) {
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
	}

	mapFinalX = targetX * TileSize

	diffPlayerToX := targetX - playerTilePositionX

	// CAH -> h = a / c
	sideDistX := diffPlayerToX / math.Cos(p.Angle)

	// finalY = initialY + (dir * hypotenuse)
	playerTilePositionY := playerCenterY / TileSize
	finalY := playerTilePositionY + (math.Sin(p.Angle) * sideDistX)
	mapFinalY = finalY * TileSize

	for range 1 {
		vector.StrokeLine(
			screen,
			float32(playerCenterX),
			float32(playerCenterY),
			float32(mapFinalX),
			float32(mapFinalY),
			4,
			color.RGBA{0, 255, 0, 255},
			false,
		)
	}
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

	DrawRayY(screen, centerX, centerY, p)
	DrawRayX(screen, centerX, centerY, p)
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
