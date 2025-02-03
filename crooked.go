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

	// if e.health <= 0 {
	// 	e.tex.getTexture() = ebiten.NewImage(16, 16)
	// }

	e.pos.x += e.vel.x
	e.pos.y += e.vel.y
}
