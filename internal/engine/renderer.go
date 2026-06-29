package engine

import (
	"image"
	"image/color"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const PlayerSize = 2
const HeightMultiplier = 2.5

func DrawMiniMap(screen *ebiten.Image, p *entity.Player, m *world.Map, rays []RayResult, scale, minimapOffsetX, minimapOffsetY float32) {
	var c color.Color
	blockSize := float32(m.TileSize) * scale

	for y := range m.Height {
		for x := range m.Width {
			xOffset := float32(x)*blockSize + minimapOffsetX
			yOffset := float32(y)*blockSize + minimapOffsetY

			if m.Grid[y*m.Width+x] >= 1 {
				c = color.White
			} else {
				c = color.Black
			}

			vector.FillRect(
				screen,
				float32(xOffset), float32(yOffset),
				float32(blockSize), float32(blockSize),
				c,
				false,
			)

			vector.StrokeRect(
				screen,
				float32(xOffset), float32(yOffset),
				float32(blockSize), float32(blockSize),
				1,
				color.RGBA{50, 50, 50, 255},
				false,
			)
		}
	}

	centerX := (p.X*float64(scale) + (float64(PlayerSize) / 2)) + float64(minimapOffsetX)
	centerY := (p.Y*float64(scale) + (float64(PlayerSize) / 2)) + float64(minimapOffsetY)
	lineLength := 6.0

	for _, ray := range rays {
		finalX := ray.FinalX*float64(scale) + float64(minimapOffsetX)
		finalY := ray.FinalY*float64(scale) + float64(minimapOffsetY)

		vector.StrokeLine(
			screen,
			float32(centerX), float32(centerY),
			float32(finalX), float32(finalY),
			1,
			color.RGBA{255, 0, 0, 255},
			false,
		)
	}

	playerX := p.X*float64(scale) + float64(minimapOffsetX)
	playerY := p.Y*float64(scale) + float64(minimapOffsetY)

	vector.FillRect(
		screen,
		float32(playerX), float32(playerY),
		float32(PlayerSize), float32(PlayerSize),
		// TODO: not alocate color every render
		color.RGBA{255, 255, 0, 255}, false,
	)

	vector.StrokeLine(
		screen,
		float32(centerX), float32(centerY),
		float32(centerX+p.DirX*lineLength), float32(centerY+p.DirY*lineLength),
		// TODO: not alocate color every render
		1,
		color.RGBA{255, 255, 0, 255},
		false,
	)
}

func Draw3dWorld(pixelBuffer []byte, screenWidth, screenHeight int, rays []RayResult, textures map[byte]*image.RGBA) {
	for _, ray := range rays {
		drawStart := int((float64(screenHeight) - ray.WallHeight) / 2)
		drawEnd := int((float64(screenHeight) + ray.WallHeight) / 2)

		if drawStart < 0 {
			drawStart = 0
		}
		if drawEnd >= screenHeight {
			drawEnd = screenHeight - 1
		}

		tex := textures[ray.TileType]
		texHeight := tex.Rect.Dy()

		step := float64(texHeight) / ray.WallHeight
		texPos := (float64(drawStart) - float64(screenHeight)/2 + ray.WallHeight/2) * step

		var fullLightDist float64 = 2.0
		var maxDarkDist float64 = 12.0
		var dist float64 = ray.PerpWallDist

		mult := (maxDarkDist - dist) / (maxDarkDist - fullLightDist)
		if dist <= fullLightDist {
			mult = 1
		}
		if dist >= maxDarkDist {
			mult = 0
		}
		if ray.Side == 0 {
			mult *= 0.7
		}

		multFixed := uint32(mult * 256)
		tx := ray.TextureX

		for y := drawStart; y <= drawEnd; y++ {
			ty := int(texPos) & (texHeight - 1)
			texPos += step

			texIndex := ty*tex.Stride + tx*4
			r := tex.Pix[texIndex]
			g := tex.Pix[texIndex+1]
			b := tex.Pix[texIndex+2]

			idx := (y*screenWidth + ray.X) * 4
			pixelBuffer[idx] = uint8((uint32(r) * multFixed) >> 8)
			pixelBuffer[idx+1] = uint8((uint32(g) * multFixed) >> 8)
			pixelBuffer[idx+2] = uint8((uint32(b) * multFixed) >> 8)
			pixelBuffer[idx+3] = 255
		}
	}
}

func DrawFloor(pixelBuffer []byte, screenWidth, screenHeight int, p *entity.Player, textures map[byte]*image.RGBA) {
	for y := screenHeight/2 + 1; y < screenHeight; y++ {
		rayDirX0 := p.DirX - p.PlaneX
		rayDirY0 := p.DirY - p.PlaneY
		rayDirX1 := p.DirX + p.PlaneX
		rayDirY1 := p.DirY + p.PlaneY

		ps := y - screenHeight/2

		posZ := 0.5 * float64(screenHeight) * HeightMultiplier

		rowDistance := posZ / float64(ps)

		floorStepX := rowDistance * (rayDirX1 - rayDirX0) / float64(screenWidth)
		floorStepY := rowDistance * (rayDirY1 - rayDirY0) / float64(screenWidth)

		floorX := (p.X / float64(world.TileSize)) + rowDistance*rayDirX0
		floorY := (p.Y / float64(world.TileSize)) + rowDistance*rayDirY0

		var fullLightDist float64 = 2.0
		var maxDarkDist float64 = 12.0

		mult := (maxDarkDist - rowDistance) / (maxDarkDist - fullLightDist)
		if rowDistance <= fullLightDist {
			mult = 1
		}
		if rowDistance >= maxDarkDist {
			mult = 0
		}

		multFixed := uint32(mult * 256)

		floorTex := textures[2]
		ceilingTex := textures[3]
		floorStride := floorTex.Stride
		ceilingStride := ceilingTex.Stride

		tileSizeFloat := float64(world.TileSize)
		tileSizeMask := world.TileSize - 1

		for x := range screenWidth {
			tx := int(floorX*tileSizeFloat) & tileSizeMask
			ty := int(floorY*tileSizeFloat) & tileSizeMask

			floorX += floorStepX
			floorY += floorStepY

			texIndexF := ty*floorStride + tx*4
			rF := floorTex.Pix[texIndexF]
			gF := floorTex.Pix[texIndexF+1]
			bF := floorTex.Pix[texIndexF+2]

			texIndexC := ty*ceilingStride + tx*4
			rC := ceilingTex.Pix[texIndexC]
			gC := ceilingTex.Pix[texIndexC+1]
			bC := ceilingTex.Pix[texIndexC+2]

			floorIdx := (y*screenWidth + x) * 4
			pixelBuffer[floorIdx] = uint8((uint32(rF) * multFixed) >> 8)
			pixelBuffer[floorIdx+1] = uint8((uint32(gF) * multFixed) >> 8)
			pixelBuffer[floorIdx+2] = uint8((uint32(bF) * multFixed) >> 8)
			pixelBuffer[floorIdx+3] = 255

			ceilY := screenHeight - y - 1
			ceilIdx := (ceilY*screenWidth + x) * 4
			pixelBuffer[ceilIdx] = uint8((uint32(rC) * multFixed) >> 8)
			pixelBuffer[ceilIdx+1] = uint8((uint32(gC) * multFixed) >> 8)
			pixelBuffer[ceilIdx+2] = uint8((uint32(bC) * multFixed) >> 8)
			pixelBuffer[ceilIdx+3] = 255
		}
	}
}
