package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
	"math"
	"math/rand"
)

func (player *Player) toothbrushGuyToothbrush() {
	player.NewEntity(player.Pos, utils.Vec2{X: player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 0, 100, textures.NewTexture("./art/entities/toothbrush_guy/toothbrush.png", ""), toothbrushGuyToothbrushAi)
}

func toothbrushGuyToothbrushAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Rotation += 10
	entity.Lifespan -= 1

	for ei := 0; ei < len(enemyai.Enemies_In_World); ei++ {
		enemy := enemyai.Enemies_In_World[ei]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			enemy.DoDamage(5)
			entity.Lifespan -= 1
		}
	}

	rot := math.Atan2(entity.Vel.X, entity.Vel.Y)
	entity.Pos.X += math.Sin(rot) * -10
	entity.Pos.Y += math.Cos(rot) * -10
}

func (player *Player) toothbrushGuyPasteBomb() {
	player.NewEntity(player.Pos, utils.Vec2{X: -(player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx()))), Y: -(player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy())))}, 0, 10, textures.NewTexture("./art/entities/toothbrush_guy/paste_bomb.png", ""), toothbrushGuyPasteBombAi)
}

func toothbrushGuyPasteBombAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
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

	for level_hitbox_index := 0; level_hitbox_index < len(level_hitbox); level_hitbox_index++ {
		hitbox := level_hitbox[level_hitbox_index]
		if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: hitbox.X, Y: hitbox.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Lifespan = -1
			Player_Ref.NewProjectile(utils.Vec2{X: entity.Pos.X - 64, Y: entity.Pos.Y - 64}, utils.Vec2{X: 0, Y: 0}, 2, 0, 50, 5, textures.NewTexture("./art/projectiles/toothbrush_guy/paste_splatter.png", ""))
		}
	}

	for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
		enemy := enemyai.Enemies_In_World[enemy_index]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			entity.Lifespan = -1
			enemyai.Enemies_In_World[enemy_index].DoDamage(10)
			Player_Ref.NewProjectile(utils.Vec2{X: entity.Pos.X - 64, Y: entity.Pos.Y - 64}, utils.Vec2{X: 0, Y: 0}, 2, 0, 50, 5, textures.NewTexture("./art/projectiles/toothbrush_guy/paste_splatter.png", ""))
		}
	}

	entity.Pos.X += entity.Vel.X / 10
	entity.Pos.Y += entity.Vel.Y / 10
}

func (player *Player) toothbrushGuyPasteSpout() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 150}, utils.Vec2{X: 0, Y: 0}, 1, 0, 10, 3, textures.NewAnimatedTexture("./art/projectiles/toothbrush_guy/toothpaste_spout.png", ""))
	player.NewEntity(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 128}, utils.Vec2{X: 7 * (rand.Float64()*50 - 25), Y: -(rand.Float64()*50 + 50)}, 0, 10, textures.NewTexture("./art/entities/toothbrush_guy/paste_bomb.png", ""), toothbrushGuyPasteBombAi)
	player.NewEntity(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 128}, utils.Vec2{X: 7 * (rand.Float64()*50 - 25), Y: -(rand.Float64()*50 + 50)}, 0, 10, textures.NewTexture("./art/entities/toothbrush_guy/paste_bomb.png", ""), toothbrushGuyPasteBombAi)
	player.NewEntity(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 128}, utils.Vec2{X: 7 * (rand.Float64()*50 - 25), Y: -(rand.Float64()*50 + 50)}, 0, 10, textures.NewTexture("./art/entities/toothbrush_guy/paste_bomb.png", ""), toothbrushGuyPasteBombAi)
	player.NewEntity(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 128}, utils.Vec2{X: 7 * (rand.Float64()*50 - 25), Y: -(rand.Float64()*50 + 50)}, 0, 10, textures.NewTexture("./art/entities/toothbrush_guy/paste_bomb.png", ""), toothbrushGuyPasteBombAi)
}

func (player *Player) toothbrushDomain(enemies []*enemyai.Enemy) {
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
		check := int(math.Mod(utils.Game_Time, 20))
		if check == 0 {
			ran := rand.Float64()*2 - 1
			if ran > 0 {
				player.NewProjectile(utils.Vec2{X: player.Pos.X - 2560 + (rand.Float64() * 256), Y: player.Pos.Y - ((ran + 1) * 512)}, utils.Vec2{X: -5, Y: 0}, 1, 10, 10, 100, textures.NewTexture("./art/projectiles/toothbrush_guy/big_brush.png", ""))
			} else {
				player.NewProjectile(utils.Vec2{X: player.Pos.X + 2560 + (rand.Float64() * 256), Y: player.Pos.Y - ((ran + 1) * 512)}, utils.Vec2{X: 5, Y: 0}, 1, 10, 10, 100, textures.NewTexture("./art/projectiles/toothbrush_guy/big_brush.png", ""))

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

var toothbrush_guy_attacks = []Attack{
	{Player_Ref.toothbrushGuyPasteBomb, 0, 30},
	{Player_Ref.toothbrushGuyToothbrush, 0, 5},
	{Player_Ref.toothbrushGuyPasteSpout, 0, 40},
}
