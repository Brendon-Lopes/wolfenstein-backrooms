package engine

import (
	"image/color"
	"math"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawPlayer(screen *ebiten.Image, p *entity.Player) {
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
		color.RGBA{255, 255, 0, 255},
		false,
	)
}
