package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"tower-defense/internal/setup"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(setup.NewGame()); err != nil {
		log.Fatal(err)
	}
}
