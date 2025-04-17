package enemyai

import (
	"jjb/shaders"
	"jjb/textures"
	"jjb/utils"
)

func greenBalloonUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	if enemy.Health <= 0 {
		Enemies_To_Add = append(Enemies_To_Add, NewEnemy(6, 20, 3, enemy.Pos, textures.NewTexture("./art/enemies/balloon.png", shaders.Enemy_Shader), balloonUpdate))
		Enemies_To_Add = append(Enemies_To_Add, NewEnemy(6, 20, 3, enemy.Pos, textures.NewTexture("./art/enemies/balloon.png", shaders.Enemy_Shader), balloonUpdate))
	}

	flyingEnemyMovement(enemy, player_pos, level_hitbox, 10, 256)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
	}
}
