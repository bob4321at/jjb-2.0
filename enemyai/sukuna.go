package enemyai

import (
	"jjb/textures"
	"jjb/utils"
	"math"
)

func sukunaUpdate(e *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	e.Vel.Y += 0.1

	if player_pos.X > e.Pos.X {
		e.Dir = false
		e.Vel.X = 4
	} else if player_pos.X < e.Pos.X {
		e.Dir = true
		e.Vel.X = -4
	}

	for ei := 0; ei < len(Enemies_In_World); ei++ {
		oe := Enemies_In_World[ei]
		if oe != e {
			if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: oe.Pos.X, Y: oe.Pos.Y}, utils.Vec2{X: float64(oe.Tex.GetTexture().Bounds().Dx()), Y: float64(oe.Tex.GetTexture().Bounds().Dy())}) {
				if e.Pos.X > oe.Pos.X {
					e.Vel.X = 1
					for ti := 0; ti < len(level_hitbox); ti++ {
						t := level_hitbox[ti]
						if utils.Collide(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + e.Vel.Y}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
							e.Vel.Y = 0
						}
						if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
							e.Vel.X = 0
						}
					}
				} else {
					e.Vel.X = -1
					for ti := 0; ti < len(level_hitbox); ti++ {
						t := level_hitbox[ti]
						if utils.Collide(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + e.Vel.Y}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
							e.Vel.Y = 0
						}
						if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
							e.Vel.X = 0
						}
					}
				}
			}
		}
	}

	for ti := 0; ti < len(level_hitbox); ti++ {
		t := level_hitbox[ti]
		if utils.Collide(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + e.Vel.Y}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			e.Vel.Y = 0
		}
		if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			e.Vel.X = 0
		}
	}

	if utils.Collide(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + e.Vel.Y + 2}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		e.Vel.Y = 0
	}

	if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		e.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: 2000 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		e.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 128}, utils.Vec2{X: 2000 + 2048 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		e.Vel.X = 0
	}

	if e.Can_Move {
		e.Pos.X += e.Vel.X
		e.Pos.Y += e.Vel.Y
		check := int(math.Mod(utils.Game_Time, 100))
		if check == 0 {
			if math.Abs(player_pos.X-e.Pos.X) > 256 {
				if e.Dir {
					e.NewProjectile(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + 32}, utils.Vec2{X: -5, Y: 0}, textures.NewTexture("./art/enemies/fuego_left.png", ""), 10, 100)
				} else {
					e.NewProjectile(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + 32}, utils.Vec2{X: 5, Y: 0}, textures.NewTexture("./art/enemies/fuego_right.png", ""), 10, 100)
				}
			} else {
				if e.Dir {
					e.NewProjectile(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + 32}, utils.Vec2{X: -6, Y: 0}, textures.NewTexture("./art/enemies/sukuna_attack_cut_left.png", ""), 5, 20)
				} else {
					e.NewProjectile(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + 32}, utils.Vec2{X: 6, Y: 0}, textures.NewTexture("./art/enemies/sukuna_attack_cut_right.png", ""), 5, 20)
				}
			}
		}
	}
}
