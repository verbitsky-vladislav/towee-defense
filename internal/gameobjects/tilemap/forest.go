package tilemap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tower-defense/internal/objects/base"
	"tower-defense/internal/system/spriteloader"
)

type ForestTile base.Tile

const (
	ForestGround    ForestTile = 0
	ForestPath                 = 1
	ForestTowerProp            = 2
)

var allForestTiles = map[ForestTile]string{
	0: "ground.png",
	1: "path.png",
	2: "tower-prop.png",
}

type Forest struct {
	Tiles map[base.Tile]*ebiten.Image
}

func NewForest() *Forest {

	var forestPath = "assets/map/forest/"
	var tiles = make(map[base.Tile]*ebiten.Image)

	for k, tile := range allForestTiles {
		sprite, err := spriteloader.LoadImage(forestPath + tile)
		if err != nil {
			panic(err) // todo add error handler
		}
		tiles[base.Tile(k)] = sprite
	}

	return &Forest{
		Tiles: tiles,
	}
}
