package input

import "github.com/hajimehoshi/ebiten/v2"

type Command struct {
	MoveForward  bool
	MoveBackward bool
	TurnLeft     bool
	TurnRight    bool
}

func ReadInput() Command {
	return Command{
		MoveForward:  ebiten.IsKeyPressed(ebiten.KeyW),
		MoveBackward: ebiten.IsKeyPressed(ebiten.KeyS),
		TurnLeft:     ebiten.IsKeyPressed(ebiten.KeyA),
		TurnRight:    ebiten.IsKeyPressed(ebiten.KeyD),
	}
}
