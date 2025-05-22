package enemyai

import (
	"jjb/utils"
	"math"

	"github.com/bob4321at/textures"
)

func sukunaUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	walkingEnemyMovement(enemy, player_pos, level_hitbox, 4)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
		check := int(math.Mod(utils.Game_Time, 100))
		if check == 0 {
			if math.Abs(player_pos.X-enemy.Pos.X) > 256 {
				if enemy.Dir {
					enemy.NewProjectile(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + 32}, utils.Vec2{X: -5, Y: 0}, textures.NewTexture("./art/enemies/fuego_left.png", ""), 10, 100)
				} else {
					enemy.NewProjectile(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + 32}, utils.Vec2{X: 5, Y: 0}, textures.NewTexture("./art/enemies/fuego_right.png", ""), 10, 100)
				}
			} else {
				if enemy.Dir {
					enemy.NewProjectile(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + 32}, utils.Vec2{X: -6, Y: 0}, textures.NewTexture("./art/enemies/sukuna_attack_cut_left.png", ""), 5, 20)
				} else {
					enemy.NewProjectile(utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y + 32}, utils.Vec2{X: 6, Y: 0}, textures.NewTexture("./art/enemies/sukuna_attack_cut_right.png", ""), 5, 20)
				}
			}
		}
	}
}
