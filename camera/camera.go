package camera

import (
	"jjb/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	Offset        utils.Vec2
	Manual_Offset utils.Vec2
}

func (c *Camera) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		c.Manual_Offset.X -= 5
	} else if ebiten.IsKeyPressed(ebiten.KeyL) {
		c.Manual_Offset.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		c.Manual_Offset.Y -= 5
	} else if ebiten.IsKeyPressed(ebiten.KeyJ) {
		c.Manual_Offset.Y += 5
	}
}

var Cam = Camera{utils.Vec2{}, utils.Vec2{}}
