package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	id         int
	health     int
	max_health int
	alive      bool
	pos        Vec2
	vel        Vec2
	img        *ebiten.Image
}

func newEnemy(id int, health int, pos Vec2, img_path string) (e Enemy) {
	e.id = id
	e.pos = pos
	e.vel = Vec2{0, 0}
	e.health = health
	e.max_health = health
	e.alive = true

	img, _, err := ebitenutil.NewImageFromFile(img_path)
	if err != nil {
		panic(err)
	}
	e.img = img

	return e
}

func (e *Enemy) Draw(s *ebiten.Image, cam *Camera) {
	if e.pos.x-cam.offset.x+640+float64(e.img.Bounds().Dx()) > 0 && e.pos.x-cam.offset.x+640-float64(e.img.Bounds().Dx()) < 1280 {
		if e.pos.y-cam.offset.y+360+float64(e.img.Bounds().Dy()) > 0 && e.pos.y-cam.offset.y+360-float64(e.img.Bounds().Dy()) < 720 {
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(e.pos.x-cam.offset.x+640, e.pos.y-cam.offset.y+360)

			s.DrawImage(e.img, &op)
		}
	}
}

func (e *Enemy) checkRemove() {
	if e.health <= 0 {
		e.alive = false
		ei := 0
		for i := 0; i < len(current_level.enemies); i++ {
			if &current_level.enemies[i] == e {
				ei = i
				i = len(current_level.enemies) + 1
			}
		}
		current_level.enemies = removeEnemy(ei, current_level.enemies)
	}
}

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
			if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{64, 64}, Vec2{oe.pos.x, oe.pos.y}, Vec2{float64(oe.img.Bounds().Dx()), float64(oe.img.Bounds().Dy())}) {
				e.vel.y = 0
			}
			if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{64, 64}, Vec2{oe.pos.x, oe.pos.y}, Vec2{float64(oe.img.Bounds().Dx()), float64(oe.img.Bounds().Dy())}) {
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

	if e.health <= 0 {
		e.img = ebiten.NewImage(16, 16)
	}

	e.pos.x += e.vel.x
	e.pos.y += e.vel.y
}

var enemy_table = map[int]Enemy{
	1: newEnemy(1, 5, Vec2{}, "./art/enemies/fliehead.png"),
	2: newEnemy(2, 10, Vec2{}, "./art/enemies/crooked.png"),
}
