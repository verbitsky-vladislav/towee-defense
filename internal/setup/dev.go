package setup

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"tower-defense/internal/gameobjects/enemy"
	level2 "tower-defense/internal/gameobjects/level"
	"tower-defense/internal/objects/base"
	"tower-defense/internal/objects/game"
)

var levelMap = [][]base.Tile{
	{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

type Game struct {
	level    *game.Level
	skeleton *game.Enemy

	width, height int
}

func NewGame() *Game {

	level := level2.NewLevel(levelMap, "forest", 16, 16)

	err := enemy.InitSkeletonAnimations()
	if err != nil {
		log.Fatal(err)
	}
	skeleton := enemy.NewSkeleton()

	return &Game{
		level:    level,
		skeleton: skeleton,

		width:  640,
		height: 480,
	}
}

func (g *Game) Update() error {
	g.skeleton.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scaleX := float64(g.width) / float64(g.level.W*g.level.TileW)
	scaleY := float64(g.height) / float64(g.level.H*g.level.TileH)
	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	offsetX := (float64(g.width) - float64(g.level.W*g.level.TileW)*scale) / 2
	offsetY := (float64(g.height) - float64(g.level.H*g.level.TileH)*scale) / 2

	g.level.Draw(screen, scale, offsetX, offsetY)

	g.skeleton.Draw(screen, 100*scale+offsetX, 100*scale+offsetY, scale)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.width = outsideWidth
	g.height = outsideHeight

	return outsideWidth, outsideHeight
}
