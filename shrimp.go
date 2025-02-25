package main

import "math"

func (e *Enemy) shrimpUpdate(p *Player, l *Level) {
	this_enemy_index := 0

	shrimps := 0

	for ei := 0; ei < len(l.enemies); ei++ {
		if e.id == 3 {
			shrimps += 1
		}

		if e == &l.enemies[ei] {
			this_enemy_index = shrimps
			ei = len(l.enemies) + 1
		}
	}

	target_pos := Vec2{p.pos.x + (math.Sin(deg2rad(game_time+float64(this_enemy_index*90))/1000) * 300), p.pos.y + (math.Cos(deg2rad(game_time+float64(this_enemy_index*90))/1000) * 300)}

	if e.pos.x > target_pos.x {
		e.vel.x = -3
	} else {
		e.vel.x = 3
	}
	if e.pos.y > target_pos.y {
		e.vel.y = -3
	} else {
		e.vel.y = 3
	}

	if e.can_move {
		e.pos.x += e.vel.x
		e.pos.y += e.vel.y
	}
}
