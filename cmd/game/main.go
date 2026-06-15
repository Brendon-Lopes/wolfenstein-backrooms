package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 1280
	windowHeight = 720
	resDivision  = 1
)

type Game struct {
	player   *entity.Player
	worldMap *world.Map
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 50, 50, 255})
	g.worldMap.DrawMap(screen)
	entity.DrawPlayer(screen, g.player)
}

func (g *Game) Update() error {
	entity.ReadInput(g.player)
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth / resDivision, windowHeight / resDivision
}

func main() {
	fmt.Println("Hello")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Wolfenrooms")

	mapGrid := []byte{
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 1, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
	}

	p := &entity.Player{X: 200, Y: 200}
	m := &world.Map{Width: 8, Height: 8, TileSize: 64, Grid: mapGrid}
	g := &Game{p, m}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
