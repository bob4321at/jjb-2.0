package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
	"math/rand"
)

func (player *Player) gojoRed() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + player.Vel.X - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + player.Vel.Y - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 5, 5, 1, -1, textures.NewTexture("./art/projectiles/gojo/red.png", ""))
}

func (player *Player) gojoBlue() {
	player.NewEntity(player.Pos, utils.Vec2{}, 0, 1, textures.NewTexture("./art/entities/gojo/blue.png", ""), gojoBlueAi)
}

func gojoBlueAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Lifespan -= 0.1

	entity.Pos.X = Player_Ref.Pos.X + ((utils.Mouse_X / 1.5) - (1280 - (1280 / 1.5))) - 72
	entity.Pos.Y = Player_Ref.Pos.Y + ((utils.Mouse_Y / 1.5) - (720 - (720 / 1.5))) - 72
	for ei := 0; ei < len(enemyai.Enemies_In_World); ei++ {
		enemy := enemyai.Enemies_In_World[ei]

		if utils.Collide(utils.Vec2{X: entity.Pos.X - 144, Y: entity.Pos.Y}, utils.Vec2{X: 144 * 3, Y: 144 * 3}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {

			enemy.Can_Move = false
			if enemy.Pos.X+float64(enemy.Tex.GetTexture().Bounds().Dx()/2) >= entity.Pos.X+72 {
				enemy.Vel.X = -0.5
			} else {
				enemy.Vel.X = 0.5
			}
			if enemy.Pos.Y+float64(enemy.Tex.GetTexture().Bounds().Dy()/2) >= entity.Pos.Y+72 {
				enemy.Vel.Y = -0.5
			} else {
				enemy.Vel.Y = 0.5
			}
			enemy.Pos.X += enemy.Vel.X
			enemy.Pos.Y += enemy.Vel.Y
			enemy.Can_Move = true
		}
	}
}

func (player *Player) gojoPurple() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + player.Vel.X - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + player.Vel.Y - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 10, 7, 40, 10, textures.NewTexture("./art/projectiles/gojo/purple.png", ""))
}

func (player *Player) gojoDomain(enemies []*enemyai.Enemy) {
	affected := []DomainedEnemy{}
	player_start_pos := player.Pos

	for enemy_index := 0; enemy_index < len(enemies); enemy_index++ {
		enemy := enemies[enemy_index]
		affected = append(affected, DomainedEnemy{enemy, true, enemy.Pos})
		if utils.Collide(utils.Vec2{X: player.Pos.X - 1024, Y: player.Pos.Y - 1024}, utils.Vec2{X: 2048, Y: 2048}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			enemy.Pos.X = 1800 + (rand.Float64() * 1000)
			enemy.Pos.Y = -1700
		}
	}
	player.Pos.X = 2000
	player.Pos.Y = -1600

	start_time := utils.Game_Time

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		enemy := affected[enemy_index].enemy
		enemy.Can_Move = false
	}

	for start_time+1500 > utils.Game_Time {
		for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
			affected_enemy := affected[enemy_index]
			affected_enemy.enemy.Can_Move = false
			if affected_enemy.enemy.Health < 0 {
				affected_enemy.alive = false
			}
		}
	}

	player.Pos = player_start_pos

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		affected_enemy := affected[enemy_index]
		affected_enemy.enemy.Can_Move = true
		affected_enemy.enemy.Pos = affected_enemy.start_pos
	}
}

var gojo_attacks = []Attack{
	{Player_Ref.gojoRed, 0, 3},
	{Player_Ref.gojoBlue, 0, 0.1},
	{Player_Ref.gojoPurple, 0, 20},
}
