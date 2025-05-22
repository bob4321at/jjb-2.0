package enemyai

import (
	"jjb/utils"
	"math"

	"github.com/bob4321at/textures"
)

func cloudHeadUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	enemy.Vel.X += -0.015 * (enemy.Pos.X - player_pos.X) * (math.Abs(enemy.Pos.Y / 100)) / 20

	flyingEnemyMovement(enemy, player_pos, level_hitbox, 10, 128)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
		check := int(math.Mod(utils.Game_Time, 100))
		if check == 0 {
			enemy.NewProjectile(enemy.Pos, utils.Vec2{X: 0, Y: 1}, textures.NewTexture("./art/enemies/cloudhead_rain.png", ""), 3, 100)
		}
	}
}
