package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Background struct {
	start        Vec2
	repeat_times int
	img          *ebiten.Image
}

func newBackground(start_pos Vec2, amount int, path string) (b Background) {
	timg, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}
	b.img = timg

	b.start = start_pos
	b.repeat_times = amount

	return b
}

func (b *Background) Draw(s *ebiten.Image, cam *Camera) {
	op := ebiten.DrawImageOptions{}

	for repition := 0; repition < b.repeat_times; repition++ {
		op.GeoM.Reset()
		op.GeoM.Translate(b.start.x+float64(repition*b.img.Bounds().Dx())-cam.offset.x, b.start.y-cam.offset.y)
		s.DrawImage(b.img, &op)
	}
}
