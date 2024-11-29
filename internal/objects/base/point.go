package base

type Point struct {
	X, Y int32
}

type Tile int32

const (
	Ground    Tile = iota
	Path      Tile = iota
	TowerProp Tile = iota
)
