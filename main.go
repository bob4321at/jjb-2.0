package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var enemy_spawned bool = false

var clicked bool = false

var mouse_x, mouse_y float64

var empty_key ebiten.Key

var set_upped_yet bool = false

func (g *Game) Setup() {
	test_place = makeLevel("./maps/test_area.png", "./art/background.png")
	player = players["temp"]

	set_upped_yet = true
}

func (g *Game) Update() error {
	if !set_upped_yet {
		g.Setup()
	}
	rmx, rmy := ebiten.CursorPosition()
	mouse_x, mouse_y = float64(rmx), float64(rmy)

	current_level = &test_place

	if !test_place.generated {
		test_place = makeLevel("./maps/test_area.png", "./art/background.png")
	}
	player.Update()
	if ebiten.IsKeyPressed(ebiten.KeyX) && !enemy_spawned {
		test_place.Spawn(newEnemy(1, 5, Vec2{}, "./art/enemies/fliehead.png"))
		enemy_spawned = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyC) && !enemy_spawned {
		test_place.Spawn(newEnemy(2, 10, Vec2{}, "./art/enemies/crooked.png"))
		enemy_spawned = true
	}
	if !ebiten.IsKeyPressed(ebiten.KeyX) && !ebiten.IsKeyPressed(ebiten.KeyC) {
		enemy_spawned = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyG) {
		player = players["greg"]
	} else if ebiten.IsKeyPressed(ebiten.KeyJ) {
		player = players["gojo"]
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		clicked = false
	}

	camera.offset.x = player.pos.x
	camera.offset.y = player.pos.y
	test_place.Update(&player)
	return nil
}

var display_img = ebiten.NewImage(1280, 720)

func (g *Game) Draw(s *ebiten.Image) {
	display_img.Fill(color.RGBA{0, 115, 255, 255})
	test_place.Draw(display_img, &camera)
	player.Draw(display_img)
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.5, 1.5)
	op.GeoM.Translate(-float64(s.Bounds().Dx()/4), -float64(s.Bounds().Dy()/4))
	s.DrawImage(display_img, &op)
}

func (g *Game) Layout(ow, oh int) (sw, sh int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
