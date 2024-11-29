package spriteloader

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// SpriteData содержит информацию об одном спрайт-листе
type SpriteData struct {
	Image       *ebiten.Image // Загруженное изображение
	FrameWidth  int           // Ширина кадра
	FrameHeight int           // Высота кадра
	FrameCount  int           // Количество кадров
	Rows        int           // Количество строк (направлений)
}

// LoadSprites загружает несколько анимаций из заданных путей
func LoadSprites(basePath string, animationConfigs map[string]SpriteConfig) (map[string]*SpriteData, error) {
	animations := make(map[string]*SpriteData)

	for name, config := range animationConfigs {
		image, err := LoadImage(basePath + config.FileName)
		if err != nil {
			return nil, fmt.Errorf("failed to load animation %s: %w", name, err)
		}

		animations[name] = &SpriteData{
			Image:       image,
			FrameWidth:  config.FrameWidth,
			FrameHeight: config.FrameHeight,
			FrameCount:  config.FrameCount,
			Rows:        config.Rows,
		}
	}

	return animations, nil
}

// SpriteConfig описывает конфигурацию для одной анимации
type SpriteConfig struct {
	FileName    string // Имя файла спрайт-листа
	FrameWidth  int    // Ширина кадра
	FrameHeight int    // Высота кадра
	FrameCount  int    // Количество кадров
	Rows        int    // Количество строк (направлений)
}

// LoadImage загружает изображение из файла
func LoadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot load image with path: %v: %w", path, err)
	}
	return img, nil
}
