package main

import "math"

func simpleGradeCurseUpdate(e *Enemy, p *Player, l *Level) {
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
			if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 64}, Vec2{oe.pos.x, oe.pos.y}, Vec2{float64(oe.tex.getTexture().Bounds().Dx()), float64(oe.tex.getTexture().Bounds().Dy())}) {
				if e.pos.x > oe.pos.x {
					e.vel.x = 1
					for ti := 0; ti < len(l.tiles); ti++ {
						t := &l.tiles[ti]
						if t.tile != 0 {
							if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{32, 64}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
								e.vel.y = 0
							}
							if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 64}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
								e.vel.x = 0
							}
						}
					}
				} else {
					e.vel.x = -1
					for ti := 0; ti < len(l.tiles); ti++ {
						t := &l.tiles[ti]
						if t.tile != 0 {
							if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{32, 64}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
								e.vel.y = 0
							}
							if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 64}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
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
			if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{32, 64}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
				e.vel.y = 0
			}
			if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 64}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
				e.vel.x = 0
			}
		}
	}

	if collide(Vec2{e.pos.x, e.pos.y + e.vel.y + 2}, Vec2{32, 64}, Vec2{2000 - (1280 / 2), -2000 - (720 / 2) + (449 * 2)}, Vec2{2048, (126 * 2)}) {
		e.vel.y = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 64}, Vec2{2000 - (1280 / 2), -2000 - (720 / 2) + (449 * 2)}, Vec2{2048, (126 * 2)}) {
		e.vel.x = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 64}, Vec2{2000 - (1280 / 2), -3000 - (720 / 2) + (449 * 2)}, Vec2{1, 1000}) {
		e.vel.x = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 64}, Vec2{2000 + 2048 - (1280 / 2), -3000 - (720 / 2) + (449 * 2)}, Vec2{1, 1000}) {
		e.vel.x = 0
	}

	if e.can_move {
		e.pos.x += e.vel.x
		e.pos.y += e.vel.y
		check := int(math.Mod(game_time, 10))
		if check == 0 {
			if e.dir {
				e.newProjectile(Vec2{e.pos.x, e.pos.y - 32}, Vec2{-7, 0}, newTexture("./art/enemies/simple_grade_curse_attack_left.png"), 5, 10)
			} else {
				e.newProjectile(Vec2{e.pos.x, e.pos.y - 32}, Vec2{7, 0}, newTexture("./art/enemies/simple_grade_curse_attack_right.png"), 5, 10)
			}
			game_time += 1
		}
	}
}
