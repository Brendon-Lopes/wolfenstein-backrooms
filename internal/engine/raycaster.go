package engine

import (
	"math"

	"github.com/Brendon-Lopes/wolfenstein-backrooms/internal/world"
)

/*
[GetDeltaDistX] calculates the length of the ray's hypotenuse when it travels exactly one map grid unit along the Y-axis.

Returns:
  - deltaDistX: The absolute length of the hypotenuse for a full grid step.
  - stepX: The direction to step on the map grid (-1 for left or 1 right).
*/
func GetDeltaDistX(rayAngle float64) (float64, int) {
	var stepX int
	rayDir := math.Cos(rayAngle)

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
[GetDeltaDistY] calculates the length of the ray's hypotenuse when it travels exactly one map grid unit along the Y-axis.

Returns:
  - deltaDistY: The absolute length of the hypotenuse for a full grid step.
  - stepY: The direction to step on the map grid (-1 for up or 1 for down).
*/
func GetDeltaDistY(rayAngle float64) (float64, int) {
	var stepY int
	rayDir := math.Sin(rayAngle)

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
[GetSideDistX] calculates the initial ray distance from the player's exact position to the first vertical grid line intersection.

Parameters:
  - playerTileX: The player's exact X coordinate in logical map units (e.g., 1.2).
  - stepX: The ray's X-axis direction (-1 or 1).
  - deltaDistX: The hypotenuse length of a full block step.

Returns:
  - sideDistX: The total ray length to reach the first vertical grid boundary.
  - mapX: The integer map coordinate where the player is currently standing.
*/
func GetSideDistX(playerTileX float64, stepX int, deltaDistX float64) (float64, int) {
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
[GetSideDistY] calculates the initial ray distance from the player's exact position to the first horizontal grid line intersection.

Parameters:
  - playerTileY: The player's exact Y coordinate in logical map units (e.g., 1.2).
  - stepY: The ray's Y-axis direction (-1 or 1).
  - deltaDistY: The hypotenuse length of a full block step.

Returns:
  - sideDistY: The total ray length to reach the first horizontal grid boundary.
  - mapY: The integer map coordinate where the player is currently standing.
*/
func GetSideDistY(playerTileY float64, stepY int, deltaDistY float64) (float64, int) {
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
func CalculateRay(centerX, centerY, angle float64, m *world.Map) (float64, float64) {
	playerTileX := centerX / float64(TileSize)
	playerTileY := centerY / float64(TileSize)

	deltaDistX, stepX := GetDeltaDistX(angle)
	deltaDistY, stepY := GetDeltaDistY(angle)

	sideDistX, mapX := GetSideDistX(playerTileX, stepX, deltaDistX)
	sideDistY, mapY := GetSideDistY(playerTileY, stepY, deltaDistY)

	hit := false
	side := 0 // 0 = hit X, 1 = hit Y

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
			if m.Grid[mapIndex] == 1 {
				hit = true
			}
		} else {
			// out of bounds
			hit = true
		}
	}

	var totalDistance float64
	if side == 0 {
		totalDistance = sideDistX - deltaDistX
	} else {
		totalDistance = sideDistY - deltaDistY
	}

	rayDirX := math.Cos(angle)
	rayDirY := math.Sin(angle)

	finalX := centerX + (rayDirX * totalDistance * float64(TileSize))
	finalY := centerY + (rayDirY * totalDistance * float64(TileSize))

	return finalX, finalY
}
