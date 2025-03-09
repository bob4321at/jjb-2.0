package camera

import (
	"jjb/utils"
)

type Camera struct {
	Offset utils.Vec2
}

var Cam = Camera{utils.Vec2{}}
