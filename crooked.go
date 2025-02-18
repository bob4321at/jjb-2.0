package main

func (e *Enemy) crookedUpdate(p *Player, l *Level) {
	e.vel.y += 0.1

	if p.pos.x > e.pos.x {
		e.vel.x = 2
	} else if p.pos.x < e.pos.x {
		e.vel.x = -2
	}

	for ei := 0; ei < len(l.enemies); ei++ {
		oe := &l.enemies[ei]
		if oe != e {
			if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{64, 64}, Vec2{oe.pos.x, oe.pos.y}, Vec2{float64(oe.tex.getTexture().Bounds().Dx()), float64(oe.tex.getTexture().Bounds().Dy())}) {
				e.vel.y = 0
			}
			if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 64}, Vec2{oe.pos.x, oe.pos.y}, Vec2{float64(oe.tex.getTexture().Bounds().Dx()), float64(oe.tex.getTexture().Bounds().Dy())}) {
				e.vel.x = 0
			}
		}
	}

	for ti := 0; ti < len(l.tiles); ti++ {
		t := &l.tiles[ti]
		if t.tile != 0 {
			if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{64, 64}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
				e.vel.y = 0
			}
			if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 64}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
				e.vel.x = 0
			}
		}
	}

	if collide(Vec2{e.pos.x, e.pos.y + e.vel.y + 2}, Vec2{32, 64}, Vec2{2000 - (1280 / 2), -2000 - (720 / 2) + (449 * 2)}, Vec2{2048, (126 * 2)}) {
		e.vel.y = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 32}, Vec2{2000 - (1280 / 2), -2000 - (720 / 2) + (449 * 2)}, Vec2{2048, (126 * 2)}) {
		e.vel.x = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 64}, Vec2{2000 - (1280 / 2), -3000 - (720 / 2) + (449 * 2)}, Vec2{1, 1000}) {
		e.vel.x = 0
	}

	if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 64}, Vec2{2000 + 2048 - (1280 / 2), -3000 - (720 / 2) + (449 * 2)}, Vec2{1, 1000}) {
		e.vel.x = 0
	}

	e.pos.x += e.vel.x
	e.pos.y += e.vel.y
}
