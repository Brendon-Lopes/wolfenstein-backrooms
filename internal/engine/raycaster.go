package engine

import (
	"math"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/entity"
	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"
)

type RayResult struct {
	X            int
	FinalX       float64
	FinalY       float64
	WallHeight   float64
	Side         int
	PerpWallDist float64
	TextureX     int
	TileType     byte
}

/*
[getDeltaDistX] calculates the length of the ray's hypotenuse when it travels exactly one map grid unit along the Y-axis.

Returns:
  - deltaDistX: The absolute length of the hypotenuse for a full grid step.
  - stepX: The direction to step on the map grid (-1 for left or 1 right).
*/
func getDeltaDistX(rayDir float64) (float64, int) {
	var stepX int

	if rayDir < 0 {
		stepX = -1
	}
	if rayDir > 0 {
		stepX = 1
	}
	if rayDir == 0 {
		return math.Inf(1), stepX
	}

	// CAH -> h=a/c
	deltaDistX := math.Abs(float64(stepX) / rayDir)

	return deltaDistX, stepX
}

/*
[getDeltaDistY] calculates the length of the ray's hypotenuse when it travels exactly one map grid unit along the Y-axis.

Returns:
  - deltaDistY: The absolute length of the hypotenuse for a full grid step.
  - stepY: The direction to step on the map grid (-1 for up or 1 for down).
*/
func getDeltaDistY(rayDir float64) (float64, int) {
	var stepY int

	if rayDir < 0 {
		stepY = -1
	}
	if rayDir > 0 {
		stepY = 1
	}
	if rayDir == 0 {
		return math.Inf(1), stepY
	}

	// SOH -> h=o/s
	deltaDistY := math.Abs(float64(stepY) / rayDir)

	return deltaDistY, stepY
}

/*
[getSideDistX] calculates the initial ray distance from the player's exact position to the first vertical grid line intersection.

Parameters:
  - playerTileX: The player's exact X coordinate in logical map units (e.g., 1.2).
  - stepX: The ray's X-axis direction (-1 or 1).
  - deltaDistX: The hypotenuse length of a full block step.

Returns:
  - sideDistX: The total ray length to reach the first vertical grid boundary.
  - mapX: The integer map coordinate where the player is currently standing.
*/
func getSideDistX(playerTileX float64, stepX int, deltaDistX float64) (float64, int) {
	var sideDistX float64
	mapX := int(playerTileX)

	if stepX < 0 {
		sideDistX = (playerTileX - float64(mapX)) * deltaDistX
	} else {
		sideDistX = (float64(mapX) + 1.0 - playerTileX) * deltaDistX
	}

	return sideDistX, mapX
}

/*
[getSideDistY] calculates the initial ray distance from the player's exact position to the first horizontal grid line intersection.

Parameters:
  - playerTileY: The player's exact Y coordinate in logical map units (e.g., 1.2).
  - stepY: The ray's Y-axis direction (-1 or 1).
  - deltaDistY: The hypotenuse length of a full block step.

Returns:
  - sideDistY: The total ray length to reach the first horizontal grid boundary.
  - mapY: The integer map coordinate where the player is currently standing.
*/
func getSideDistY(playerTileY float64, stepY int, deltaDistY float64) (float64, int) {
	var sideDistY float64
	mapY := int(playerTileY)

	if stepY < 0 {
		sideDistY = (playerTileY - float64(mapY)) * deltaDistY
	} else {
		sideDistY = (float64(mapY) + 1.0 - playerTileY) * deltaDistY
	}

	return sideDistY, mapY
}

