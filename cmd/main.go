package main

import (
	"log"
	"tower-defense/internal/setup"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&setup.Game{}); err != nil {
		log.Fatal(err)
	}
}
