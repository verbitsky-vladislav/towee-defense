package setup

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"tower-defense/internal/gameobjects/enemy"
	"tower-defense/internal/objects/game"
)

type Game struct {
	skeleton *game.Enemy // Экземпляр скелета
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Init() error {
	// Инициализируем анимации скелета
	if err := enemy.InitSkeletonAnimations(); err != nil {
		log.Fatalf(err.Error())
		return err
	}

	// Создаем экземпляр скелета
	g.skeleton = enemy.NewSkeleton()
	return nil
}

func (g *Game) Update() error {
	// Если скелет существует, обновляем анимацию
	if g.skeleton != nil {
		g.skeleton.TypeInfo.Animations[game.Move].Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.skeleton != nil {
		ebitenutil.DebugPrint(screen, "Drawing animation...\n")
		g.skeleton.TypeInfo.Animations[game.Move].Draw(screen, 100, 200)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240 // Примерные размеры экрана
}
