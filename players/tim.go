package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/utils"
	"math"

	"github.com/bob4321at/textures"
)

func (player *Player) timSnake() {
	for i := 0; i < 100; i++ {
		snake_balls := player.NewEntity(player.Pos, utils.Vec2{X: 0, Y: 0}, 0, 75, textures.NewTexture("./art/entities/tim/snake_part.png", ""), player.timSnakeBallAi)
		snake_balls.SetID(1)
	}
	snake_head := player.NewEntity(player.Pos, utils.Vec2{X: -(player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx()))), Y: -(player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy())))}, 0, 75, textures.NewTexture("./art/entities/tim/snakehead.png", ""), player.timSnakeAi)
	snake_head.SetID(0)
}

func (player *Player) timSnakeAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	snake_balls := []*PlayerEntity{}

	// entity.Lifespan -= 0.1

	for entity_index := 0; entity_index < len(player.Entities); entity_index++ {
		entity := &player.Entities[entity_index]
		if entity.ID == 1 {
			snake_balls = append(snake_balls, entity)
		}
	}
	if len(snake_balls) == 0 {
		entity.Lifespan = -1
	}

	if len(snake_balls) != 0 {
		point := snake_balls[0]
		if utils.GetDist(point.Pos, utils.Vec2{X: entity.Pos.X + 16, Y: entity.Pos.Y + 16}) > 10 {
			dist := utils.GetDist(point.Pos, utils.Vec2{X: entity.Pos.X + 16, Y: entity.Pos.Y + 16}) / 10
			point.Vel.X = -math.Sin((utils.GetAngle(point.Pos, utils.Vec2{X: entity.Pos.X + 16, Y: entity.Pos.Y + 16}))) * dist
			point.Vel.Y = -math.Cos((utils.GetAngle(point.Pos, utils.Vec2{X: entity.Pos.X + 16, Y: entity.Pos.Y + 16}))) * dist
		}

		if len(snake_balls) > 1 {
			for ball_index := 1; ball_index < len(snake_balls); ball_index++ {
				point := snake_balls[ball_index]
				if utils.GetDist(point.Pos, snake_balls[ball_index-1].Pos) > 10 {
					dist := utils.GetDist(point.Pos, utils.Vec2{X: entity.Pos.X + 16, Y: entity.Pos.Y + 16}) / 10
					point.Vel.X = -math.Sin((utils.GetAngle(point.Pos, snake_balls[ball_index-1].Pos))) * dist
					point.Vel.Y = -math.Cos((utils.GetAngle(point.Pos, snake_balls[ball_index-1].Pos))) * dist
				}
			}
		}
	}

	for i := 0; i < 3; i++ {
		if len(enemyai.Enemies_In_World) > 0 {
			closest_enemy := enemyai.Enemies_In_World[0]

			last_dist := utils.GetDist(entity.Pos, enemyai.Enemies_In_World[0].Pos)

			for ei := 0; ei < len(enemyai.Enemies_In_World); ei++ {
				dist := utils.GetDist(entity.Pos, enemyai.Enemies_In_World[ei].Pos)
				if dist < last_dist && math.Abs(dist) > 64 {
					closest_enemy = enemyai.Enemies_In_World[ei]
				}
			}

			entity.Rotation = -(utils.GetAngle(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y + entity.Vel.Y}, utils.Vec2{X: closest_enemy.Pos.X + closest_enemy.Vel.X, Y: closest_enemy.Pos.Y + closest_enemy.Vel.Y})) - 90

			entity.Vel.X = math.Cos(entity.Rotation)
			entity.Vel.Y = math.Sin(entity.Rotation)
			entity.Pos.X += math.Cos(entity.Rotation)
			entity.Pos.Y += math.Sin(entity.Rotation)
		} else {
			target_pos := utils.Vec2{X: Player_Ref.Pos.X + math.Sin(utils.Game_Time/100)*300, Y: Player_Ref.Pos.Y + math.Cos(utils.Game_Time/100)*300}
			entity.Rotation = -(utils.GetAngle(entity.Pos, target_pos)) - 90
			entity.Vel.X = math.Cos(entity.Rotation)
			entity.Vel.Y = math.Sin(entity.Rotation)
			entity.Pos.X += math.Cos(entity.Rotation)
			entity.Pos.Y += math.Sin(entity.Rotation)
		}
	}
	entity.Rotation = entity.Rotation*(180/3.14159) + 90
}

func (player *Player) timSnakeBallAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Lifespan -= 0.1
	for ei := 0; ei < len(enemyai.Enemies_In_World); ei++ {
		enemy := enemyai.Enemies_In_World[ei]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			enemy.DoDamage(3)
			snake_balls := []*PlayerEntity{}
			for entity_index := 0; entity_index < len(player.Entities); entity_index++ {
				entitty := &player.Entities[entity_index]
				if entitty.ID == 1 {
					snake_balls = append(snake_balls, entitty)
				}
			}
			snake_balls[len(snake_balls)-1].Lifespan = -1
		}
	}

	entity.Pos.X += entity.Vel.X * 3
	entity.Pos.Y += entity.Vel.Y * 3
}

func (player *Player) timSandBall() {
	player.NewEntity(player.Pos, utils.Vec2{X: -(player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx()))), Y: -(player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy())))}, 0, 50, textures.NewTexture("./art/entities/tim/sandball.png", ""), timSandBallAi)
}

func timSandBallAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Lifespan -= 0.05

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
	entity.Vel.Y += 3

	for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
		enemy := enemyai.Enemies_In_World[enemy_index]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			entity.Lifespan -= 1
			enemyai.Enemies_In_World[enemy_index].DoDamage(4)
		}
	}

	entity.Pos.X += entity.Vel.X / 10
	entity.Pos.Y += entity.Vel.Y / 10
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
	{Player_Ref.timSandBall, 0, 15},
	{Player_Ref.timSnake, 0, 70},
}
