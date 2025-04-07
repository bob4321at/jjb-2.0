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

func newBackground(start_pos utils.Vec2, amount int, img textures.RenderableTexture) (background Background) {
	background.img = img

	background.start = start_pos
	background.repeat_times = amount

	return background
}

func (background *Background) Draw(screen *ebiten.Image, cam *camera.Camera) {
	op := ebiten.DrawImageOptions{}

	for repition := 0; repition < background.repeat_times; repition++ {
		op.GeoM.Reset()
		op.GeoM.Translate(background.start.X+float64(repition*background.img.GetTexture().Bounds().Dx())-cam.Offset.X-cam.Manual_Offset.X, background.start.Y-cam.Offset.Y-cam.Manual_Offset.Y)
		background.img.Draw(screen, &op)
	}
}
