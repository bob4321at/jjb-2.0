package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/utils"
	"math"
	"math/rand"

	"github.com/bob4321at/textures"
)

func (player *Player) pyroFirePiller() {
	player.Vel.Y = -10
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 256 + 64}, utils.Vec2{X: 0, Y: 0}, 1, 0, 10, 4, textures.NewAnimatedTexture("./art/projectiles/pyro/firetornado.png", ""))
}

func (player *Player) pyroFireDrop() {
	player.NewEntity(player.Pos, utils.Vec2{X: -(player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx()))), Y: -(player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy())))}, 0, 10, textures.NewTexture("./art/entities/pyro/firedrop.png", ""), pyroFireDropAi)
}

func pyroFireDropAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Lifespan -= 0.1

	if entity.Vel.X > 150 {
		entity.Vel.X = 150
	} else if entity.Vel.X < -150 {
		entity.Vel.X = -150
	}
	if entity.Vel.Y > 100 {
		entity.Vel.Y = 100
	} else if entity.Vel.Y < -100 {
		entity.Vel.Y = -100
	}

	if entity.Vel.X > 0 {
		entity.Vel.X -= 1
	} else if entity.Vel.X < 0 {
		entity.Vel.X += 1
	}
	entity.Vel.Y += 3

	for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
		enemy := enemyai.Enemies_In_World[enemy_index]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			entity.Lifespan -= 1
			enemyai.Enemies_In_World[enemy_index].DoDamage(1)
		}
	}

	entity.Pos.X += entity.Vel.X / 10
	entity.Pos.Y += entity.Vel.Y / 10
}

func (player *Player) pyroFireBlasts() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 64}, utils.Vec2{X: -1, Y: 0}, 10, 10, 10, 3, textures.NewAnimatedTexture("./art/projectiles/pyro/fireblast.png", ""))
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 64}, utils.Vec2{X: 1, Y: 0}, 10, 10, 10, 3, textures.NewAnimatedTexture("./art/projectiles/pyro/firedropflipped.png", ""))
}

func (player *Player) pyroDomain(enemies []*enemyai.Enemy) {
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
		if player.Domain_Start_Time+1500 > utils.Game_Time {
			for enemy_index := 0; enemy_index < len(player.Domained_Enemies); enemy_index++ {
				de := player.Domained_Enemies[enemy_index]
				if de.enemy.Health < 0 {
					de.alive = false
				}
			}
		}
		check := int(math.Mod(utils.Game_Time, 1))
		if check == 0 {
			Player_Ref.NewEntity(utils.Vec2{X: 1350 + rand.Float64()*2000, Y: -4000 - (rand.Float64() * 300)}, utils.Vec2{X: 0, Y: 0}, 0, 30, textures.NewTexture("./art/entities/pyro/firedrop.png", ""), pyroFireDropAi)
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

var pyro_attacks = []Attack{
	{Player_Ref.pyroFirePiller, 0, 30},
	{Player_Ref.pyroFireDrop, 0, 5},
	{Player_Ref.pyroFireBlasts, 0, 20},
}
