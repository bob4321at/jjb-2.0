package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyProjectile struct {
	pos      Vec2
	vel      Vec2
	img      RenderableTexture
	damage   int
	lifetime float64
}

func (e *Enemy) newProjectile(pos, vel Vec2, img RenderableTexture, damage int, lifetime float64) {
	p := EnemyProjectile{}

	p.pos = pos
	p.vel = vel

	p.img = img

	p.damage = damage
	p.lifetime = lifetime

	e.projectiles = append(e.projectiles, p)
}

type Enemy struct {
	id          int
	health      int
	max_health  int
	damage      int
	projectiles []EnemyProjectile
	alive       bool
	pos         Vec2
	vel         Vec2
	can_move    bool
	tex         RenderableTexture
	dir         bool
	update      func(e *Enemy, p *Player, l *Level)
}

func newEnemy(id int, health int, damage int, pos Vec2, img RenderableTexture, update func(e *Enemy, p *Player, l *Level)) (e Enemy) {
	e.id = id
	e.pos = pos
	e.vel = Vec2{0, 0}
	e.can_move = true
	e.health = health
	e.max_health = health
	e.alive = true

	e.tex = img

	e.damage = damage

	e.update = update

	return e
}

func (e *Enemy) Draw(s *ebiten.Image, cam *Camera) {
	if e.pos.x-cam.offset.x+640+float64(e.tex.getTexture().Bounds().Dx()) > 0 && e.pos.x-cam.offset.x+640-float64(e.tex.getTexture().Bounds().Dx()) < 1280 {
		if e.pos.y-cam.offset.y+360+float64(e.tex.getTexture().Bounds().Dy()) > 0 && e.pos.y-cam.offset.y+360-float64(e.tex.getTexture().Bounds().Dy()) < 720 {
			op := ebiten.DrawImageOptions{}

			if !e.dir {
				op.GeoM.Translate(e.pos.x-cam.offset.x+640, e.pos.y-cam.offset.y+360)
			} else {
				op.GeoM.Scale(-1, 1)
				op.GeoM.Translate(e.pos.x-cam.offset.x+640+float64(e.tex.getTexture().Bounds().Dx()), e.pos.y-cam.offset.y+360)
			}

			e.tex.draw(s, &op)
		}
	}

	for projectile_index := 0; projectile_index < len(e.projectiles); projectile_index++ {
		p := &e.projectiles[projectile_index]

		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.pos.x-cam.offset.x+640, p.pos.y-cam.offset.y+360)

		s.DrawImage(p.img.getTexture(), &op)
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

func (e *Enemy) updateProjectiles(pl *Player) {
	for projectile_index := 0; projectile_index < len(e.projectiles); projectile_index++ {
		p := &e.projectiles[projectile_index]

		p.pos.x += p.vel.x
		p.pos.y += p.vel.y

		if collide(pl.pos, Vec2{32, 64}, p.pos, Vec2{float64(p.img.getTexture().Bounds().Dx()), float64(p.img.getTexture().Bounds().Dy())}) {
			pl.health -= p.damage
			e.projectiles = removeEnemyProjectile(projectile_index, e.projectiles)
			projectile_index -= 1
		}

		p.lifetime -= 0.1
		if p.lifetime < 0 {
			e.projectiles = removeEnemyProjectile(projectile_index, e.projectiles)
			projectile_index -= 1
		}
	}
}

var func_giver_enemy = Enemy{}

var enemy_table = map[int]Enemy{}

func init() {
	enemy_table = map[int]Enemy{
		1: newEnemy(1, 10, 1, Vec2{}, newAnimatedTexture("./art/enemies/fliehead.png"), flieHeadUpdate),
		2: newEnemy(2, 20, 2, Vec2{}, newTexture("./art/enemies/crooked.png"), crookedUpdate),
		3: newEnemy(3, 10, 2, Vec2{}, newTexture("./art/enemies/shrimp.png"), shrimpUpdate),
		4: newEnemy(4, 100, 5, Vec2{}, newAnimatedTexture("./art/enemies/bosshead.png"), bossHeadUpdate),
		5: newEnemy(5, 10, 5, Vec2{}, newTexture("./art/enemies/cloudhead.png"), cloudHeadUpdate),
		6: newEnemy(6, 20, 3, Vec2{}, newTexture("./art/enemies/balloon.png"), balloonUpdate),
		7: newEnemy(7, 10, 3, Vec2{}, newAnimatedTexture("./art/enemies/bunny.png"), bunnyUpdate),
		8: newEnemy(8, 20, 3, Vec2{}, newAnimatedTexture("./art/enemies/fuzzball.png"), fuzzBallUpdate),
	}
}
