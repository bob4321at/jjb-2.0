package enemyai

import (
	"jjb/utils"
)

func flieHeadUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	flyingEnemyMovement(enemy, player_pos, level_hitbox, 10, 128)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
	}
}
