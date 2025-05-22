package enemyai

import (
	"jjb/utils"
	"math"

	"github.com/bob4321at/textures"
)

func simpleGradeCurseUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	walkingEnemyMovement(enemy, player_pos, level_hitbox, 4)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
		check := int(math.Mod(utils.Game_Time, 100))
		if check == 0 {
			if enemy.Dir {
				enemy.NewProjectile(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y - 32}, utils.Vec2{X: -7, Y: 0}, textures.NewTexture("./art/enemies/simple_grade_curse_attack_left.png", ""), 5, 10)
			} else {
				enemy.NewProjectile(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y - 32}, utils.Vec2{X: 7, Y: 0}, textures.NewTexture("./art/enemies/simple_grade_curse_attack_right.png", ""), 5, 10)
			}
		}
	}
}
