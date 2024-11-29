package spriteloader

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadAnimations(basePath string, animationNames []string) (map[string]*ebiten.Image, error) {
	animations := make(map[string]*ebiten.Image)

	for _, animName := range animationNames {
		img := LoadImage(basePath + animName)
		//if err != nil {
		//	return nil, err
		//}
		animations[animName] = img
	}

	return animations, nil
}

// LoadImage helper function to load an image from a file path
func LoadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(fmt.Errorf(`cannot load image with path: %v`, path))
		return nil
	}
	return img
}
