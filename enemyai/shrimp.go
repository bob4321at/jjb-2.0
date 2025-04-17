package enemyai

import (
	"jjb/utils"
)

func shrimpUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	circlingEnemyMovement(enemy, player_pos, 3)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
	}
}
