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
	windowWidth  = 1280
	windowHeight = 720
	resDivision  = 2

	minimapScale   = 0.125
	minimapOffsetX = 10
	minimapOffsetY = 10

	rayStep   = 4
	rectWidth = 4
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
	rays     []engine.RayResult
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(palette.Gray)

	centerX := g.player.X + (float64(engine.PlayerSize) / 2)
	centerY := g.player.Y + (float64(engine.PlayerSize) / 2)
	screenWidth := windowWidth / resDivision
	screenHeight := windowHeight / resDivision

	g.rays = engine.CalculateAllRays(
		rayStep,
		screenWidth, screenHeight,
		centerX, centerY,
		g.rays, g.worldMap, g.player,
	)

	engine.Draw3dWorld(screen, g.rays, rectWidth)
	engine.DrawMiniMap(screen, g.player, g.worldMap, g.rays, minimapScale, minimapOffsetX, minimapOffsetY)

	fps := ebiten.ActualFPS()
	ebitenutil.DebugPrintAt(screen, "FPS: "+strconv.FormatFloat(fps, 'f', 0, 64), windowWidth-60, 10)
}

func (g *Game) Update() error {
	cmd := input.ReadInput()
	entity.UpdatePlayer(g.player, cmd, g.worldMap)
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
	g := &Game{
		player:   p,
		worldMap: m,
		rays:     make([]engine.RayResult, 0, (windowWidth/resDivision)/rayStep+1),
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
