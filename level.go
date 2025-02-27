package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var temp_tile_img, _, _ = ebitenutil.NewImageFromFile("./art/temp_tile.png")

type Tile struct {
	pos  Vec2
	tile int
}

type Waves struct {
	Waves [][]int
}

type Level struct {
	tile_map             [][]uint8
	tiles                []Tile
	player_spawn         Vec2
	spawn_points         []Vec2
	enemies              []Enemy
	background           Background
	generated            bool
	tileset              map[int]*ebiten.Image
	waves                Waves
	current_wave         int
	spawn_timer          float64
	origonal_spawn_timer float64
	spawned              bool
}

var levels = []Level{}

var current_level *Level
var current_level_index int

func (l *Level) Draw(s *ebiten.Image, cam *Camera) {
	op := ebiten.DrawImageOptions{}
	l.background.Draw(s, cam)

	for ti := 0; ti < len(l.tiles); ti++ {
		t := &l.tiles[ti]
		op.GeoM.Reset()
		op.GeoM.Translate(t.pos.x-cam.offset.x+640, t.pos.y-cam.offset.y+360)
		s.DrawImage(l.tileset[t.tile], &op)
	}

	op.GeoM.Reset()

	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(1000-camera.offset.x, -2500-camera.offset.y)
	s.DrawImage(domain_background, &op)

	op.GeoM.Reset()

	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(2000-camera.offset.x, -2000-camera.offset.y)
	s.DrawImage(player.domain.img.getTexture(), &op)

	for e := 0; e < len(l.enemies); e++ {
		l.enemies[e].Draw(s, cam)
	}
}

func (l *Level) SpawnWave() {
	for enemy_index := 0; enemy_index < len(l.waves.Waves[l.current_wave]); enemy_index += 1 - 1 {
		time.Sleep(20000)
		if l.spawn_timer < 0 {
			l.spawn_timer = l.origonal_spawn_timer
			l.Spawn(enemy_table[l.waves.Waves[l.current_wave][enemy_index]])
			enemy_index += 1
		} else {
			l.spawn_timer -= 0.01
		}
	}

	for l.spawned {
		if len(l.enemies) == 0 {
			l.spawned = false
			l.current_wave += 1
		}
	}
}

func (l *Level) Update(p *Player) {
	p.damageCheck(l)

	for e := 0; e < len(l.enemies); e++ {
		l.enemies[e].update(&l.enemies[e], p, l)

		l.enemies[e].updateProjectiles(&player)

		l.enemies[e].tex.update()

		l.enemies[e].checkRemove()
	}

	if !l.spawned && l.current_wave < len(l.waves.Waves) {
		go l.SpawnWave()
		l.spawned = true
	}
	if l.current_wave >= len(l.waves.Waves) && current_level_index+1 < len(levels) {
		current_level_index += 1
	}
}

func (l *Level) Spawn(e Enemy) {
	for p := 0; p < len(l.spawn_points); p++ {
		e.pos = l.spawn_points[p]
		e.pos.x += rand.Float64()
		e.pos.y += rand.Float64()
		l.enemies = append(l.enemies, e)
	}
}
