package assets

import (
	"embed"
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed textures/*
var texturesFS embed.FS

func LoadTextures() (map[byte]*ebiten.Image, error) {
	files := map[byte]string{
		1: "textures/redbrick.png",
		2: "textures/eagle.png",
		3: "textures/wood.png",
		4: "textures/bluestone.png",
	}

	textures := make(map[byte]*ebiten.Image, len(files))

	for tileType, path := range files {
		img, _, err := ebitenutil.NewImageFromFileSystem(texturesFS, path)
		if err != nil {
			return nil, fmt.Errorf("decoding %s: %w", path, err)
		}
		textures[tileType] = img
	}

	return textures, nil
}
