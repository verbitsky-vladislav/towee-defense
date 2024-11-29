package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Resource struct {
	Sprite      *ebiten.Image     // Спрайт-лист
	FrameWidth  int               // Ширина кадра
	FrameHeight int               // Высота кадра
	FrameCount  int               // Количество кадров
	Cache       [][]*ebiten.Image // Кэш кадров [строка][кадр]
}

// NewAnimationResource создает новый объект ресурса анимации
func NewAnimationResource(sprite *ebiten.Image, frameWidth, frameHeight, frameCount, rows int) *Resource {
	resource := &Resource{
		Sprite:      sprite,
		FrameWidth:  frameWidth,
		FrameHeight: frameHeight,
		FrameCount:  frameCount,
		Cache:       make([][]*ebiten.Image, rows),
	}

	// Кэшируем кадры
	for row := 0; row < rows; row++ {
		resource.Cache[row] = make([]*ebiten.Image, frameCount)
		for frame := 0; frame < frameCount; frame++ {
			x0 := frame * frameWidth
			y0 := row * frameHeight
			x1 := x0 + frameWidth
			y1 := y0 + frameHeight
			frameImage := sprite.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
			resource.Cache[row][frame] = frameImage
		}
	}

	return resource
}
