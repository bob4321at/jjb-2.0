package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/utils"

	"github.com/bob4321at/textures"
)

func (player *Player) agent_21Gun() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 10, 5, 1, -1, textures.NewTexture("./art/projectiles/agent_21/gun.png", ""))
}

func (player *Player) agent_21C4() {
	player.NewEntity(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y + 32}, utils.Vec2{X: 0, Y: 1}, 0, 100, textures.NewTexture("./art/entities/agent_21/c4.png", ""), agent_21C4Ai)
}

func agent_21C4Ai(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Vel.Y += 0.1

	for level_hitbox_index := 0; level_hitbox_index < len(level_hitbox); level_hitbox_index++ {
		hitbox := level_hitbox[level_hitbox_index]
		if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y + entity.Vel.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: hitbox.X, Y: hitbox.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Vel.Y = 0
		}
	}

	for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
		enemy := enemyai.Enemies_In_World[enemy_index]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			entity.Lifespan = -1
			Player_Ref.NewProjectile(utils.Vec2{X: entity.Pos.X - 128, Y: entity.Pos.Y - 128}, utils.Vec2{X: 0, Y: 0}, 10, 0, 10, 10, textures.NewTexture("./art/projectiles/greg/explosion.png", ""))
		}
	}

	entity.Pos.Y += entity.Vel.Y
}

func (player *Player) agent_21Molotov() {
	player.NewEntity(player.Pos, utils.Vec2{X: -(player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx()))), Y: -(player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy())))}, 0, 100, textures.NewTexture("./art/entities/agent_21/molotov.png", ""), agent_21MolotovAi)
}

func agent_21MolotovAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	if entity.Vel.X > 100 {
		entity.Vel.X = 100
	} else if entity.Vel.X < -100 {
		entity.Vel.X = -100
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
	for level_hitbox_index := 0; level_hitbox_index < len(level_hitbox); level_hitbox_index++ {
		hitbox := level_hitbox[level_hitbox_index]
		if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y + (entity.Vel.Y / 10)}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: hitbox.X, Y: hitbox.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Lifespan = -1
			Player_Ref.NewEntity(entity.Pos, utils.Vec2{X: 0, Y: 0}, 10, 100, textures.NewTexture("./art/entities/agent_21/fire.png", ""), agent_21Firewall)
		}
		if utils.Collide(utils.Vec2{X: entity.Pos.X + (entity.Vel.X / 10), Y: entity.Pos.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: hitbox.X, Y: hitbox.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Lifespan = -1
			Player_Ref.NewProjectile(utils.Vec2{X: entity.Pos.X - 128, Y: entity.Pos.Y - 128}, utils.Vec2{X: 0, Y: 0}, 10, 0, 10, 10, textures.NewTexture("./art/projectiles/greg/explosion.png", ""))
		}
	}
	entity.Vel.Y += 5

	for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
		enemy := enemyai.Enemies_In_World[enemy_index]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			entity.Lifespan = -1
			Player_Ref.NewProjectile(utils.Vec2{X: entity.Pos.X - 128, Y: entity.Pos.Y - 128}, utils.Vec2{X: 0, Y: 0}, 10, 0, 10, 10, textures.NewTexture("./art/projectiles/greg/explosion.png", ""))
		}
	}
	entity.Pos.X += entity.Vel.X / 10
	entity.Pos.Y += entity.Vel.Y / 10
}

func agent_21Firewall(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Lifespan -= 1
	if entity.Cooldown <= 0 {
		for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
			enemy := enemyai.Enemies_In_World[enemy_index]
			if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
				enemy.DoDamage(1)
				entity.Lifespan -= 1
				entity.Cooldown = entity.Max_Cooldown
			}
		}
	} else {
		entity.Cooldown -= 0.1
	}
}

var agent_21_attacks = []Attack{
	{Player_Ref.agent_21C4, 0, 20},
	{Player_Ref.agent_21Gun, 0, 5},
	{Player_Ref.agent_21Molotov, 0, 30},
}
