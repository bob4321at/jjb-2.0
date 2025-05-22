package players

import (
	"fmt"
	"jjb/camera"
	"jjb/enemyai"
	"jjb/utils"
	"math/rand"

	"github.com/bob4321at/textures"
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
	if player.Activate_Domain {
		player.Domained_Enemies = []DomainedEnemy{}
		player.Player_Return_Pos = player.Pos

		for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
			e := enemyai.Enemies_In_World[enemy_index]
			player.Domained_Enemies = append(player.Domained_Enemies, DomainedEnemy{e, true, e.Pos})
			if utils.Collide(utils.Vec2{X: player.Pos.X - 1024, Y: player.Pos.Y - 1024}, utils.Vec2{X: 2048, Y: 2048}, e.Pos, utils.Vec2{X: float64(e.Tex.GetTexture().Bounds().Dx()), Y: float64(e.Tex.GetTexture().Bounds().Dy())}) {
				e.Return_To_Pos = e.Pos
				e.Pos.X = 1800 + (rand.Float64() * 1000)
				e.Pos.Y = -1700 - (rand.Float64() * 300)
			}
		}
		player.Pos.X = 2000
		player.Pos.Y = -1600

		player.Domain_Start_Time = utils.Game_Time

		for enemy_index := 0; enemy_index < len(player.Domained_Enemies); enemy_index++ {
			e := player.Domained_Enemies[enemy_index].enemy
			e.Can_Move = true
		}
		player.Activate_Domain = false
		player.Domain_Active = true

		player.Damage_Multiplier *= 2
	}

	if player.Domain_Active && player.Domain_Start_Time+1499 < utils.Game_Time {
		player.Pos = player.Player_Return_Pos
		player.Damage_Multiplier /= 2
		for _, e := range enemies {
			empty_vec2 := utils.Vec2{X: 0, Y: 0}
			if e.Return_To_Pos != empty_vec2 {
				e.Pos = e.Return_To_Pos
				e.Return_To_Pos = utils.Vec2{X: 0, Y: 0}
			}
		}
		player.Domain_Active = false
	}
}

var boberto_attacks = []Attack{
	{Player_Ref.realBobertoDamageBuff, 0, 75},
	{Player_Ref.bobertoFireball, 0, 5},
	{Player_Ref.bobertoFirePiller, 0, 20},
}
