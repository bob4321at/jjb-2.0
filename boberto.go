package main

import (
	"time"
)

func (p *Player) bobertoDamageBuff() {
	p.damage_multiplier = 1.3
	p.img = *newAnimatedTexture("./art/players/strong_boberto.png")

	time.Sleep(time.Second * 5)

	p.damage_multiplier = 1
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
	for enemy_index := 0; enemy_index < len(l.enemies); enemy_index++ {
		e := &l.enemies[enemy_index]
		e.pos.x = 2000
		e.pos.y = -2000
	}
	p.pos.x = 2000
	p.pos.y = -1600
}

var boberto_attacks = []Attack{
	{player.realBobertoDamageBuff, 0, 100},
	{player.bobertoFireball, 0, 1},
	{player.bobertoFirePiller, 0, 20},
}
