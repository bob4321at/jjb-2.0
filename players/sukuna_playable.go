package players

import (
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
)

func (player *Player) sukunaDismantle() {
	if !player.Dir {
		player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y - 32}, utils.Vec2{X: -4, Y: 0}, 5, 3, 10, 10, textures.NewTexture("./art/projectiles/sukuna_playable/dismantle_right.png"))
	} else {
		player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y - 32}, utils.Vec2{X: 4, Y: 0}, 5, 3, 10, 10, textures.NewTexture("./art/projectiles/sukuna_playable/dismantle_left.png"))
	}
}

func (player *Player) sukunaFireArrow() {
	if !player.Dir {
		player.NewEntity(player.Pos, utils.Vec2{X: 7, Y: 0}, 3, 10, textures.NewTexture("./art/projectiles/sukuna_playable/fuego_right.png"), sukunaFireArrowAi)
	} else {
		player.NewEntity(player.Pos, utils.Vec2{X: -7, Y: 0}, 3, 10, textures.NewTexture("./art/projectiles/sukuna_playable/fuego_left.png"), sukunaFireArrowAi)
	}
}

func sukunaFireArrowAi(entity *PlayerEntity, level_hitbox []utils.HitBox) {
	entity.Pos.X += entity.Vel.X

	for tile_index := 0; tile_index < len(level_hitbox); tile_index++ {
		tile := level_hitbox[tile_index]
		if utils.Collide(utils.Vec2{X: entity.Pos.X, Y: entity.Pos.Y + entity.Vel.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: tile.X, Y: tile.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Lifespan = -1
			Player_Ref.NewProjectile(utils.Vec2{X: entity.Pos.X - 64, Y: entity.Pos.Y - 512 + 64}, utils.Vec2{X: 0, Y: 0}, 1, 0, 40, 20, textures.NewTexture("./art/projectiles/boberto/fire_pillar.png"))
		}
		if utils.Collide(utils.Vec2{X: entity.Pos.X + entity.Vel.X, Y: entity.Pos.Y}, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: tile.X, Y: tile.Y}, utils.Vec2{X: 32, Y: 32}) {
			entity.Lifespan = -1
			Player_Ref.NewProjectile(utils.Vec2{X: entity.Pos.X - 64, Y: entity.Pos.Y - 512 + 64}, utils.Vec2{X: 0, Y: 0}, 1, 0, 40, 20, textures.NewTexture("./art/projectiles/boberto/fire_pillar.png"))
		}
	}
	for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
		enemy := enemyai.Enemies_In_World[enemy_index]
		if utils.Collide(entity.Pos, utils.Vec2{X: float64(entity.Img.GetTexture().Bounds().Dx()), Y: float64(entity.Img.GetTexture().Bounds().Dy())}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
			entity.Lifespan = -1
			Player_Ref.NewProjectile(utils.Vec2{X: entity.Pos.X - 64, Y: entity.Pos.Y - 512 + 64}, utils.Vec2{X: 0, Y: 0}, 1, 0, 40, 20, textures.NewTexture("./art/projectiles/boberto/fire_pillar.png"))
		}
	}
}

var sukuna_attacks = []Attack{
	{Player_Ref.sukunaFireArrow, 0, 3},
	{Player_Ref.sukunaDismantle, 0, 5},
	{Player_Ref.gojoPurple, 0, 20},
}
