package enemyai

import (
	"jjb/utils"
)

func TreeGuyUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	walkingEnemyMovement(enemy, player_pos, level_hitbox, 3)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
	}
}
