package enemyai

import (
	"jjb/shaders"
	"jjb/utils"

	"github.com/bob4321at/textures"
)

func purpleBalloonUpdate(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox) {
	if enemy.Health <= 0 {
		Enemies_To_Add = append(Enemies_To_Add, NewEnemy(5, 10, 5, enemy.Pos, textures.NewTexture("./art/enemies/green_balloon.png", shaders.Enemy_Shader), greenBalloonUpdate))
		Enemies_To_Add = append(Enemies_To_Add, NewEnemy(5, 10, 5, enemy.Pos, textures.NewTexture("./art/enemies/green_balloon.png", shaders.Enemy_Shader), greenBalloonUpdate))
		Enemies_To_Add = append(Enemies_To_Add, NewEnemy(5, 10, 5, enemy.Pos, textures.NewTexture("./art/enemies/green_balloon.png", shaders.Enemy_Shader), greenBalloonUpdate))
	}

	flyingEnemyMovement(enemy, player_pos, level_hitbox, 10, 256)

	if enemy.Can_Move {
		enemy.Pos.X += enemy.Vel.X
		enemy.Pos.Y += enemy.Vel.Y
	}
}
