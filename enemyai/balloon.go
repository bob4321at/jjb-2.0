package enemyai

import (
	"jjb/textures"
	"jjb/utils"
	"math"
)

func balloonUpdate(e *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	if e.Health <= 0 {
		Enemies_To_Add = append(Enemies_To_Add, NewEnemy(5, 10, 5, e.Pos, textures.NewTexture("./art/enemies/small_balloon.png", ""), smallBalloonUpdate))
	}

	e.Vel.X += -0.015 * (e.Pos.X - player_pos.X) * (math.Abs(e.Pos.Y / 100)) / 20

	if e.Pos.Y > player_pos.Y-128 {
		e.Vel.Y -= 1
	}

	if e.Vel.X > 10 {
		e.Vel.X = 10
	} else if e.Vel.X < -10 {
		e.Vel.X = -10
	}

	e.Vel.Y += 0.5

	for le := 0; le < len(Enemies_In_World); le++ {
		oe := Enemies_In_World[le]
		if e != oe {
			if utils.Collide(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + e.Vel.Y}, utils.Vec2{X: 64, Y: 64}, utils.Vec2{X: float64(oe.Pos.X), Y: float64(oe.Pos.Y)}, utils.Vec2{X: float64(oe.Tex.GetTexture().Bounds().Dx()), Y: float64(oe.Tex.GetTexture().Bounds().Dy())}) {
				e.Vel.Y = -e.Vel.Y / 1.2
				oe.Vel.Y = e.Vel.Y / 1.2
			}
			if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 64}, utils.Vec2{X: float64(oe.Pos.X), Y: float64(oe.Pos.Y)}, utils.Vec2{X: float64(oe.Tex.GetTexture().Bounds().Dx()), Y: float64(oe.Tex.GetTexture().Bounds().Dy())}) {
				e.Vel.X = -e.Vel.X / 1.2
				oe.Vel.X = e.Vel.X / 1.2
			}
		}
	}

	for ti := 0; ti < len(level_hitbox); ti++ {
		t := level_hitbox[ti]
		if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 64}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			e.Vel.X = -e.Vel.X / 1.5
		}
		if utils.Collide(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + e.Vel.Y}, utils.Vec2{X: 64, Y: 64}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			e.Vel.Y = 0
		}
		if utils.Collide(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + e.Vel.Y}, utils.Vec2{X: 64, Y: 64}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			e.Vel.Y = 0
		}
	}
	if utils.Collide(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + e.Vel.Y + 2}, utils.Vec2{X: 64, Y: 64}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		e.Vel.Y = 0
	}
	if utils.Collide(utils.Vec2{X: e.Pos.X, Y: e.Pos.Y + e.Vel.Y + 2}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) - (575/2 - 32)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		e.Vel.Y = 0
	}

	if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 64}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		e.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 64}, utils.Vec2{X: 2000 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		e.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: e.Pos.X + e.Vel.X, Y: e.Pos.Y}, utils.Vec2{X: 64, Y: 64}, utils.Vec2{X: 2000 + 2048 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		e.Vel.X = 0
	}

	if e.Can_Move {
		e.Pos.X += e.Vel.X
		e.Pos.Y += e.Vel.Y
	}
}
