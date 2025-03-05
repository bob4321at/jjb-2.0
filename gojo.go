package main

import (
	"math/rand"
)

func (p *Player) gojoRed() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x + p.vel.x - mouse_x - camera.offset.x + 640 + (float64(p.img.getTexture().Bounds().Dx())), p.pos.y + p.vel.y - mouse_y - camera.offset.y + 320 + (float64(p.img.getTexture().Bounds().Dy()))}, 5, 5, 1, -1, newTexture("./art/projectiles/gojo/red.png"))
}

func (p *Player) gojoBlue() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x + p.vel.x - mouse_x - camera.offset.x + 640 + (float64(p.img.getTexture().Bounds().Dx())), p.pos.y + p.vel.y - mouse_y - camera.offset.y + 320 + (float64(p.img.getTexture().Bounds().Dy()))}, 1, 10, 5, 5, newTexture("./art/projectiles/gojo/blue.png"))
}

func (p *Player) gojoPurple() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x + p.vel.x - mouse_x - camera.offset.x + 640 + (float64(p.img.getTexture().Bounds().Dx())), p.pos.y + p.vel.y - mouse_y - camera.offset.y + 320 + (float64(p.img.getTexture().Bounds().Dy()))}, 1, 7, 40, 10, newTexture("./art/projectiles/gojo/purple.png"))
}

func (p *Player) gojoDomain(l *Level) {
	affected := []DomainedEnemy{}
	player_start_pos := p.pos

	for enemy_index := 0; enemy_index < len(l.enemies); enemy_index++ {
		e := &l.enemies[enemy_index]
		affected = append(affected, DomainedEnemy{e, true, e.pos})
		if collide(Vec2{p.pos.x - 1024, p.pos.y - 1024}, Vec2{2048, 2048}, e.pos, Vec2{float64(e.tex.getTexture().Bounds().Dx()), float64(e.tex.getTexture().Bounds().Dy())}) {
			e.pos.x = 1800 + (rand.Float64() * 1000)
			e.pos.y = -1700
		}
	}
	p.pos.x = 2000
	p.pos.y = -1600

	start_time := game_time

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		e := affected[enemy_index].enemy
		e.can_move = false
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
		de.enemy.can_move = true
		de.enemy.pos = de.start_pos
	}
}

var gojo_attacks = []Attack{
	{player.gojoBlue, 0, 3},
	{player.gojoRed, 0, 5},
	{player.gojoPurple, 0, 20},
}
