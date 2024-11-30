package enemy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"time"
	"tower-defense/internal/objects/game"
	"tower-defense/internal/system/animation"
	"tower-defense/internal/system/spriteloader"
)

var skeletonAnimations = make(map[game.Animation]*animation.Resource)

func InitSkeletonAnimations() error {
	animationConfigs := map[string]spriteloader.SpriteConfig{
		"Idle": {
			FileName:    "idle.png",
			FrameWidth:  16,
			FrameHeight: 16,
			FrameCount:  4,
			Rows:        1,
		},
		"Move": {
			FileName:    "move.png",
			FrameWidth:  16,
			FrameHeight: 16,
			FrameCount:  8,
			Rows:        1,
		},
		"Attack": {
			FileName:    "attack.png",
			FrameWidth:  16,
			FrameHeight: 16,
			FrameCount:  7,
			Rows:        1,
		},
		"Hit": {
			FileName:    "hit.png",
			FrameWidth:  16,
			FrameHeight: 16,
			FrameCount:  4,
			Rows:        1,
		},
		"Death": {
			FileName:    "death.png",
			FrameWidth:  16,
			FrameHeight: 16,
			FrameCount:  10,
			Rows:        1,
		},
	}

	animations, err := spriteloader.LoadSprites("assets/enemy/skeleton/animations/", animationConfigs)
	if err != nil {
		log.Fatalf("Ошибка загрузки анимаций: %v", err)
		return err
	}

	for name, data := range animations {
		switch name {
		case "Idle":
			skeletonAnimations[game.Idle] = animation.NewAnimationResource(data.Image, data.FrameWidth, data.FrameHeight, data.FrameCount, data.Rows)
		case "Move":
			skeletonAnimations[game.Move] = animation.NewAnimationResource(data.Image, data.FrameWidth, data.FrameHeight, data.FrameCount, data.Rows)
		case "Attack":
			skeletonAnimations[game.Attack] = animation.NewAnimationResource(data.Image, data.FrameWidth, data.FrameHeight, data.FrameCount, data.Rows)
		case "Hit":
			skeletonAnimations[game.Hit] = animation.NewAnimationResource(data.Image, data.FrameWidth, data.FrameHeight, data.FrameCount, data.Rows)
		case "Death":
			skeletonAnimations[game.Death] = animation.NewAnimationResource(data.Image, data.FrameWidth, data.FrameHeight, data.FrameCount, data.Rows)
		}
	}

	return nil
}

type Skeleton struct {
	Enemy *game.Enemy
}

func newSkeletonTypeInfo() *game.EnemyTypeInfo {
	// Сопоставление анимаций для скелета
	var skAnimations = map[game.Animation]*animation.Instance{
		game.Idle:   animation.NewAnimationInstance(skeletonAnimations[game.Idle], 150*time.Millisecond),
		game.Move:   animation.NewAnimationInstance(skeletonAnimations[game.Move], 100*time.Millisecond),
		game.Attack: animation.NewAnimationInstance(skeletonAnimations[game.Attack], 200*time.Millisecond),
		game.Hit:    animation.NewAnimationInstance(skeletonAnimations[game.Hit], 80*time.Millisecond),
		game.Death:  animation.NewAnimationInstance(skeletonAnimations[game.Death], 300*time.Millisecond),
	}

	return &game.EnemyTypeInfo{
		Speed:      1.5,
		Animations: skAnimations,
	}
}

func NewSkeleton() *game.Enemy {
	typeInfo := newSkeletonTypeInfo()

	return &game.Enemy{
		Type:     game.Skeleton,
		TypeInfo: *typeInfo,
		Health:   100,
	}
}

func (s *Skeleton) Update() {
	for _, animInstance := range s.Enemy.TypeInfo.Animations {
		animInstance.Update()
	}
}

func (s *Skeleton) Draw(screen *ebiten.Image, x, y float64, scale float64) {
	if animInstance, exists := s.Enemy.TypeInfo.Animations[game.Idle]; exists {
		animInstance.Draw(screen, x, y, scale)
	}
}

func (s *Skeleton) TakeDamage(amount int64) {
	s.Enemy.Health -= amount
	if s.Enemy.Health < 0 {
		s.Enemy.Health = 0
	}
}

func (s *Skeleton) IsDead() bool {
	return s.Enemy.Health <= 0
}
