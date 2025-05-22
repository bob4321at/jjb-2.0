package players

import (
	"jjb/enemyai"
	"jjb/utils"
	"math"

	"github.com/bob4321at/textures"
)

func (player *Player) birdmanBirdSummon() {
	player.NewEntity(player.Pos, utils.Vec2{X: 0, Y: -5}, 10, 1000, textures.NewAnimatedTexture("./art/entities/birdman/bird.png", ""), birdmanBirdSummonAi)
}

func birdmanBirdSummonAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Lifespan -= 1

	var closest_enemy *enemyai.Enemy
	for enemy_index, enemy := range enemyai.Enemies_In_World {
		if enemy_index == 0 {
			closest_enemy = enemy
		}

		if utils.GetDist(entity.Pos, enemy.Pos) < utils.GetDist(entity.Pos, closest_enemy.Pos) {
			closest_enemy = enemy
		}
	}

	if closest_enemy != nil {
		if entity.Pos.X > closest_enemy.Pos.X {
			entity.Dir = true
		} else {
			entity.Dir = false
		}

		// rot := utils.GetAngle(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y + entity.Vel.Y}, closest_enemy.Pos)
		rot := utils.GetAngle(entity.Pos, closest_enemy.Pos)
		entity.Vel.X += math.Cos(rot + 90)
		entity.Vel.Y += -math.Sin(rot + 90)

		if math.Abs(entity.Vel.X) > 10 {
			entity.Vel.X += -(entity.Vel.X / 10)
		}
		if math.Abs(entity.Vel.Y) > 10 {
			entity.Vel.Y += -(entity.Vel.Y / 10)
		}

		entity.Pos.X += entity.Vel.X
		entity.Pos.Y += entity.Vel.Y
	} else {
		rot := utils.GetAngle(entity.Pos, Player_Ref.Pos)
		entity.Vel.X += math.Cos(rot + 90)
		entity.Vel.Y += -math.Sin(rot + 90)

		if math.Abs(entity.Vel.X) > 10 {
			entity.Vel.X += -(entity.Vel.X / 10)
		}
		if math.Abs(entity.Vel.Y) > 10 {
			entity.Vel.Y += -(entity.Vel.Y / 10)
		}

		entity.Pos.X += entity.Vel.X
		entity.Pos.Y += entity.Vel.Y
	}

	if entity.Cooldown <= 0 {
		for _, enemy := range enemyai.Enemies_In_World {
			if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
				enemy.DoDamage(1)
			}
		}
	} else {
		entity.Cooldown -= 1
	}

	entity.Img.Update()
}

func (player *Player) birdmanFlight() {
	rot := utils.GetAngle(utils.Vec2{X: 0, Y: 0}, utils.Vec2{X: ((utils.Mouse_X / 1.5) - (1280 - (1280 / 1.5))) - 72, Y: ((utils.Mouse_Y / 1.5) - (720 - (720 / 1.5))) - 72})
	player.Vel.X += -math.Sin(rot) / 3
	player.Vel.Y += -math.Cos(rot) / 8
	if player.Vel.X < 0 {
		player.Dir = true
	} else {
		player.Dir = false
	}
}

func (player *Player) birdmanFallingBird() {
	spawn_point := utils.Vec2{}
	var texture textures.RenderableTexture

	if ((utils.Mouse_X)-640)/1.5 > 0 {
		spawn_point = utils.Vec2{X: player.Pos.X + 128*6/2, Y: player.Pos.Y - 72*10/2}
		texture = textures.NewTexture("./art/entities/birdman/falling_birdleft.png", "")
	} else {
		spawn_point = utils.Vec2{X: player.Pos.X - 128*6/2, Y: player.Pos.Y - 72*10/2}
		texture = textures.NewTexture("./art/entities/birdman/falling_birdright.png", "")
	}

	rot := utils.GetAngle(utils.Vec2{X: spawn_point.X, Y: spawn_point.Y}, utils.Vec2{X: player.Pos.X + (((utils.Mouse_X) - 640) / 1.5), Y: player.Pos.Y + (((utils.Mouse_Y) - 360) / 1.5)})

	player.NewProjectile(spawn_point, utils.Vec2{X: math.Sin(rot), Y: math.Cos(rot)}, 10, 5, 2, 100, texture)
}

var birdman_attacks = []Attack{
	{Player_Ref.birdmanBirdSummon, 0, 10},
	{Player_Ref.birdmanFlight, 0, 0},
	{Player_Ref.birdmanFallingBird, 0, 5},
}
