package main

func (e *Enemy) flieHeadUpdate(p *Player, l *Level) {
	e.vel.x += -0.015 * (e.pos.x - p.pos.x) * (e.pos.y / 100) / 20

	if e.pos.y > p.pos.y-128 {
		e.vel.y -= 1
	}

	if e.vel.x > 10 {
		e.vel.x = 10
	} else if e.vel.x < -10 {
		e.vel.x = -10
	}

	e.vel.y += 0.5

	for le := 0; le < len(l.enemies); le++ {
		oe := &l.enemies[le]
		if e != oe {
			if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{32, 32}, Vec2{float64(oe.pos.x), float64(oe.pos.y)}, Vec2{float64(oe.img.Bounds().Dx()), float64(oe.img.Bounds().Dy())}) {
				e.vel.y = -e.vel.y / 1.2
				oe.vel.y = e.vel.y / 1.2
			}
			if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 32}, Vec2{float64(oe.pos.x), float64(oe.pos.y)}, Vec2{float64(oe.img.Bounds().Dx()), float64(oe.img.Bounds().Dy())}) {
				e.vel.x = -e.vel.x / 1.2
				oe.vel.x = e.vel.x / 1.2
			}
		}
	}

	for ti := 0; ti < len(l.tiles); ti++ {
		t := &l.tiles[ti]
		if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{32, 32}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
			e.vel.x = -e.vel.x / 1.5
		}
		if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{32, 128}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
			e.vel.y = 0
		}
		if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{32, 32}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
			e.vel.y = 0
		}
	}

	e.pos.x += e.vel.x
	e.pos.y += e.vel.y
}
