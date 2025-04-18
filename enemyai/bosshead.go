package enemyai

import (
	"jjb/utils"
	"math"
)

func bossHeadUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	flyingEnemyMovement(enemy, player_pos, level_hitbox, 10, 128)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
		check := int(math.Mod(utils.Game_Time, 100))
		if check == 0 {
			Enemies_To_Add = append(Enemies_To_Add, Enemy_Table[1])
		}
	}
}
