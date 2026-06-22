package input

import "github.com/hajimehoshi/ebiten/v2"

type Command struct {
	MoveForward  bool
	MoveBackward bool
	TurnRight    bool
	TurnLeft     bool
	StrafeRight  bool
	StrafeLeft   bool
}

func ReadInput() Command {
	return Command{
		MoveForward:  ebiten.IsKeyPressed(ebiten.KeyW),
		MoveBackward: ebiten.IsKeyPressed(ebiten.KeyS),
		TurnRight:    ebiten.IsKeyPressed(ebiten.KeyD),
		TurnLeft:     ebiten.IsKeyPressed(ebiten.KeyA),
		StrafeRight:  ebiten.IsKeyPressed(ebiten.KeyE),
		StrafeLeft:   ebiten.IsKeyPressed(ebiten.KeyQ),
	}
}
