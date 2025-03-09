package level

import (
	"jjb/camera"
	"jjb/textures"
	"jjb/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	start        utils.Vec2
	repeat_times int
	img          textures.RenderableTexture
}

func newBackground(start_pos utils.Vec2, amount int, img textures.RenderableTexture) (b Background) {
	b.img = img

	b.start = start_pos
	b.repeat_times = amount

	return b
}

func (b *Background) Draw(s *ebiten.Image, cam *camera.Camera) {
	op := ebiten.DrawImageOptions{}

	for repition := 0; repition < b.repeat_times; repition++ {
		op.GeoM.Reset()
		op.GeoM.Translate(b.start.X+float64(repition*b.img.GetTexture().Bounds().Dx())-cam.Offset.X, b.start.Y-cam.Offset.Y)
		b.img.Draw(s, &op)
	}
}
