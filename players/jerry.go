package players

import (
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
)

func (player *Player) jerrySlide() {
	player.Vel.X *= 4
	player.I_Frames = 40
	player.jerrySpike()
}

func (player *Player) jerryMaxBust() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 128 - 16}, utils.Vec2{}, 5, 0, 100, 10, textures.NewTexture("./art/projectiles/jerry/max_bust.png"))
}

func (player *Player) jerryWhiteStuff() {
	player.NewEntity(player.Pos, utils.Vec2{X: 0, Y: 1}, 2, 50, textures.NewTexture("./art/entities/jerry/jerry_white_stuff.png"), jerryWhiteStuffAi)
}

func jerryWhiteStuffAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Vel.Y += 0.1

	entity.Lifespan -= 0.1

	for tile_index := 0; tile_index < len(level_hitbox); tile_index++ {
		tile := level_hitbox[tile_index]
		if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y + entity.Vel.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: tile.X, Y: tile.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Vel.Y = 0
		}
		if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: tile.X, Y: tile.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Vel.X = 0
		}
	}

	if entity.Cooldown < entity.Max_Cooldown {
		for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
			enemy := enemyai.Enemies_In_World[enemy_index]
			if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y - 8}, utils.Vec2{X: 64, Y: 16}, utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
				enemy.Health -= 1
				enemy.Vel.X *= -3
				entity.Cooldown = entity.Max_Cooldown
			}
		}
	} else {
		entity.Cooldown -= 0.1
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y + entity.Vel.Y + 2}, utils.Vec2{X: 64, Y: 8}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		entity.Vel.Y = 0
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: 64, Y: 8}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		entity.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: 64, Y: 8}, utils.Vec2{X: 2000 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		entity.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: 64, Y: 8}, utils.Vec2{X: 2000 + 2048 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		entity.Vel.X = 0
	}
	entity.Pos.Y += entity.Vel.Y
}

func (player *Player) jerrySpike() {
	if player.Dir {
		player.NewEntity(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 64}, utils.Vec2{X: -6, Y: 0}, 2, 50, textures.NewAnimatedTexture("./art/entities/jerry/jerry_spike.png"), JerrySpikeAi)
	} else {
		player.NewEntity(utils.Vec2{X: player.Pos.X - 64, Y: player.Pos.Y - 64}, utils.Vec2{X: 6, Y: 0}, 2, 50, textures.NewAnimatedTexture("./art/entities/jerry/jerry_spike.png"), JerrySpikeAi)

	}
}
func JerrySpikeAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Img.Update()

	if entity.Vel.X >= 0 {
		entity.Dir = false
	} else {
		entity.Dir = true
	}
	entity.Vel.Y += 0.1

	for tile_index := 0; tile_index < len(level_hitbox); tile_index++ {
		tile := level_hitbox[tile_index]
		if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y + entity.Vel.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: tile.X, Y: tile.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Vel.Y = 0
		}
	}

	if entity.Cooldown < 0 {
		for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
			enemy := enemyai.Enemies_In_World[enemy_index]
			if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y - 8}, utils.Vec2{X: 128, Y: 128}, utils.Vec2{X: enemy.Pos.X, Y: enemy.Pos.Y}, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
				enemy.Health -= 2
				entity.Cooldown = entity.Max_Cooldown
			}
		}
	} else {
		entity.Cooldown -= 0.1
	}

	if entity.Img.GetTexture().Bounds().Dx() == 16 {
		entity.Lifespan = -10000
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y + entity.Vel.Y + 2}, utils.Vec2{X: 64, Y: 8}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		entity.Vel.Y = 0
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: 64, Y: 8}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		entity.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: 64, Y: 8}, utils.Vec2{X: 2000 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		entity.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: 64, Y: 8}, utils.Vec2{X: 2000 + 2048 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		entity.Vel.X = 0
	}

	entity.Pos.X += entity.Vel.X
	entity.Pos.Y += entity.Vel.Y
}

var jerry_attacks = []Attack{
	{Player_Ref.jerrySlide, 0, 5},
	{Player_Ref.jerryWhiteStuff, 0, 5},
	{Player_Ref.jerryMaxBust, 0, 1},
}
