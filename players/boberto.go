package players

import (
	"fmt"
	"jjb/camera"
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
	"math/rand"
)

func (player *Player) bobertoDamageBuff() {
	player.Damage_Multiplier *= 2
	player.Img = *textures.NewAnimatedTexture("./art/players/strong_boberto.png", "")

	start_time := utils.Game_Time

	for start_time+300 >= utils.Game_Time {
		fmt.Println("")
	}

	player.Damage_Multiplier /= 2
	player.Img = *textures.NewAnimatedTexture("./art/players/boberto.png", "")
}

func (player *Player) realBobertoDamageBuff() {
	go player.bobertoDamageBuff()
}

func (player *Player) bobertoFireball() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 5, 8, 1, -1, textures.NewTexture("./art/projectiles/boberto/fireball.png", ""))
}

func (player *Player) bobertoFirePiller() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 512 + 64}, utils.Vec2{X: 0, Y: 0}, 3, 0, 40, 20, textures.NewTexture("./art/projectiles/boberto/fire_pillar.png", ""))
}

func (player *Player) bobertoDomain(enemies []*enemyai.Enemy) {
	affected := []DomainedEnemy{}
	player_start_pos := player.Pos

	for enemy_index := 0; enemy_index < len(enemies); enemy_index++ {
		enemy := enemies[enemy_index]
		affected = append(affected, DomainedEnemy{enemy, true, enemy.Pos})
		if utils.Collide(utils.Vec2{X: player.Pos.X - 1024, Y: player.Pos.Y - 1024}, utils.Vec2{X: 2048, Y: 2048}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			enemy.Pos.X = 1800 + (rand.Float64() * 1000)
			enemy.Pos.Y = -1800 - (rand.Float64() * 300)
		}
	}
	player.Pos.X = 2000
	player.Pos.Y = -1600

	start_time := utils.Game_Time

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		affected_enemy := affected[enemy_index]
		affected_enemy.enemy.Damage /= 2
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
		if affected_enemy.alive {
			affected_enemy.enemy.Pos = affected_enemy.start_pos
		}
		affected_enemy.enemy.Damage *= 2
	}
}

var boberto_attacks = []Attack{
	{Player_Ref.realBobertoDamageBuff, 0, 75},
	{Player_Ref.bobertoFireball, 0, 5},
	{Player_Ref.bobertoFirePiller, 0, 20},
}
