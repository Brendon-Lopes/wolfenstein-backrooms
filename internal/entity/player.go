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
	vector.FillRect(
		screen,
		float32(p.X),
		float32(p.Y),
		8,
		8,
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
