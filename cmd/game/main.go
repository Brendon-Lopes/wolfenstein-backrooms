package main

import (
	"image/color"
	"log"
	"math"
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
	engine.DrawMiniMap(screen, g.worldMap)
	engine.DrawMiniPlayer(screen, g.player, g.worldMap)

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

	// TODO: initialize map somewhere else
	mapGrid := [...]byte{
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 1, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
	}

	// TODO: initialize player somewhere else (maybe object pool for inits)
	pdx := math.Cos(0) * entity.PlayerSpeed
	pdy := math.Sin(0) * entity.PlayerSpeed

	p := &entity.Player{X: 200, Y: 200, DeltaX: pdx, DeltaY: pdy, DirX: 1, DirY: 0, PlaneX: 0, PlaneY: 0.66}
	m := &world.Map{Width: 8, Height: 8, TileSize: engine.TileSize, Grid: mapGrid}
	g := &Game{p, m}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
