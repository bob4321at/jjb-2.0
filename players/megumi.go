package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
	"math/rand"
)

func (player *Player) megumiTp() {
	player.Pos.X += utils.Mouse_X - (1280 / 2)
	player.Pos.Y += utils.Mouse_Y - (720 / 2)
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
		player.NewEntity(utils.Vec2{X: player.Pos.X - 16, Y: player.Pos.Y - 32}, utils.Vec2{X: 1, Y: 0}, 1, 100, textures.NewTexture("./art/entities/megumi/mahoraga.png", ""), mahoragaUpdate)
	} else {
		player.NewEntity(utils.Vec2{X: player.Pos.X - 16, Y: player.Pos.Y - 32}, utils.Vec2{X: -1, Y: 0}, 1, 100, textures.NewTexture("./art/entities/megumi/mahoraga.png", ""), mahoragaUpdate)
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
	{Player_Ref.megumiTp, 0, 4},
	{Player_Ref.megumiBird, 0, 4},
	{Player_Ref.megumiMahoraga, 0, 33},
}
