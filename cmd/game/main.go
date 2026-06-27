package main

import (
	"image"
	"image/color"
	"log"
	"strconv"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/assets"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/engine"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/input"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	windowWidth  = 1920
	windowHeight = 1080
	resDivision  = 3

	minimapScale   = 0.125
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
	player      *entity.Player
	worldMap    *world.Map
	textures    map[byte]*image.RGBA
	rays        []engine.RayResult
	pixelBuffer []byte
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the buffer
	for i := 0; i < len(g.pixelBuffer); i += 4 {
		g.pixelBuffer[i] = palette.Gray.R
		g.pixelBuffer[i+1] = palette.Gray.G
		g.pixelBuffer[i+2] = palette.Gray.B
		g.pixelBuffer[i+3] = 255
	}

	screenWidth := windowWidth / resDivision
	screenHeight := windowHeight / resDivision

	g.rays = engine.CalculateAllRays(
		screenWidth, screenHeight,
		g.player.X+(float64(engine.PlayerSize)/2), g.player.Y+(float64(engine.PlayerSize)/2),
		g.rays, g.worldMap, g.player,
	)

	engine.DrawFloor(g.pixelBuffer, screenWidth, screenHeight, g.player, g.textures)
	engine.Draw3dWorld(g.pixelBuffer, screenWidth, screenHeight, g.rays, g.textures)

	screen.WritePixels(g.pixelBuffer)
	// engine.DrawMiniMap(screen, g.player, g.worldMap, g.rays, minimapScale, minimapOffsetX, minimapOffsetY)

	fps := ebiten.ActualFPS()
	ebitenutil.DebugPrintAt(screen, "FPS: "+strconv.FormatFloat(fps, 'f', 0, 64), screenWidth-60, 10)
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

	t, err := assets.LoadTextures()
	if err != nil {
		log.Fatal(err)
	}
	p := entity.NewPlayer()
	m := world.NewMap()
	g := &Game{
		player:      p,
		worldMap:    m,
		textures:    t,
		rays:        make([]engine.RayResult, 0, (windowWidth/resDivision)+1),
		pixelBuffer: make([]byte, (windowWidth/resDivision)*(windowHeight/resDivision)*4),
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
