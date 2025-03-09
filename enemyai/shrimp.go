package enemyai

import (
	"jjb/utils"
	"math"
)

func shrimpUpdate(e *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	this_enemy_index := 0

	shrimps := 0

	for enemy_index := 0; enemy_index < len(Enemies_In_World); enemy_index++ {
		if e.Id == 3 {
			shrimps += 1
		}

		if e == Enemies_In_World[enemy_index] {
			this_enemy_index = shrimps
			enemy_index = len(Enemies_In_World) + 1
		}
	}

	target_pos := utils.Vec2{X: player_pos.X + (math.Sin(utils.Deg2Rad(utils.Game_Time+float64(this_enemy_index*90))/1000) * 300), Y: player_pos.Y + (math.Cos(utils.Deg2Rad(utils.Game_Time+float64(this_enemy_index*90))/1000) * 300)}

	if e.Pos.X > target_pos.X {
		e.Vel.X = -3
	} else {
		e.Vel.X = 3
	}
	if e.Pos.Y > target_pos.Y {
		e.Vel.Y = -3
	} else {
		e.Vel.Y = 3
	}

	if e.Can_Move {
		e.Pos.X += e.Vel.X
		e.Pos.Y += e.Vel.Y
	}
}
