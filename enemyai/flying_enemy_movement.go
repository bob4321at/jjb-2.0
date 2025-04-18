package enemyai

import (
	"jjb/utils"
	"math"
)

func flyingEnemyMovement(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox, speed float64, height_above_ground float64) {
	enemy.Vel.X += -0.015 * (enemy.Pos.X - player_pos.X) * (math.Abs(enemy.Pos.Y / 100)) / 20

	if enemy.Pos.Y > player_pos.Y-height_above_ground {
		enemy.Vel.Y -= 1
	}

	if enemy.Vel.X > speed {
		enemy.Vel.X = speed
	} else if enemy.Vel.X < -speed {
		enemy.Vel.X = -speed
	}

	enemy.Vel.Y += 0.5

	for le := 0; le < len(Enemies_In_World); le++ {
		oe := Enemies_In_World[le]
		if enemy != oe {
			if utils.Collide(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + enemy.Vel.Y}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}, utils.Vec2{X: float64(oe.Pos.X), Y: float64(oe.Pos.Y)}, utils.Vec2{X: float64(oe.Tex.GetTexture().Bounds().Dx()), Y: float64(oe.Tex.GetTexture().Bounds().Dy())}) {
				enemy.Vel.Y = -enemy.Vel.Y / 1.5
			}
			if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}, utils.Vec2{X: float64(oe.Pos.X), Y: float64(oe.Pos.Y)}, utils.Vec2{X: float64(oe.Tex.GetTexture().Bounds().Dx()), Y: float64(oe.Tex.GetTexture().Bounds().Dy())}) {
				enemy.Vel.X = -enemy.Vel.X / 1.5
			}
		}
	}

	for ti := 0; ti < len(level_hitbox); ti++ {
		t := level_hitbox[ti]
		if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			enemy.Vel.X = -enemy.Vel.X / 1.5
		}
		if utils.Collide(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + enemy.Vel.Y}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy()) + height_above_ground}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			enemy.Vel.Y -= 1
		}
		if utils.Collide(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + enemy.Vel.Y}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			enemy.Vel.Y = 0
		}
	}

	if utils.Collide(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + enemy.Vel.Y + 2}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		enemy.Vel.Y = 0
	}
	if utils.Collide(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + enemy.Vel.Y + 2}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) - (575/2 - 32)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		enemy.Vel.Y = 0
	}

	if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		enemy.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}, utils.Vec2{X: 2000 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		enemy.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}, utils.Vec2{X: 2000 + 2048 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		enemy.Vel.X = 0
	}
}
