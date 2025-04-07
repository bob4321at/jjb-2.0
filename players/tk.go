package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
)

func (player *Player) tkTrafficCone() {
	player.NewEntity(player.Pos, utils.Vec2{X: -(player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx()))), Y: -(player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy())))}, 0, 1, textures.NewTexture("./art/entities/tk/traffic_cone.png", ""), tkTrafficConeAi)
}

func tkTrafficConeAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
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
	entity.Vel.Y += 5
	for level_hitbox_index := 0; level_hitbox_index < len(level_hitbox); level_hitbox_index++ {
		hitbox := level_hitbox[level_hitbox_index]
		if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y + (entity.Vel.Y / 10)}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: hitbox.X, Y: hitbox.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Vel.Y = 0
		}
		if utils.Collide(utils.Vec2{X: entity.Pos.X + (entity.Vel.X / 10), Y: entity.Pos.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: hitbox.X, Y: hitbox.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Vel.X = 0
		}
	}

	for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
		enemy := enemyai.Enemies_In_World[enemy_index]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			enemy.DoDamage(5)
			entity.Lifespan -= 1
		}
	}
	entity.Pos.X += entity.Vel.X / 10
	entity.Pos.Y += entity.Vel.Y / 10
}

func (player *Player) tkCar() {
	if player.Dir {
		player.NewProjectile(utils.Vec2{X: player.Pos.X + 640, Y: player.Pos.Y - 360}, utils.Vec2{X: 1, Y: -0.5}, 1, 3, 50, 100, textures.NewTexture("./art/projectiles/tk/car.png", ""))
	} else {
		player.NewProjectile(utils.Vec2{X: player.Pos.X - 640, Y: player.Pos.Y - 360}, utils.Vec2{X: -1, Y: -0.5}, 1, 3, 50, 100, textures.NewTexture("./art/projectiles/tk/carleft.png", ""))

	}
}

func (player *Player) tkJump() {
	player.Vel.Y = -7
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 32, Y: player.Pos.Y + 16}, utils.Vec2{X: 0, Y: -1}, 10, 5, 10, 10, textures.NewTexture("./art/projectiles/tk/tires.png", ""))
}

var tk_attacks = []Attack{
	{Player_Ref.tkJump, 0, 5},
	{Player_Ref.tkTrafficCone, 0, 5},
	{Player_Ref.tkCar, 0, 20},
}
