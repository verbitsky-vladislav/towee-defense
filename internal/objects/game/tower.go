package game

type AttackTactic int

const (
	ClosestEnemy   AttackTactic = iota // Атаковать ближайшего врага
	FurthestEnemy                      // Атаковать самого дальнего врага
	StrongestEnemy                     // Атаковать самого сильного врага (с максимальным здоровьем)
	WeakestEnemy                       // Атаковать самого слабого врага (с минимальным здоровьем)
)

type TowerType int

const (
	ArcherTower TowerType = iota // as example
)

// DamageType определяет тип урона башни.
type DamageType int

const (
	SingleTarget DamageType = iota // Урон по одной цели
	AreaOfEffect                   // Урон по области (AOE)
)

type Tower struct {
	Type         TowerType
	AttackTactic AttackTactic
	DamageType   DamageType

	AttackDamage int64
	AttackSpeed  int64
	AttackRange  float64
}
