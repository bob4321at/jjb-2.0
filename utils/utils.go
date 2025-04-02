package utils

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Vec2 struct {
	X, Y float64
}

type HitBox struct {
	X, Y, W, H float64
}

func NewHitBox(x, y, w, h float64) HitBox {
	return HitBox{X: x, Y: y, W: w, H: h}
}

var Domain_Background, _, _ = ebitenutil.NewImageFromFile("./art/domains/domain_backdrop.png")

var Empty_Key ebiten.Key
var Clicked bool = false

var Mouse_X, Mouse_Y float64
var Game_Time float64 = 0

func Collide(pos1, size1, pos2, size2 Vec2) bool {
	if pos1.X < pos2.X+size2.X && pos1.X+size1.X > pos2.X {
		if pos1.Y < pos2.Y+size2.Y && pos1.Y+size1.Y > pos2.Y {
			return true
		}
	}
	return false
}

func RemoveArrayElement[T any](index_to_remove int, slice *[]T) {
	*slice = append((*slice)[:index_to_remove], (*slice)[index_to_remove+1:]...)
}

func Deg2Rad(num float64) float64 {
	return num * (180 / 3.14159)
}

func GetDist(p1, p2 Vec2) float64 {
	offx := math.Abs(p1.X - p2.X)
	offy := math.Abs(p1.Y - p2.Y)

	return math.Sqrt((offx * offx) + (offy * offy))
}

func GetAngle(p1, p2 Vec2) float64 {
	offx := p1.X - p2.X
	offy := p1.Y - p2.Y

	return math.Atan2(offx, offy)
}
