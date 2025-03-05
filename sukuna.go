package main

import "math"

func sukunaUpdate(e *Enemy, p *Player, l *Level) {
	e.vel.y += 0.1

	if p.pos.x > e.pos.x {
		e.dir = false
		e.vel.x = 4
	} else if p.pos.x < e.pos.x {
		e.dir = true
		e.vel.x = -4
	}

	for ei := 0; ei < len(l.enemies); ei++ {
		oe := &l.enemies[ei]
		if oe != e {
			if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 128}, Vec2{oe.pos.x, oe.pos.y}, Vec2{float64(oe.tex.getTexture().Bounds().Dx()), float64(oe.tex.getTexture().Bounds().Dy())}) {
				if e.pos.x > oe.pos.x {
					e.vel.x = 1
					for ti := 0; ti < len(l.tiles); ti++ {
						t := &l.tiles[ti]
						if t.tile != 0 {
							if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{64, 128}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
								e.vel.y = 0
							}
							if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 128}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
								e.vel.x = 0
							}
						}
					}
				} else {
					e.vel.x = -1
					for ti := 0; ti < len(l.tiles); ti++ {
						t := &l.tiles[ti]
						if t.tile != 0 {
							if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{64, 128}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
								e.vel.y = 0
							}
							if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 128}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
								e.vel.x = 0
							}
						}
					}
				}
			}
		}
	}

	for ti := 0; ti < len(l.tiles); ti++ {
		t := &l.tiles[ti]
		if t.tile != 0 {
			if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{64, 128}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
				e.vel.y = 0
			}
			if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 128}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
				e.vel.x = 0
			}
		}
	}

	if collide(Vec2{e.pos.x, e.pos.y + e.vel.y + 2}, Vec2{64, 128}, Vec2{2000 - (1280 / 2), -2000 - (720 / 2) + (449 * 2)}, Vec2{2048, (126 * 2)}) {
		e.vel.y = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 128}, Vec2{2000 - (1280 / 2), -2000 - (720 / 2) + (449 * 2)}, Vec2{2048, (126 * 2)}) {
		e.vel.x = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 128}, Vec2{2000 - (1280 / 2), -3000 - (720 / 2) + (449 * 2)}, Vec2{1, 1000}) {
		e.vel.x = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 128}, Vec2{2000 + 2048 - (1280 / 2), -3000 - (720 / 2) + (449 * 2)}, Vec2{1, 1000}) {
		e.vel.x = 0
	}

	if e.can_move {
		e.pos.x += e.vel.x
		e.pos.y += e.vel.y
		check := int(math.Mod(game_time, 100))
		if check == 0 {
			if math.Abs(p.pos.x-e.pos.x) > 256 {
				if e.dir {
					e.newProjectile(Vec2{e.pos.x, e.pos.y + 32}, Vec2{-5, 0}, newTexture("./art/enemies/sukuna_fire_attack.png"), 10, 100)
				} else {
					e.newProjectile(Vec2{e.pos.x, e.pos.y + 32}, Vec2{5, 0}, newTexture("./art/enemies/sukuna_fire_attack.png"), 10, 100)
				}
			} else {
				if e.dir {
					e.newProjectile(Vec2{e.pos.x, e.pos.y + 32}, Vec2{-6, 0}, newTexture("./art/enemies/sukuna_attack_cut_left.png"), 5, 20)
				} else {
					e.newProjectile(Vec2{e.pos.x, e.pos.y + 32}, Vec2{6, 0}, newTexture("./art/enemies/sukuna_attack_cut_right.png"), 5, 20)
				}
			}
		}
	}
}