/*
[CalculateRay] performs raycasting from the player's position to find the first wall intersection point.

Parameters:
  - centerX: The player's X coordinate in pixel space.
  - centerY: The player's Y coordinate in pixel space.
  - angle: The ray direction angle in radians.
  - m: Pointer to the world Map containing the grid layout.

Returns:
  - finalX: The X coordinate where the ray hits a wall (in pixel space).
  - finalY: The Y coordinate where the ray hits a wall (in pixel space).
*/
func CalculateRay(x, screenWidth, screenHeight int, centerX, centerY float64, m *world.Map, p *entity.Player) (float64, float64, float64, int, float64, byte, int) {
	cameraX := (2.0 * float64(x) / float64(screenWidth)) - 1.0

	rayDirX := p.DirX + (p.PlaneX * cameraX)
	rayDirY := p.DirY + (p.PlaneY * cameraX)

	playerTileX := centerX / float64(world.TileSize)
	playerTileY := centerY / float64(world.TileSize)

	deltaDistX, stepX := getDeltaDistX(rayDirX)
	deltaDistY, stepY := getDeltaDistY(rayDirY)

	sideDistX, mapX := getSideDistX(playerTileX, stepX, deltaDistX)
	sideDistY, mapY := getSideDistY(playerTileY, stepY, deltaDistY)

	hit := false
	side := 0 // 0 = hit X, 1 = hit Y

	var tileType byte

	for !hit {
		if sideDistX < sideDistY {
			sideDistX += deltaDistX
			mapX += stepX
			side = 0
		} else {
			sideDistY += deltaDistY
			mapY += stepY
			side = 1
		}

		if mapX >= 0 && mapX < m.Width && mapY >= 0 && mapY < m.Height {
			mapIndex := (mapY * m.Width) + mapX
			if m.Grid[mapIndex] > 0 {
				tileType = m.Grid[mapIndex]
				hit = true
			}
		} else {
			// out of bounds
			tileType = 1
			hit = true
		}
	}

	var totalDistance float64
	if side == 0 {
		totalDistance = sideDistX - deltaDistX
	} else {
		totalDistance = sideDistY - deltaDistY
	}

	finalX := centerX + (rayDirX * totalDistance * float64(world.TileSize))
	finalY := centerY + (rayDirY * totalDistance * float64(world.TileSize))

	var perpWallDist float64
	if side == 0 {
		perpWallDist = sideDistX - deltaDistX
	} else {
		perpWallDist = sideDistY - deltaDistY
	}

	var wallX float64
	if side == 0 {
		// vertical walls |
		wallX = playerTileY + perpWallDist*rayDirY
	} else {
		// horizontal walls --
		wallX = playerTileX + perpWallDist*rayDirX
	}
	// get the fractional part of wall (eg.: 14.73 - 14 = 0.73 -> 73%)
	wallX -= math.Floor(wallX)

	// 64px
	textureWidth := world.TileSize
	textureX := int(wallX * float64(textureWidth)) // eg. 0.73 * 64 = 46.72 -> 46th pixel

	// if the wall is facing left or up, flip the texture coordinate
	if side == 0 && rayDirX > 0 {
		textureX = textureWidth - textureX - 1
	}
	if side == 1 && rayDirY < 0 {
		textureX = textureWidth - textureX - 1
	}

	wallHeight := (float64(screenHeight) / perpWallDist) * HeightMultiplier

	return finalX, finalY, wallHeight, side, perpWallDist, tileType, textureX
}

/*
[CalculateAllRays] casts rays for each vertical stripe of the screen to determine wall intersections and their properties.

Parameters:
  - step: The interval at which rays are cast (e.g., every 4 pixels).
  - screenWidth: The width of the screen in pixels.
  - screenHeight: The height of the screen in pixels.
  - centerX: The player's X coordinate in pixel space.
  - centerY: The player's Y coordinate in pixel space.
  - rays: A slice to store the results of each raycast.
  - m: Pointer to the world Map containing the grid layout.
  - p: Pointer to the Player entity containing position and direction information.

Returns:
  - A slice of [RayResult] containing the properties of each raycast for rendering.
*/
func CalculateAllRays(screenWidth, screenHeight int, centerX, centerY float64, rays []RayResult, m *world.Map, p *entity.Player) []RayResult {
	// the slice object/header is copied, but the array is a pointer
	// set Len to 0
	rays = rays[:0]

	for x := range screenWidth {
		finalX, finalY, wallHeight, side, perpWallDist, tileType, textureX := CalculateRay(x, screenWidth, screenHeight, centerX, centerY, m, p)
		// overwrite old values
		rays = append(rays, RayResult{
			X:            x,
			FinalX:       finalX,
			FinalY:       finalY,
			WallHeight:   wallHeight,
			Side:         side,
			PerpWallDist: perpWallDist,
			TextureX:     textureX,
			TileType:     tileType,
		})
	}

	return rays
}
