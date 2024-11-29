package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"tower-defense/internal/gameobjects/enemy"
	"tower-defense/internal/objects/game"
)

type Game struct {
	skeleton *game.Enemy
}

func (g *Game) Update() error {
	g.skeleton.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.skeleton.Draw(screen, 100, 100)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	err := enemy.InitSkeletonAnimations()
	if err != nil {
		log.Fatal(err)
	}

	skeleton := enemy.NewSkeleton()

	g := &Game{
		skeleton: skeleton,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
