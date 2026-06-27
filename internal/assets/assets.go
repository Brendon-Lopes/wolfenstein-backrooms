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
	files := map[byte]string{
		1: "textures/redbrick.png",
		2: "textures/eagle.png",
		3: "textures/wood.png",
		4: "textures/bluestone.png",
	}

	textures := make(map[byte]*image.RGBA, len(files))

	for tileType, path := range files {
		file, err := texturesFS.Open(path)
		if err != nil {
			return nil, fmt.Errorf("opening %s: %w", path, err)
		}
		img, _, err := image.Decode(file)
		file.Close()
		if err != nil {
			return nil, fmt.Errorf("decoding %s: %w", path, err)
		}

		rgba := image.NewRGBA(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)

		textures[tileType] = rgba
	}

	return textures, nil
}
