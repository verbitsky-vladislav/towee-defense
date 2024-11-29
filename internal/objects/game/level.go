package game

type TowerLevelProps struct {
	Tower  *Tower
	Target *Enemy
}

type EnemyLevelProps struct {
	Enemy *Enemy
	Path  int64
}

type Level struct {
	Money int64
	Path  []Point
	X, Y  int64

	Turrets map[Point]*TowerLevelProps
	Enimies []*EnemyLevelProps
}
