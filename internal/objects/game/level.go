package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tower-defense/internal/objects/base"
)

type TowerLevelProps struct {
	Tower  *Tower
	Target *Enemy
}

type EnemyLevelProps struct {
	Enemy *Enemy
	Path  int64
}

type LevelType string

const (
	Forest LevelType = "forest"
)

type Level struct {
	Map  [][]base.Tile
	Path []base.Point

	W, H         int32
	TileW, TileH int32

	MapCache *ebiten.Image

	Tiles map[base.Tile]*ebiten.Image

	Towers  map[base.Point]*TowerLevelProps
	Enemies []*EnemyLevelProps
}

func (l *Level) Draw(screen *ebiten.Image, scale float64, offsetX, offsetY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)         // Применяем масштаб
	op.GeoM.Translate(offsetX, offsetY) // Центрируем карту
	screen.DrawImage(l.MapCache, op)
}

func (l *Level) CacheMap() {
	if l.MapCache == nil {
		l.MapCache = ebiten.NewImage(int(l.W*l.TileW), int(l.H*l.TileH))
	}

	for y, row := range l.Map {
		for x, tile := range row {
			if img, ok := l.Tiles[tile]; ok {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*int(l.TileW)), float64(y*int(l.TileH)))
				l.MapCache.DrawImage(img, op)
			}
		}
	}
}
