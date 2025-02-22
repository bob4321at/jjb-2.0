package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	id         int
	health     int
	max_health int
	damage     int
	alive      bool
	pos        Vec2
	vel        Vec2
	tex        RenderableTexture
	dir        bool
}

func newEnemy(id int, health int, damage int, pos Vec2, img RenderableTexture) (e Enemy) {
	e.id = id
	e.pos = pos
	e.vel = Vec2{0, 0}
	e.health = health
	e.max_health = health
	e.alive = true

	e.tex = img

	e.damage = damage

	return e
}

func (e *Enemy) Draw(s *ebiten.Image, cam *Camera) {
	if e.pos.x-cam.offset.x+640+float64(e.tex.getTexture().Bounds().Dx()) > 0 && e.pos.x-cam.offset.x+640-float64(e.tex.getTexture().Bounds().Dx()) < 1280 {
		if e.pos.y-cam.offset.y+360+float64(e.tex.getTexture().Bounds().Dy()) > 0 && e.pos.y-cam.offset.y+360-float64(e.tex.getTexture().Bounds().Dy()) < 720 {
			op := ebiten.DrawImageOptions{}

			op.GeoM.Translate(e.pos.x-cam.offset.x+640, e.pos.y-cam.offset.y+360)

			e.tex.draw(s, &op)
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

var enemy_table = map[int]Enemy{
	1: newEnemy(1, 10, 1, Vec2{}, newAnimatedTexture("./art/enemies/fliehead.png")),
	2: newEnemy(2, 20, 2, Vec2{}, newTexture("./art/enemies/crooked.png")),
	3: newEnemy(3, 10, 2, Vec2{}, newTexture("./art/enemies/shrimp.png")),
	4: newEnemy(4, 100, 5, Vec2{}, newAnimatedTexture("./art/enemies/bosshead.png")),
}
