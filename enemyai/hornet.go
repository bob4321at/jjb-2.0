package enemyai

import (
	"jjb/textures"
	"jjb/utils"
	"math"
)

func hornetUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	this_enemy_index := 0

	shrimps := 0

	if player_pos.X > enemy.Pos.X {
		enemy.Dir = false
	} else {
		enemy.Dir = true
	}

	for enemy_index := 0; enemy_index < len(Enemies_In_World); enemy_index++ {
		if enemy.Id == 13 {
			shrimps += 1
		}

		if enemy == Enemies_In_World[enemy_index] {
			this_enemy_index = shrimps
			enemy_index = len(Enemies_In_World) + 1
		}
	}

	target_pos := utils.Vec2{X: player_pos.X + (math.Sin(utils.Deg2Rad(utils.Game_Time+float64(this_enemy_index*90)/3)/3) * 300), Y: player_pos.Y + (math.Cos(utils.Deg2Rad(utils.Game_Time+float64(this_enemy_index*90)/3)/3) * 300)}

	if enemy.Pos.X > target_pos.X {
		enemy.Vel.X = -3
	} else {
		enemy.Vel.X = 3
	}
	if enemy.Pos.Y > target_pos.Y {
		enemy.Vel.Y = -3
	} else {
		enemy.Vel.Y = 3
	}

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y

		if math.Mod(utils.Game_Time, 100) == 0 {
			angle := -utils.GetAngle(enemy.Pos, player_pos)

			enemy.NewProjectile(enemy.Pos, utils.Vec2{X: math.Sin(angle), Y: -math.Cos(angle)}, textures.NewTexture("./art/enemies/hornet_stinger.png", ""), 5, 20)
		}
	}
}
