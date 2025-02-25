package main

import (
	"fmt"
	"math/rand"
)

func (p *Player) bobertoDamageBuff() {
	p.damage_multiplier *= 1.3
	p.img = *newAnimatedTexture("./art/players/strong_boberto.png")

	start_time := game_time

	for start_time+30 >= game_time {
		fmt.Println("")
	}

	p.damage_multiplier /= 1.3
	p.img = *newAnimatedTexture("./art/players/boberto.png")
}

func (p *Player) realBobertoDamageBuff() {
	go p.bobertoDamageBuff()
}

func (p *Player) bobertoFireball() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x + (p.vel.x * 2) - mouse_x - camera.offset.x + 640 + (float64(p.img.getTexture().Bounds().Dx())), p.pos.y + (p.vel.y * 2) - mouse_y - camera.offset.y + 320 + (float64(p.img.getTexture().Bounds().Dy()))}, 5, 8, 1, -1, "./art/projectiles/boberto/fireball.png")
}

func (p *Player) bobertoFirePiller() {
	p.newProjectile(Vec2{p.pos.x - 64, p.pos.y - 512 + 64}, Vec2{0, 0}, 20, 0, 100, 20, "./art/projectiles/boberto/fire_pillar.png")
}

func (p *Player) bobertoDomain(l *Level) {
	affected := []DomainedEnemy{}
	player_start_pos := p.pos

	for enemy_index := 0; enemy_index < len(l.enemies); enemy_index++ {
		e := &l.enemies[enemy_index]
		affected = append(affected, DomainedEnemy{e, true, e.pos})
		if collide(Vec2{p.pos.x - 1024, p.pos.y - 1024}, Vec2{2048, 2048}, e.pos, Vec2{float64(e.tex.getTexture().Bounds().Dx()), float64(e.tex.getTexture().Bounds().Dy())}) {
			e.pos.x = 1800 + (rand.Float64() * 1000)
			e.pos.y = -1800 - (rand.Float64() * 300)
		}
	}
	p.pos.x = 2000
	p.pos.y = -1600

	start_time := game_time

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		de := affected[enemy_index]
		de.enemy.damage /= 2
	}

	for start_time+150 > game_time {
		for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
			de := affected[enemy_index]
			if de.enemy.health < 0 {
				de.alive = false
			}
		}
	}

	p.pos = player_start_pos

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		de := affected[enemy_index]
		if de.alive {
			de.enemy.pos = de.start_pos
		}
		de.enemy.damage *= 2
	}
}

var boberto_attacks = []Attack{
	{player.realBobertoDamageBuff, 0, 75},
	{player.bobertoFireball, 0, 5},
	{player.bobertoFirePiller, 0, 20},
}
