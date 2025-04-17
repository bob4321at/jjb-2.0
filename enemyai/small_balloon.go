package enemyai

import (
	"jjb/utils"
	"math"
)

func smallBalloonUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	enemy.Vel.X += -0.025 * (enemy.Pos.X - player_pos.X) * (math.Abs(enemy.Pos.Y / 100)) / 20

	flyingEnemyMovement(enemy, player_pos, level_hitbox, 10, 128)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
	}
}
