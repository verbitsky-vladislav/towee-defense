package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"time"
)

// todo add cacheImages
// todo optimize parameters in Animation struct
// todo delete x,y,direction controllers from Animation struct (logic separation)

type Animation struct {
	Sprite        *ebiten.Image // Изображение, содержащее спрайт-лист анимации
	FrameWidth    int           // Ширина одного кадра анимации
	FrameHeight   int           // Высота одного кадра анимации
	FrameCount    int           // Общее количество кадров в анимации
	FrameDelay    time.Duration // Задержка между сменой кадров
	LastFrameTime time.Time     // Время последнего обновления кадра
	FrameIndex    int           // Индекс текущего кадра анимации

	X, Y      *float64 // Указатели на координаты x и y объекта
	Direction *int     // Указатель на направление движения объекта
}

func NewAnimation(
	sprite *ebiten.Image,
	frameWidth, frameHeight, frameCount int,
	frameDelay time.Duration,
	lastFrameTime time.Time,
	frameIndex int,
	x, y *float64,
	direction *int,
) *Animation {
	return &Animation{
		Sprite:        sprite,
		FrameWidth:    frameWidth,
		FrameHeight:   frameHeight,
		FrameCount:    frameCount,
		FrameDelay:    frameDelay,
		LastFrameTime: lastFrameTime,
		FrameIndex:    frameIndex,

		X:         x,
		Y:         y,
		Direction: direction,
	}
}

// Update обновляет состояние анимации, переходя к следующему кадру, если прошла заданная задержка.
func (a *Animation) Update() {
	now := time.Now()

	// Если прошло достаточно времени с последнего обновления, переключить кадр
	if now.Sub(a.LastFrameTime) >= a.FrameDelay {
		a.LastFrameTime = now
		a.FrameIndex = (a.FrameIndex + 1) % a.FrameCount // Переход на следующий кадр, цикл по индексам
	}
}

// DrawLine анимирует объект по line spritу - спрайту, который идет слева направо
func (a *Animation) DrawLine(screen *ebiten.Image) {
	x0 := a.FrameIndex * a.FrameWidth
	y0 := 0
	x1 := x0 + a.FrameWidth
	y1 := y0 + a.FrameHeight
	frame := image.Rect(x0, y0, x1, y1)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(*a.X, *a.Y)

	screen.DrawImage(a.Sprite.SubImage(frame).(*ebiten.Image), op)
}

func (a *Animation) DrawGrid(screen *ebiten.Image) {
	// Расчет позиции кадра на спрайте (как по двумерному массиву считаем)
	x0 := *a.Direction * a.FrameWidth
	y0 := a.FrameIndex * a.FrameHeight

	x1 := x0 + a.FrameWidth
	y1 := y0 + a.FrameHeight

	frame := image.Rect(x0, y0, x1, y1)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(*a.X, *a.Y)

	screen.DrawImage(a.Sprite.SubImage(frame).(*ebiten.Image), op)
}
