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

func Move(p *Player, dir float64) {
	p.X += p.DeltaX * dir
	p.Y += p.DeltaY * dir
}

func Rotate(p *Player, dir float64) {
	p.Angle += AngleStep * dir

	if p.Angle < 0 {
		p.Angle += 2 * math.Pi
	} else if p.Angle > 2*math.Pi {
		p.Angle -= 2 * math.Pi
	}

	p.DeltaX = math.Cos(p.Angle) * PlayerSpeed
	p.DeltaY = math.Sin(p.Angle) * PlayerSpeed
}

func UpdatePlayer(p *Player, cmd input.Command) {
	const positive float64 = 1
	const negative float64 = -1

	if cmd.TurnLeft {
		Rotate(p, negative)
	}
	if cmd.TurnRight {
		Rotate(p, positive)
	}
	if cmd.MoveForward {
		Move(p, positive)
	}
	if cmd.MoveBackward {
		Move(p, negative)
	}
}
