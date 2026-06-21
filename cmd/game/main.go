package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/engine"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/input"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	windowWidth  = 1024
	windowHeight = 512
	resDivision  = 1

	minimapScale   = 0.25
	minimapOffsetX = 10
	minimapOffsetY = 10
)

// TODO: move palette to own package and initialize once
type Palette struct {
	Gray color.RGBA
}

// TODO: move palette to own package and initialize once
var palette = Palette{
	Gray: color.RGBA{50, 50, 50, 255},
}

type Game struct {
	player   *entity.Player
	worldMap *world.Map
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(palette.Gray)
	engine.Draw3dWorld(screen, g.player, g.worldMap)
	engine.DrawMiniMap(screen, g.player, g.worldMap, minimapScale, minimapOffsetX, minimapOffsetY)

	fps := ebiten.ActualFPS()
	ebitenutil.DebugPrintAt(screen, "FPS: "+strconv.FormatFloat(fps, 'f', 0, 64), windowWidth-60, 10)
}

func (g *Game) Update() error {
	cmd := input.ReadInput()
	entity.UpdatePlayer(g.player, cmd)
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth / resDivision, windowHeight / resDivision
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Wolfenrooms")

	p := entity.NewPlayer()
	m := world.NewMap()
	g := &Game{p, m}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
