package entity

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y, DeltaX, DeltaY, Angle float64
}

// TODO: mover pra UI
func (p *Player) DrawPlayer(screen *ebiten.Image) {
	playerSize := 8

	vector.FillRect(
		screen,
		float32(p.X),
		float32(p.Y),
		float32(playerSize),
		float32(playerSize),
		color.RGBA{255, 255, 0, 255},
		false,
	)

	centerX := p.X + (float64(playerSize) / 2)
	centerY := p.Y + (float64(playerSize) / 2)

	vector.StrokeLine(
		screen,
		float32(centerX),
		float32(centerY),
		float32(centerX+p.DeltaX*5),
		float32(centerY+p.DeltaY*5),
		2,
		color.RGBA{255, 255, 0, 255},
		false,
	)
}

// TODO: mover pra input
func ReadInput(p *Player) {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Angle -= 0.1

		if p.Angle < 0 {
			p.Angle += 2 * math.Pi
		}

		p.DeltaX = math.Cos(p.Angle) * 5
		p.DeltaY = math.Sin(p.Angle) * 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Angle += 0.1

		if p.Angle > 2*math.Pi {
			p.Angle -= 2 * math.Pi
		}

		p.DeltaX = math.Cos(p.Angle) * 5
		p.DeltaY = math.Sin(p.Angle) * 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.X += p.DeltaX
		p.Y += p.DeltaY
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.X -= p.DeltaX
		p.Y -= p.DeltaY
	}
}
