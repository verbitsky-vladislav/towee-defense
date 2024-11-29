package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tower-defense/internal/system/animation"
)

type Animation string

const (
	Idle   = Animation("Idle")
	Move   = Animation("Move")
	Attack = Animation("Attack")
	Hit    = Animation("Hit")
	Death  = Animation("Death")
)

type EnemyType = int

const (
	Skeleton EnemyType = iota
)

type EnemyTypeInfo struct {
	Speed      float64
	Animations map[Animation]*animation.Instance
}

type Enemy struct {
	Type     EnemyType // Тип врага -> конкретная информация указана
	TypeInfo EnemyTypeInfo
	Health   int64
}

func (e *Enemy) Update() {
	e.TypeInfo.Animations[Idle].Update()
}

func (e *Enemy) Draw(screen *ebiten.Image, x, y float64) {
	e.TypeInfo.Animations[Idle].Draw(screen, x, y)
}
