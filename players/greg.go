package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
	"math"
	"math/rand"
)

func (player *Player) gregLaunch() {
	player.Vel.Y = -10
	if !player.Dir {
		player.Vel.X = 10
	} else {
		player.Vel.X = -10
	}
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 128, Y: player.Pos.Y - 128}, utils.Vec2{X: 0, Y: 0}, 2, 0, 10, 2, textures.NewTexture("./art/projectiles/greg/launch_explosion.png", ""))
}

func (player *Player) gregThrow() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 5, 10, 1, -1, textures.NewTexture("./art/projectiles/greg/rock.png", ""))
}

func (player *Player) gregNuke() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 128, Y: player.Pos.Y - 128}, utils.Vec2{X: 0, Y: 0}, 1, 0, 5, 10, textures.NewTexture("./art/projectiles/greg/explosion.png", ""))
}

func (player *Player) gregDomain(enemies []*enemyai.Enemy) {
	affected := []DomainedEnemy{}
	player_start_pos := player.Pos

	for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
		e := enemyai.Enemies_In_World[enemy_index]
		affected = append(affected, DomainedEnemy{e, true, e.Pos})
		if utils.Collide(utils.Vec2{X: player.Pos.X - 1024, Y: player.Pos.Y - 1024}, utils.Vec2{X: 2048, Y: 2048}, e.Pos, utils.Vec2{X: float64(e.Tex.GetTexture().Bounds().Dx()), Y: float64(e.Tex.GetTexture().Bounds().Dy())}) {
			e.Pos.X = 1800 + (rand.Float64() * 1000)
			e.Pos.Y = -1700 - (rand.Float64() * 300)
		}
	}
	player.Pos.X = 2000
	player.Pos.Y = -1600

	start_time := utils.Game_Time

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		e := affected[enemy_index].enemy
		e.Can_Move = true
	}

	for start_time+1500 > utils.Game_Time {
		for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
			de := affected[enemy_index]
			if math.Mod(utils.Game_Time, 30) == 0 {
				player.NewProjectile(utils.Vec2{X: de.enemy.Pos.X - 64, Y: de.enemy.Pos.Y - 64}, utils.Vec2{X: 0, Y: 0}, 1, 0, 5, 5, textures.NewTexture("./art/projectiles/greg/domain_explosion.png", ""))
			}
			if de.enemy.Health < 0 {
				de.alive = false
			}
		}
	}

	player.Pos = player_start_pos

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		de := affected[enemy_index]
		de.enemy.Pos = de.start_pos
	}
}

var greg_attacks = []Attack{
	{Player_Ref.gregLaunch, 0, 20},
	{Player_Ref.gregThrow, 0, 5},
	{Player_Ref.gregNuke, 0, 1},
}
