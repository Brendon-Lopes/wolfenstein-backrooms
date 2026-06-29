package assets

import (
	"embed"
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
)

//go:embed textures/*
var texturesFS embed.FS

func LoadTextures() (map[byte]*image.RGBA, error) {
	file, err := texturesFS.Open("textures/backrooms_map_textures.png")
	if err != nil {
		return nil, fmt.Errorf("opening spritesheet: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("decoding spritesheet: %w", err)
	}

	textures := make(map[byte]*image.RGBA)

	tileSize := 64
	tilesPerRow := 8

	tileAmount := 3

	for index := range tileAmount {
		col := index % tilesPerRow
		row := index / tilesPerRow

		startX := col * tileSize
		startY := row * tileSize

		rgba := image.NewRGBA(image.Rect(0, 0, tileSize, tileSize))

		draw.Draw(rgba, rgba.Bounds(), img, image.Point{X: startX, Y: startY}, draw.Src)

		textures[byte(index+1)] = rgba
	}

	return textures, nil
}
