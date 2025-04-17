package enemyai

import (
	"jjb/textures"
	"jjb/utils"
	"math"
)

func hornetUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	circlingEnemyMovement(enemy, player_pos, 13)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y

		if math.Mod(utils.Game_Time, 100) == 0 {
			angle := -utils.GetAngle(enemy.Pos, player_pos)

			enemy.NewProjectile(enemy.Pos, utils.Vec2{X: math.Sin(angle), Y: -math.Cos(angle)}, textures.NewTexture("./art/enemies/hornet_stinger.png", ""), 5, 20)
		}
	}
}
