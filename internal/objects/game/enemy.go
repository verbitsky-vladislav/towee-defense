package game

import (
	"tower-defense/internal/gameobjects/enemy"
	"tower-defense/internal/system/animation"
)

type EnemyType = int

const (
	Skeleton EnemyType = iota
)

type EnemyTypeInfo struct {
	Speed int

	Animations map[enemy.Animation]*animation.Animation
}

type Enemy struct {
	Type     EnemyType // Тип врага -> конкретная информация указана
	TypeInfo any
	Health   int64
}
