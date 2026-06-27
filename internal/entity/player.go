package entity

import (
	"math"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/input"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"
)

type Player struct {
	X, Y, DirX, DirY, PlaneX, PlaneY, Angle float64
}

const PlayerSpeed float64 = 4
const AngleStep = 0.04

func NewPlayer() *Player {
	return &Player{
		X: 200, Y: 200,
		DirX: 1, DirY: 0,
		PlaneX: 0, PlaneY: 0.66,
	}
}

/*
[move] moves the player in a given direction.
*/
func move(p *Player, moveDirX, moveDirY float64, m *world.Map) {
	length := math.Hypot(moveDirX, moveDirY)

	moveX := 0.0
	moveY := 0.0

	if length != 0 {
		moveX = (moveDirX / length) * PlayerSpeed
		moveY = (moveDirY / length) * PlayerSpeed
	} else {
		return
	}

	safetyMargin := 20.0
	bufferX := 0.0
	bufferY := 0.0

	if moveX > 0 {
		bufferX = safetyMargin
	} else if moveX < 0 {
		bufferX = -safetyMargin
	}

	if moveY > 0 {
		bufferY = safetyMargin
	} else if moveY < 0 {
		bufferY = -safetyMargin
	}

	futureX := p.X + moveX + bufferX

	gridX := int(futureX / world.TileSize)
	currentGridY := int(p.Y / world.TileSize)

	if m.Grid[currentGridY*m.Width+gridX] == 0 {
		p.X += moveX
	}

	futureY := p.Y + moveY + bufferY

	currentGridX := int(p.X / world.TileSize)
	gridY := int(futureY / world.TileSize)

	if m.Grid[gridY*m.Width+currentGridX] == 0 {
		p.Y += moveY
	}
}

/*
[rotate] rotates the player by a given direction (dir).
The direction is 1 or -1.
The player's angle is updated based on the AngleStep constant,
and the DirX and DirY values are recalculated based on the new angle.
The angle is kept within the range of 0 to 2π radians.
*/
func rotate(p *Player, dir float64) {
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
}

func UpdatePlayer(p *Player, cmd input.Command, m *world.Map) {
	const positive float64 = 1
	const negative float64 = -1

	if cmd.TurnRight {
		rotate(p, positive)
	}
	if cmd.TurnLeft {
		rotate(p, negative)
	}

	intentX := 0.0
	intentY := 0.0

	if cmd.MoveForward {
		intentX += p.DirX
		intentY += p.DirY
	}
	if cmd.MoveBackward {
		intentX -= p.DirX
		intentY -= p.DirY
	}
	if cmd.StrafeRight {
		intentX -= p.DirY
		intentY += p.DirX
	}
	if cmd.StrafeLeft {
		intentX += p.DirY
		intentY -= p.DirX
	}

	if intentX != 0 || intentY != 0 {
		move(p, intentX, intentY, m)
	}
}
