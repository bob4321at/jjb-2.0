package enemyai

import (
	"jjb/textures"
	"jjb/utils"
	"math"
)

func simpleGradeCurseUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	enemy.Vel.Y += 0.1

	if player_pos.X > enemy.Pos.X {
		enemy.Dir = false
		enemy.Vel.X = 4
	} else if player_pos.X < enemy.Pos.X {
		enemy.Dir = true
		enemy.Vel.X = -4
	}
	for ei := 0; ei < len(Enemies_In_World); ei++ {
		oe := Enemies_In_World[ei]
		if oe != enemy {
			if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: oe.Pos.X, Y: oe.Pos.Y}, utils.Vec2{X: float64(oe.Tex.GetTexture().Bounds().Dx()), Y: float64(oe.Tex.GetTexture().Bounds().Dy())}) {
				if enemy.Pos.X > oe.Pos.X {
					enemy.Vel.X = 1
					for ti := 0; ti < len(level_hitbox); ti++ {
						t := level_hitbox[ti]
						if utils.Collide(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + enemy.Vel.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
							enemy.Vel.Y = 0
						}
						if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
							enemy.Vel.X = 0
						}
					}
				} else {
					enemy.Vel.X = -1
					for ti := 0; ti < len(level_hitbox); ti++ {
						t := level_hitbox[ti]
						if utils.Collide(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + enemy.Vel.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
							enemy.Vel.Y = 0
						}
						if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
							enemy.Vel.X = 0
						}
					}
				}
			}
		}
	}

	for ti := 0; ti < len(level_hitbox); ti++ {
		t := level_hitbox[ti]
		if utils.Collide(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + enemy.Vel.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			enemy.Vel.Y = 0
		}
		if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			enemy.Vel.X = 0
		}
	}

	if utils.Collide(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + enemy.Vel.Y + 2}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		enemy.Vel.Y = 0
	}

	if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		enemy.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: 2000 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		enemy.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: enemy.Pos.X + enemy.Vel.X, Y: enemy.Pos.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: 2000 + 2048 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		enemy.Vel.X = 0
	}

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
		check := int(math.Mod(utils.Game_Time, 100))
		if check == 0 {
			if enemy.Dir {
				enemy.NewProjectile(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y - 32}, utils.Vec2{X: -7, Y: 0}, textures.NewTexture("./art/enemies/simple_grade_curse_attack_left.png", ""), 5, 10)
			} else {
				enemy.NewProjectile(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y - 32}, utils.Vec2{X: 7, Y: 0}, textures.NewTexture("./art/enemies/simple_grade_curse_attack_right.png", ""), 5, 10)
			}
		}
	}
}
