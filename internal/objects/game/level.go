package game

import "tower-defense/internal/objects/base"

type TowerLevelProps struct {
	Tower  *Tower
	Target *Enemy
}

type EnemyLevelProps struct {
	Enemy *Enemy
	Path  int64
}

type Level struct {
	Money         int64
	Path          []base.Point
	X, Y          int64
	Width, Height int64

	Turrets map[base.Point]*TowerLevelProps
	Enemies []*EnemyLevelProps
}
