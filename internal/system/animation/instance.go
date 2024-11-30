package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Instance struct {
	Resource       *Resource     // Ссылка на общий ресурс анимации
	FrameIndex     int           // Текущий кадр
	FrameDelay     time.Duration // Задержка между кадрами
	LastFrameTime  time.Time     // Время последнего обновления
	DirectionIndex int           // Текущая строка спрайтов (направление)
}

// NewAnimationInstance создает новый экземпляр анимации
func NewAnimationInstance(resource *Resource, frameDelay time.Duration) *Instance {
	return &Instance{
		Resource:       resource,
		FrameIndex:     0,
		FrameDelay:     frameDelay,
		LastFrameTime:  time.Now(),
		DirectionIndex: 0,
	}
}

// Update обновляет состояние экземпляра
func (ai *Instance) Update() {
	now := time.Now()
	if now.Sub(ai.LastFrameTime) >= ai.FrameDelay {
		ai.LastFrameTime = now
		ai.FrameIndex = (ai.FrameIndex + 1) % ai.Resource.FrameCount
	}
}

// Draw рисует текущий кадр экземпляра анимации
func (ai *Instance) Draw(screen *ebiten.Image, x, y float64, scale float64) {
	frame := ai.Resource.Cache[ai.DirectionIndex][ai.FrameIndex]
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale) // Масштабируем врага
	op.GeoM.Translate(x, y)
	op.GeoM.Translate(x, y)
	screen.DrawImage(frame, op)
}

// SetDirection изменяет направление анимации (строка в спрайт-листе)
func (ai *Instance) SetDirection(direction int) {
	if direction >= 0 && direction < len(ai.Resource.Cache) {
		ai.DirectionIndex = direction
	}
}
