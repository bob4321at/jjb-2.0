package main

import (
	"math/rand"

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

func (l *Level) Draw(s *ebiten.Image, cam *Camera) {
	op := ebiten.DrawImageOptions{}
	l.background.Draw(s, cam)

	for ti := 0; ti < len(l.tiles); ti++ {
		t := &l.tiles[ti]
		op.GeoM.Reset()
		op.GeoM.Translate(t.pos.x-cam.offset.x+640, t.pos.y-cam.offset.y+360)
		s.DrawImage(l.tileset[t.tile], &op)
	}

	for e := 0; e < len(l.enemies); e++ {
		l.enemies[e].Draw(s, cam)
	}
}

func (l *Level) Update(p *Player) {
	for e := 0; e < len(l.enemies); e++ {
		if l.enemies[e].id == 1 {
			l.enemies[e].flieHeadUpdate(p, l)
		} else if l.enemies[e].id == 2 {
			l.enemies[e].crookedUpdate(p, l)
		}
		l.enemies[e].checkRemove()
	}

	if !l.spawned {
		for enemy_index := 0; enemy_index < len(l.waves.Waves[l.current_wave]); enemy_index += 1 - 1 {
			if l.spawn_timer < 0 {
				l.spawn_timer = l.origonal_spawn_timer
				l.Spawn(enemy_table[l.waves.Waves[l.current_wave][l.waves.Waves[l.current_wave][enemy_index]]])
				enemy_index += 1
			} else {
				l.spawn_timer -= 0.1
			}
		}
		l.spawned = true
	}

	if len(l.enemies) == 0 {
		l.spawned = false
		if l.current_wave+1 < len(l.waves.Waves) {
			l.current_wave += 1
		} else {
			l.spawned = true
		}
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
