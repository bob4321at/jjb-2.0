package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/shaders"
	"jjb/utils"
	"math"
	"math/rand"

	"github.com/bob4321at/textures"
)

func (player *Player) test_playerGlitchProjectile() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 25, 20, 1, -1, textures.NewTexture("./art/projectiles/test_player/glitch.png", ""))
}

func (player *Player) test_playerDeathBar() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 256, Y: player.Pos.Y - 512}, utils.Vec2{X: 0, Y: -1}, 30, 5, 5, 100, textures.NewTexture("./art/projectiles/test_player/death_bar.png", ""))
}

func (player *Player) test_playerTp() {
	player.Pos.X += utils.Mouse_X - (1280 / 2)
	player.Pos.Y += utils.Mouse_Y - (720 / 2)
}

func (player *Player) testGuyDomain(enemies []*enemyai.Enemy) {
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
	}

	if player.Domain_Active {
		for _, e := range enemyai.Enemies_In_World {
			empty_vec2 := utils.Vec2{X: 0, Y: 0}
			if e.Return_To_Pos != empty_vec2 {
				de := e
				if math.Mod(utils.Game_Time, 10) == 0 {
					rand_x := math.Pi * rand.Float64()
					rand_y := math.Pi * rand.Float64()
					player.NewEntity(utils.Vec2{X: de.Pos.X + (math.Cos(rand_x) * 100), Y: de.Pos.Y + (math.Sin(rand_y) * 100)}, utils.Vec2{X: math.Cos(-rand_x) * -10, Y: math.Sin(-rand_y) * 10}, 0, 100, textures.NewTexture("./art/entities/test_guy/balls.png", shaders.Test_Guy_Ball_Shader), testGuyBallAi)
				}
			}
		}
	}

	if player.Domain_Active && player.Domain_Start_Time+1499 < utils.Game_Time {
		player.Pos = player.Player_Return_Pos
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

func testGuyBallAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	if entity.Lifespan == 100 {
		entity.Img.SetUniforms(map[string]any{
			"X": rand.Float64(),
			"Y": rand.Float64(),
			"Z": rand.Float64(),
		})
		entity.Img.Update()
	}
	entity.Lifespan -= 1

	for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
		enemy := enemyai.Enemies_In_World[enemy_index]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			enemy.DoDamage(2)
			entity.Lifespan = -1
		}
	}

	entity.Pos.X += entity.Vel.X
	entity.Pos.Y += entity.Vel.Y
}

var test_player_attacks = []Attack{
	{Player_Ref.test_playerTp, 0, 3},
	{Player_Ref.test_playerGlitchProjectile, 0, 10},
	{Player_Ref.test_playerDeathBar, 0, 40},
}
