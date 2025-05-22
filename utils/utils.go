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
	X, Y, Width, Height float64
}

func NewHitBox(x, y, w, h float64) HitBox {
	return HitBox{X: x, Y: y, Width: w, Height: h}
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
	return num * (3.14159 / 180)
}

func Rad2Deg(num float64) float64 {
	return num * (180 / 3.14159)
}

func GetDist(point_1, point_2 Vec2) float64 {
	offx := math.Abs(point_1.X - point_2.X)
	offy := math.Abs(point_1.Y - point_2.Y)

	return math.Sqrt((offx * offx) + (offy * offy))
}

func GetAngle(point_1, point_2 Vec2) float64 {
	offset_x := point_1.X - point_2.X
	offset_y := point_1.Y - point_2.Y

	return math.Atan2(offset_x, offset_y)
}

func Raycast(pos, dir Vec2, length int, hitboxes []HitBox) (Vec2, bool) {
	for _, hitbox := range hitboxes {
		for l := range length {
			offx := dir.X * float64(l)
			offy := dir.Y * float64(l)

			if Collide(Vec2{pos.X + offx, pos.Y + offy}, Vec2{1, 1}, Vec2{hitbox.X, hitbox.Y}, Vec2{hitbox.Width, hitbox.Height}) {
				return Vec2{pos.X + offx, pos.Y + offy}, true
			}
		}
	}

	return Vec2{0, 0}, false
}
