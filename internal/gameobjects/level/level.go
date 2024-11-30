package level

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tower-defense/internal/gameobjects/tilemap"
	"tower-defense/internal/objects/base"
	"tower-defense/internal/objects/game"
)

func NewLevel(mapData [][]base.Tile, levelType game.LevelType, tileW, tileH int32) *game.Level {

	tiles := make(map[base.Tile]*ebiten.Image)
	switch levelType {
	case game.Forest:
		tiles = tilemap.NewForest().Tiles
	default:
		tiles = tilemap.NewForest().Tiles
	}

	w := int32(len(mapData[0]))
	h := int32(len(mapData))

	level := &game.Level{
		Map:      mapData,
		Path:     realizePath(mapData),
		Tiles:    tiles,
		MapCache: ebiten.NewImage(int(w*32), int(h*32)),

		W: w, H: h,
		TileW: tileW, TileH: tileH,

		Towers:  nil, // todo add
		Enemies: nil, // todo add
	}
	level.CacheMap()
	return level
}

// todo think about name of func
func realizePath(mapData [][]base.Tile) []base.Point {
	var path []base.Point

	for y, row := range mapData {
		for x, tile := range row {
			if tile == base.Path {
				path = append(path, base.Point{X: int32(x), Y: int32(y)})
			}
		}
	}

	return path
}
