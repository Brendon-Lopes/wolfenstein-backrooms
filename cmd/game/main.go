package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	windowWidth  = 1280
	windowHeight = 720
	resDivision  = 2
)

type Game struct{}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth / resDivision, windowHeight / resDivision
}

func main() {
	fmt.Println("Hello")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Wolfenrooms")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
