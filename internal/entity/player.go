package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y float64
}

// TODO: mover pra UI
func DrawPlayer(screen *ebiten.Image, p *Player) {
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
		p.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += 5
	}
}
