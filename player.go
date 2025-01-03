package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	pos         Vec2
	vel         Vec2
	img         *ebiten.Image
	dir         bool
	attacks     []Attack
	projectiles []Projectile
	entities    []PlayerEntity
}

type Projectile struct {
	pos      Vec2
	vel      Vec2
	damage   int
	speed    float64
	pierce   float64
	lifetime float64
	img      *ebiten.Image
}

type PlayerEntity struct {
	pos    Vec2
	vel    Vec2
	img    *ebiten.Image
	dir    bool
	Update func(e *PlayerEntity)
}

type Attack struct {
	attack       func()
	cooldown     float64
	max_cooldown float64
}

var attack_keys = map[int]ebiten.Key{
	0: ebiten.KeyE,
	2: ebiten.KeyF,
}

func (p *Player) newProjectile(pos, vel Vec2, damage int, speed float64, pierce float64, lifetime float64, img_path string) {
	temp_img, _, err := ebitenutil.NewImageFromFile(img_path)
	if err != nil {
		panic(err)
	}

	projectile := Projectile{pos, vel, damage, speed, pierce, lifetime, temp_img}

	p.projectiles = append(p.projectiles, projectile)
}

func (p *Player) newEntity(pos Vec2, starting_vel Vec2, path string, Update func(e *PlayerEntity)) {
	entity := PlayerEntity{}

	timg, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}
	entity.img = timg

	entity.pos = pos
	entity.vel = starting_vel
	entity.Update = Update

	p.entities = append(p.entities, entity)
}

func newPlayer(pos Vec2, img_path string, attacks []Attack) (p Player) {
	p.pos = pos
	p.vel = Vec2{0, 0}

	img, _, err := ebitenutil.NewImageFromFile(img_path)
	if err != nil {
		panic(err)
	}
	p.img = img

	p.attacks = attacks
	p.dir = false

	return p
}

func (p *Player) Punch() {
	for ie := 0; ie < len(current_level.enemies); ie++ {
		e := &current_level.enemies[ie]
		if collide(Vec2{p.pos.x - 32, p.pos.y}, Vec2{96, 64}, e.pos, Vec2{float64(e.img.Bounds().Dx()), float64(e.img.Bounds().Dy())}) {
			e.health -= 1
		}
	}
}

func (p *Player) Draw(s *ebiten.Image) {
	op := ebiten.DrawImageOptions{}

	if !p.dir {
		op.GeoM.Translate(640, 360)
		s.DrawImage(p.img, &op)
	} else {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(640+32, 360)
		s.DrawImage(p.img, &op)
	}

	for entity_index := 0; entity_index < len(p.entities); entity_index++ {
		e := &p.entities[entity_index]
		op := ebiten.DrawImageOptions{}
		if !e.dir {
			op.GeoM.Translate(e.pos.x-camera.offset.x+640, e.pos.y-camera.offset.y+360)
		} else {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(e.pos.x-camera.offset.x+640+float64(e.img.Bounds().Dx()), e.pos.y-camera.offset.y+360)
		}
		s.DrawImage(e.img, &op)
	}

	for projectile_index := 0; projectile_index < len(p.projectiles); projectile_index++ {
		op.GeoM.Reset()
		op.GeoM.Translate(p.projectiles[projectile_index].pos.x-camera.offset.x+650, p.projectiles[projectile_index].pos.y-camera.offset.y+380)
		s.DrawImage(p.projectiles[projectile_index].img, &op)
	}
}

func (p *Player) Update() {
	p.vel.y += 0.1

	if p.vel.x > 5 {
		p.vel.x -= 0.1
		if p.vel.x > 10 {
			p.vel.x += 0.2
		}
	} else if p.vel.x < -5 {
		p.vel.x += 0.1
		if p.vel.x < -10 {
			p.vel.x += 0.2
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.vel.x -= 0.1
		p.dir = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.vel.x += 0.1
		p.dir = false
	} else {
		if p.vel.x > 0 {
			p.vel.x -= 0.2
			if p.vel.x > -0.6 && p.vel.x < 0.6 {
				p.vel.x = 0
			}
		} else if p.vel.x < 0 {
			p.vel.x += 0.2
			if p.vel.x > -0.6 && p.vel.x < 0.6 {
				p.vel.x = 0
			}
		}
	}

	for b := 0; b < len(p.attacks); b++ {
		p.attacks[b].cooldown -= 0.1
		if ebiten.IsKeyPressed(attack_keys[b]) && p.attacks[b].cooldown < 0 && attack_keys[b] != empty_key {
			p.attacks[b].attack()
			p.attacks[b].cooldown = p.attacks[b].max_cooldown
		} else if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) && p.attacks[1].cooldown < 0 {
			p.attacks[1].attack()
			p.attacks[1].cooldown = p.attacks[1].max_cooldown
		}
	}

	for ti := 0; ti < len(current_level.tiles); ti++ {
		t := &current_level.tiles[ti]
		if t.tile != 0 {
			if collide(Vec2{p.pos.x, p.pos.y + p.vel.y + 2}, Vec2{32, 62}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
				p.vel.y = 0
				if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
					p.vel.y = -5.1
					if collide(Vec2{p.pos.x, p.pos.y + p.vel.y + 2}, Vec2{32, 62}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
						p.vel.y = 0
					}
				}
			}
			if collide(Vec2{p.pos.x + p.vel.x, p.pos.y + 2}, Vec2{32, 62}, Vec2{t.pos.x, t.pos.y}, Vec2{32, 32}) {
				p.vel.x = 0
			}

		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !clicked {
		p.Punch()
		clicked = true
	}

	for projectile_index := 0; projectile_index < len(p.projectiles); projectile_index++ {
		projectile := &p.projectiles[projectile_index]
		projectile_move_dir := math.Atan2(projectile.vel.y, projectile.vel.x)
		projectile.pos.x -= math.Cos(projectile_move_dir) * projectile.speed
		projectile.pos.y -= math.Sin(projectile_move_dir) * projectile.speed

		for ei := 0; ei < len(current_level.enemies); ei++ {
			e := &current_level.enemies[ei]
			if collide(projectile.pos, Vec2{float64(projectile.img.Bounds().Dx()), float64(projectile.img.Bounds().Dy())}, e.pos, Vec2{float64(e.img.Bounds().Dx()), float64(e.img.Bounds().Dy())}) {
				e.health -= projectile.damage
				if projectile.pierce == -1 {
					p.projectiles = removeProjectile(projectile_index, p.projectiles)
					break
				} else {
					projectile.pierce -= 1.1
					if projectile.pierce <= 0 {
						projectile.damage = 0
					}
				}
			}
		}
		if projectile.lifetime != -1 {
			projectile.lifetime -= 0.1
			if projectile.lifetime < 0 {
				p.projectiles = removeProjectile(projectile_index, p.projectiles)
			}
		}
	}

	for entity_index := 0; entity_index < len(p.entities); entity_index++ {
		e := &p.entities[entity_index]
		e.Update(e)
	}

	p.pos.y += p.vel.y
	p.pos.x += p.vel.x
}

var player Player

func init() {
	player = newPlayer(Vec2{0, 0}, "./art/temp_player.png", []Attack{})
}
