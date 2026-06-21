package entity

import (
	"math"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/input"
)

type Player struct {
	X, Y, DeltaX, DeltaY, DirX, DirY, PlaneX, PlaneY, Angle float64
}

const PlayerSpeed float64 = 2
const AngleStep = 0.1

/*
[Move] moves the player in a given direction (dir).
Dir is 1 for forward and -1 for backward.
*/
func Move(p *Player, dir float64) {
	p.X += p.DeltaX * dir
	p.Y += p.DeltaY * dir
}

/*
[Rotate] rotates the player by a given direction (dir).
The direction is 1 or -1.
The player's angle is updated based on the AngleStep constant,
and the DeltaX and DeltaY values are recalculated based on the new angle.
The angle is kept within the range of 0 to 2π radians.
*/
func Rotate(p *Player, dir float64) {
	p.Angle += AngleStep * dir

	if p.Angle < 0 {
		p.Angle += 2 * math.Pi
	} else if p.Angle > 2*math.Pi {
		p.Angle -= 2 * math.Pi
	}

	p.DirX = math.Cos(p.Angle)
	p.DirY = math.Sin(p.Angle)

	// TODO: centralize FOV somewhere else
	p.PlaneX = -p.DirY * 0.66
	p.PlaneY = p.DirX * 0.66

	p.DeltaX = p.DirX * PlayerSpeed
	p.DeltaY = p.DirY * PlayerSpeed
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
