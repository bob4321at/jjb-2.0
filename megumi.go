package main

import "math/rand"

func (p *Player) megumiTp() {
	p.pos.x += mouse_x - (1280 / 2)
	p.pos.y += mouse_y - (720 / 2)
}

func (p *Player) megumiBird() {
	if p.pos.x-camera.offset.x+640 < mouse_x {
		p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{-1, 0.5}, 1, 3, 5, 100, newTexture("./art/projectiles/megumi/birdright.png"))
	} else {
		p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{1, 0.5}, 1, 3, 5, 100, newTexture("./art/projectiles/megumi/birdleft.png"))
	}
}

func (p *Player) megumiMahoraga() {
	if !p.dir {
		p.newEntity(Vec2{p.pos.x - 16, p.pos.y - 32}, Vec2{1, 0}, 1, 100, newTexture("./art/entities/megumi/mahoraga.png"), mahoragaUpdate)
	} else {
		p.newEntity(Vec2{p.pos.x - 16, p.pos.y - 32}, Vec2{-1, 0}, 1, 100, newTexture("./art/entities/megumi/mahoraga.png"), mahoragaUpdate)
	}
}

func mahoragaUpdate(e *PlayerEntity) {
	e.vel.y += 0.1

	e.lifespan -= 0.1

	for tile_index := 0; tile_index < len(current_level.tiles); tile_index++ {
		t := &current_level.tiles[tile_index]
		if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{float64(e.img.getTexture().Bounds().Dx()), float64(e.img.getTexture().Bounds().Dy())}, t.pos, Vec2{32, 32}) {
			if e.vel.y >= 0 {
				e.vel.y = -3
			} else {
				e.vel.y = 0
			}
		}
		if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{float64(e.img.getTexture().Bounds().Dx()), float64(e.img.getTexture().Bounds().Dy())}, t.pos, Vec2{32, 32}) {
			e.vel.x = -e.vel.x
		}
	}

	if e.vel.x > 0 {
		e.dir = false
	} else {
		e.dir = true
	}

	if e.cooldown < 0 {
		for enemy_index := 0; enemy_index < len(current_level.enemies); enemy_index++ {
			le := &current_level.enemies[enemy_index]
			if collide(e.pos, Vec2{float64(e.img.getTexture().Bounds().Dx()), float64(e.img.getTexture().Bounds().Dy())}, le.pos, Vec2{float64(le.tex.getTexture().Bounds().Dx()), float64(le.tex.getTexture().Bounds().Dy())}) {
				le.health -= 2
				e.cooldown = 0.5
			}
		}
	} else {
		e.cooldown -= 0.1
	}

	if collide(Vec2{e.pos.x, e.pos.y + e.vel.y + 2}, Vec2{64, 96}, Vec2{2000 - (1280 / 2), -2000 - (720 / 2) + (449 * 2)}, Vec2{2048, (126 * 2)}) {
		e.vel.y = -3
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 96}, Vec2{2000 - (1280 / 2), -2000 - (720 / 2) + (449 * 2)}, Vec2{2048, (126 * 2)}) {
		e.vel.x = -e.vel.x
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 96}, Vec2{2000 - (1280 / 2), -3000 - (720 / 2) + (449 * 2)}, Vec2{1, 1000}) {
		e.vel.x = -e.vel.x
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 96}, Vec2{2000 + 2048 - (1280 / 2), -3000 - (720 / 2) + (449 * 2)}, Vec2{1, 1000}) {
		e.vel.x = -e.vel.x
	}

	e.pos.x += e.vel.x
	e.pos.y += e.vel.y
}

func (p *Player) megumiDomain(l *Level) {
	affected := []DomainedEnemy{}
	player_start_pos := p.pos

	for enemy_index := 0; enemy_index < len(l.enemies); enemy_index++ {
		e := &l.enemies[enemy_index]
		affected = append(affected, DomainedEnemy{e, true, e.pos})
		if collide(Vec2{p.pos.x - 1024, p.pos.y - 1024}, Vec2{2048, 2048}, e.pos, Vec2{float64(e.tex.getTexture().Bounds().Dx()), float64(e.tex.getTexture().Bounds().Dy())}) {
			e.pos.x = 1800 + (rand.Float64() * 1000)
			e.pos.y = -1700 - (rand.Float64() * 300)
		}
	}
	p.pos.x = 2000
	p.pos.y = -1600

	start_time := game_time

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		e := affected[enemy_index].enemy
		e.can_move = true
	}

	for start_time+150 > game_time {
		for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
			de := affected[enemy_index]
			if de.enemy.health < 0 {
				de.alive = false
			}
		}
		for attack_index := 0; attack_index < len(p.attacks); attack_index++ {
			attack := &p.attacks[attack_index]
			if attack.cooldown > attack.max_cooldown/3 {
				attack.cooldown = attack.max_cooldown / 3
			}
		}
	}

	p.pos = player_start_pos

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		de := affected[enemy_index]
		de.enemy.pos = de.start_pos
	}
}

var megumi_attacks = []Attack{
	{player.megumiTp, 0, 4},
	{player.megumiBird, 0, 4},
	{player.megumiMahoraga, 0, 33},
}
