package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/utils"

	"github.com/bob4321at/textures"
)

func (player *Player) timSandBall() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 5, 10, 1, -1, textures.NewTexture("./art/projectiles/tim/sandball.png", ""))
}

func (player *Player) timCactus() {
	pos, hit := utils.Raycast(utils.Vec2{X: player.Pos.X + (((utils.Mouse_X) - 640) / 1.5), Y: player.Pos.Y + (((utils.Mouse_Y) - 360) / 1.5)}, utils.Vec2{X: 0, Y: 1}, 10000, Emergency_Level_Hitbox)
	if hit {
		player.NewEntity(utils.Vec2{X: pos.X - 48, Y: pos.Y - 128}, utils.Vec2{X: 0, Y: 0}, 1, 40, textures.NewAnimatedTexture("./art/entities/tim/cactus.png", ""), timCactusAi)
	}
}

func timCactusAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	if entity.Lifespan > 38.7 {
		entity.Img.Update()
		entity.Lifespan -= 0.1
	}
	if entity.Cooldown < 0 {
		for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
			enemy := enemyai.Enemies_In_World[enemy_index]
			if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
				enemy.DoDamage(2)
				entity.Cooldown = 0.5
				entity.Lifespan -= 1
			}
		}
	} else {
		entity.Cooldown -= 0.1
	}
}

var tim_attacks = []Attack{
	{Player_Ref.timCactus, 0, 10},
	{Player_Ref.timSandBall, 0, 0},
	{Player_Ref.birdmanFallingBird, 0, 5},
}
