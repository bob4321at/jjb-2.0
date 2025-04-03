package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
	"math"
	"math/rand"
)

func (player *Player) megumiTp() {
	player.Pos.X += utils.Mouse_X - (1280 / 2)
	player.Pos.Y += utils.Mouse_Y - (720 / 2)
}

func (player *Player) megumiSnake() {
	for i := 0; i < 35; i++ {
		snake_balls := player.NewEntity(player.Pos, utils.Vec2{X: 0, Y: 0}, 0, 75, textures.NewTexture("./art/entities/megumi/snake_part.png", ""), player.megumiSnakeBallsAi)
		snake_balls.SetID(1)
	}
	snake_head := player.NewEntity(player.Pos, utils.Vec2{X: -(player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx()))), Y: -(player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy())))}, 0, 75, textures.NewTexture("./art/entities/megumi/snake_head.png", ""), player.megumiSnakeAi)
	snake_head.SetID(0)
}

func (player *Player) megumiSnakeAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	snake_balls := []*PlayerEntity{}

	entity.Lifespan -= 0.1

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
			point.Vel.X = -math.Sin((utils.GetAngle(point.Pos, utils.Vec2{X: entity.Pos.X + 16, Y: entity.Pos.Y + 16})))
			point.Vel.Y = -math.Cos((utils.GetAngle(point.Pos, utils.Vec2{X: entity.Pos.X + 16, Y: entity.Pos.Y + 16})))
		}

		if len(snake_balls) > 1 {
			for ball_index := 1; ball_index < len(snake_balls); ball_index++ {
				point := snake_balls[ball_index]
				if utils.GetDist(point.Pos, snake_balls[ball_index-1].Pos) > 10 {
					point.Vel.X = -math.Sin((utils.GetAngle(point.Pos, snake_balls[ball_index-1].Pos)))
					point.Vel.Y = -math.Cos((utils.GetAngle(point.Pos, snake_balls[ball_index-1].Pos)))
				}
			}
		}
	}

	for i := 0; i < 3; i++ {
		if len(enemyai.Enemies_In_World) > 0 {
			closest_enemy := enemyai.Enemies_In_World[0]

			last_dist := utils.GetDist(entity.Pos, enemyai.Enemies_In_World[0].Pos)

			for ei := 0; ei < len(enemyai.Enemies_In_World); ei++ {
				if utils.GetDist(entity.Pos, enemyai.Enemies_In_World[ei].Pos) < last_dist {
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

func (player *Player) megumiSnakeBallsAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Lifespan -= 0.1
	for ei := 0; ei < len(enemyai.Enemies_In_World); ei++ {
		enemy := enemyai.Enemies_In_World[ei]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			enemy.DoDamage(1)
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

func (player *Player) megumiBird() {
	if player.Pos.X-camera.Cam.Offset.X+640 < utils.Mouse_X {
		player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: -1, Y: 0.5}, 1, 3, 5, 100, textures.NewTexture("./art/projectiles/megumi/birdright.png", ""))
	} else {
		player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: 1, Y: 0.5}, 1, 3, 5, 100, textures.NewTexture("./art/projectiles/megumi/birdleft.png", ""))
	}
}

func (player *Player) megumiMahoraga() {
	if !player.Dir {
		e := player.NewEntity(utils.Vec2{X: player.Pos.X - 16, Y: player.Pos.Y - 32}, utils.Vec2{X: 1, Y: 0}, 1, 100, textures.NewTexture("./art/entities/megumi/mahoraga.png", ""), mahoragaUpdate)
		e.SetID(0)
	} else {
		e := player.NewEntity(utils.Vec2{X: player.Pos.X - 16, Y: player.Pos.Y - 32}, utils.Vec2{X: -1, Y: 0}, 1, 100, textures.NewTexture("./art/entities/megumi/mahoraga.png", ""), mahoragaUpdate)
		e.SetID(0)
	}
}

func mahoragaUpdate(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Vel.Y += 0.1

	entity.Lifespan -= 0.1

	for tile_index := 0; tile_index < len(level_hitbox); tile_index++ {
		tile := level_hitbox[tile_index]
		if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y + entity.Vel.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: tile.X, Y: tile.Y}, utils.Vec2{X: 32, Y: 32}) {
			if entity.Vel.Y >= 0 {
				entity.Vel.Y = -3
			} else {
				entity.Vel.Y = 0
			}
		}
		if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: tile.X, Y: tile.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Vel.X = -entity.Vel.X
		}
	}

	if entity.Vel.X > 0 {
		entity.Dir = false
	} else {
		entity.Dir = true
	}

	if entity.Cooldown < 0 {
		for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
			enemy := enemyai.Enemies_In_World[enemy_index]
			if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
				enemy.DoDamage(2)
				entity.Cooldown = 0.5
			}
		}
	} else {
		entity.Cooldown -= 0.1
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y + entity.Vel.Y + 2}, utils.Vec2{X: 64, Y: 96}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		entity.Vel.Y = -3
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: 64, Y: 96}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		entity.Vel.X = -entity.Vel.X
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: 64, Y: 96}, utils.Vec2{X: 2000 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		entity.Vel.X = -entity.Vel.X
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: 64, Y: 96}, utils.Vec2{X: 2000 + 2048 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		entity.Vel.X = -entity.Vel.X
	}

	entity.Pos.X += entity.Vel.X
	entity.Pos.Y += entity.Vel.Y
}

func (player *Player) megumiDomain(enemies []*enemyai.Enemy) {
	affected := []DomainedEnemy{}
	player_start_pos := player.Pos

	for enemy_index := 0; enemy_index < len(enemies); enemy_index++ {
		enemy := enemies[enemy_index]
		affected = append(affected, DomainedEnemy{enemy, true, enemy.Pos})
		if utils.Collide(utils.Vec2{X: player.Pos.X - 1024, Y: player.Pos.Y - 1024}, utils.Vec2{X: 2048, Y: 2048}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			enemy.Pos.X = 1800 + (rand.Float64() * 1000)
			enemy.Pos.Y = -1700 - (rand.Float64() * 300)
		}
	}
	player.Pos.X = 2000
	player.Pos.Y = -1600

	start_time := utils.Game_Time

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		affected_enemy := affected[enemy_index].enemy
		affected_enemy.Can_Move = true
	}

	for start_time+1500 > utils.Game_Time {
		for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
			affected_enemy := affected[enemy_index]
			if affected_enemy.enemy.Health < 0 {
				affected_enemy.alive = false
			}
		}
		for attack_index := 0; attack_index < len(player.Attacks); attack_index++ {
			attack := &player.Attacks[attack_index]
			if attack.Cooldown > attack.Max_Cooldown/3 {
				attack.Cooldown = attack.Max_Cooldown / 3
			}
		}
	}

	player.Pos = player_start_pos

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		affected_enemy := affected[enemy_index]
		affected_enemy.enemy.Pos = affected_enemy.start_pos
	}
}

var megumi_attacks = []Attack{
	{Player_Ref.megumiSnake, 0, 100},
	{Player_Ref.megumiBird, 0, 4},
	{Player_Ref.megumiMahoraga, 0, 33},
}
