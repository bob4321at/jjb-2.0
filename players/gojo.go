package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
	"math/rand"
)

func (player *Player) gojoRed() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + player.Vel.X - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + player.Vel.Y - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 5, 5, 1, -1, textures.NewTexture("./art/projectiles/gojo/red.png"))
}

func (player *Player) gojoBlue() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + player.Vel.X - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + player.Vel.Y - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 1, 10, 5, 5, textures.NewTexture("./art/projectiles/gojo/blue.png"))
}

func (player *Player) gojoPurple() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + player.Vel.X - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + player.Vel.Y - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 1, 7, 40, 10, textures.NewTexture("./art/projectiles/gojo/purple.png"))
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
	{Player_Ref.gojoBlue, 0, 3},
	{Player_Ref.gojoRed, 0, 5},
	{Player_Ref.gojoPurple, 0, 20},
}
