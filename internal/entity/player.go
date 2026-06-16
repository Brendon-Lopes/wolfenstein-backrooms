package entity

import (
	"math"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/input"
)

type Player struct {
	X, Y, DeltaX, DeltaY, Angle float64
}

const PlayerSpeed float64 = 2
const AngleStep = 0.1

func Move(p *Player, forward bool) {
	if forward {
		p.X += p.DeltaX
		p.Y += p.DeltaY
		return
	}

	p.X -= p.DeltaX
	p.Y -= p.DeltaY
}

func Rotate(p *Player, dir float64) {
	p.Angle += dir * AngleStep

	if p.Angle < 0 {
		p.Angle += 2 * math.Pi
	} else if p.Angle > 2*math.Pi {
		p.Angle -= 2 * math.Pi
	}

	p.DeltaX = math.Cos(p.Angle) * PlayerSpeed
	p.DeltaY = math.Sin(p.Angle) * PlayerSpeed
}

func UpdatePlayer(p *Player, cmd input.Command) {
	if cmd.TurnLeft {
		Rotate(p, -1)
	}
	if cmd.TurnRight {
		Rotate(p, 1)
	}
	if cmd.MoveForward {
		Move(p, true)
	}
	if cmd.MoveBackward {
		Move(p, false)
	}
}
