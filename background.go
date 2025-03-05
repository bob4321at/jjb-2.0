package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	start        Vec2
	repeat_times int
	img          RenderableTexture
}

func newBackground(start_pos Vec2, amount int, img RenderableTexture) (b Background) {
	b.img = img

	b.start = start_pos
	b.repeat_times = amount

	return b
}

func (b *Background) Draw(s *ebiten.Image, cam *Camera) {
	op := ebiten.DrawImageOptions{}

	for repition := 0; repition < b.repeat_times; repition++ {
		op.GeoM.Reset()
		op.GeoM.Translate(b.start.x+float64(repition*b.img.getTexture().Bounds().Dx())-cam.offset.x, b.start.y-cam.offset.y)
		s.DrawImage(b.img.getTexture(), &op)
	}
}
