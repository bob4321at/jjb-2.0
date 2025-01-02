package main

import (
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var temp_tile_img, _, _ = ebitenutil.NewImageFromFile("./art/temp_tile.png")

type Tile struct {
	pos  Vec2
	tile int
}

type Level struct {
	tile_map     [][]uint8
	tiles        []Tile
	player_spawn Vec2
	spawn_points []Vec2
	enemies      []Enemy
	background   Background
	generated    bool
}

var temp_tileset = map[int]*ebiten.Image{
	0:  ebiten.NewImage(32, 32),
	1:  ebiten.NewImage(32, 32),
	2:  ebiten.NewImage(32, 32),
	3:  ebiten.NewImage(32, 32),
	4:  ebiten.NewImage(32, 32),
	5:  ebiten.NewImage(32, 32),
	6:  ebiten.NewImage(32, 32),
	7:  ebiten.NewImage(32, 32),
	8:  ebiten.NewImage(32, 32),
	9:  ebiten.NewImage(32, 32),
	10: ebiten.NewImage(32, 32),
	11: ebiten.NewImage(32, 32),
	12: ebiten.NewImage(32, 32),
	13: ebiten.NewImage(32, 32),
	14: ebiten.NewImage(32, 32),
	15: ebiten.NewImage(32, 32),
	16: ebiten.NewImage(32, 32),
}

var current_level *Level

func init() {
	temp_tileset_img, _, err := ebitenutil.NewImageFromFile("./art/temp_tileset.png")
	if err != nil {
		panic(err)
	}
	temp_tileset[1] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(0, 0, 32, 32)))
	temp_tileset[2] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(32, 0, 64, 32)))
	temp_tileset[3] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(64, 0, 96, 32)))
	temp_tileset[4] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(0, 32, 32, 64)))
	temp_tileset[5] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(32, 32, 64, 64)))
	temp_tileset[6] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(64, 32, 96, 64)))
	temp_tileset[7] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(0, 64, 32, 96)))
	temp_tileset[8] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(32, 64, 64, 96)))
	temp_tileset[9] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(64, 64, 96, 96)))
	temp_tileset[10] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(96, 0, 128, 32)))
	temp_tileset[11] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(96, 32, 128, 64)))
	temp_tileset[12] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(96, 64, 128, 96)))
	temp_tileset[13] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(0, 96, 32, 128)))
	temp_tileset[14] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(32, 96, 64, 128)))
	temp_tileset[15] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(64, 96, 96, 128)))
	temp_tileset[16] = ebiten.NewImageFromImage(temp_tileset_img.SubImage(image.Rect(96, 96, 128, 128)))
}

func (l *Level) Draw(s *ebiten.Image, cam *Camera) {
	op := ebiten.DrawImageOptions{}
	l.background.Draw(s, cam)

	for ti := 0; ti < len(l.tiles); ti++ {
		t := &l.tiles[ti]
		op.GeoM.Reset()
		op.GeoM.Translate(t.pos.x-cam.offset.x+640, t.pos.y-cam.offset.y+360)
		s.DrawImage(temp_tileset[t.tile], &op)
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
}

func (l *Level) Spawn(e Enemy) {
	for p := 0; p < len(l.spawn_points); p++ {
		e.pos = l.spawn_points[p]
		e.pos.x += rand.Float64()
		e.pos.y += rand.Float64()
		l.enemies = append(l.enemies, e)
	}
}

var test_place Level
