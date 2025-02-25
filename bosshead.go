package main

import (
	"math"
)

func (e *Enemy) bossHeadUpdate(p *Player, l *Level) {
	e.vel.x += -0.015 * (e.pos.x - p.pos.x) * (math.Abs(e.pos.y / 100)) / 20

	if e.pos.y > p.pos.y-256 {
		e.vel.y -= 1
	}

	if e.vel.x > 10 {
		e.vel.x = 10
	} else if e.vel.x < -10 {
		e.vel.x = -10
	}

	check := int(math.Mod(game_time, 10))
	if check == 0 {
		l.Spawn(enemy_table[1])
		game_time += 1
	}

	e.vel.y += 0.5

	for le := 0; le < len(l.enemies); le++ {
		oe := &l.enemies[le]
		if e != oe {
			if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{128, 96}, Vec2{float64(oe.pos.x), float64(oe.pos.y)}, Vec2{float64(oe.tex.getTexture().Bounds().Dx()), float64(oe.tex.getTexture().Bounds().Dy())}) {
				e.vel.y = -e.vel.y / 1.2
				oe.vel.y = e.vel.y / 1.2
			}
			if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{128, 96}, Vec2{float64(oe.pos.x), float64(oe.pos.y)}, Vec2{float64(oe.tex.getTexture().Bounds().Dx()), float64(oe.tex.getTexture().Bounds().Dy())}) {
				e.vel.x = -e.vel.x / 1.2
				oe.vel.x = e.vel.x / 1.2
			}
		}
	}

	for ti := 0; ti < len(l.tiles); ti++ {
		t := &l.tiles[ti]
		if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{128, 96}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
			e.vel.x = -e.vel.x / 1.5
		}
		if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{128, 256}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
			e.vel.y = 0
		}
		if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{128, 96}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
			e.vel.y = 0
		}
	}

	if collide(Vec2{e.pos.x, e.pos.y + e.vel.y + 2}, Vec2{128, 96}, Vec2{2000 - (1280 / 2), -2000 - (720 / 2) + (449 * 2)}, Vec2{2048, (126 * 2)}) {
		e.vel.y = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{128, 96}, Vec2{2000 - (1280 / 2), -2000 - (720 / 2) + (449 * 2)}, Vec2{2048, (126 * 2)}) {
		e.vel.x = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{128, 96}, Vec2{2000 - (1280 / 2), -3000 - (720 / 2) + (449 * 2)}, Vec2{1, 1000}) {
		e.vel.x = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{128, 96}, Vec2{2000 + 2048 - (1280 / 2), -3000 - (720 / 2) + (449 * 2)}, Vec2{1, 1000}) {
		e.vel.x = 0
	}

	if e.can_move {
		e.pos.x += e.vel.x
		e.pos.y += e.vel.y
	}
}
