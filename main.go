package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var enemy_spawned bool = false

var clicked bool = false

var mouse_x, mouse_y float64

var empty_key ebiten.Key

var started = false

var game_time float64 = 0

var domain_background, _, _ = ebitenutil.NewImageFromFile("./art/domains/domain_backdrop.png")

func (g *Game) Setup() {
	if !started {
		levels = loadAllLevels("./maps/")
		current_level_index = 0
		current_level = &levels[current_level_index]
		initPlayer()
		player = players["greg"]
	}

	started = true
}

func (g *Game) Update() error {
	g.Setup()

	if &levels[current_level_index] != current_level {
		current_level = &levels[current_level_index]
		initPlayer()
		player = players["greg"]
	}

	game_time += 0.1

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		clicked = false
	}

	rmx, rmy := ebiten.CursorPosition()
	mouse_x, mouse_y = float64(rmx), float64(rmy)

	player.Update()
	if ebiten.IsKeyPressed(ebiten.KeyX) && !enemy_spawned {
		current_level.Spawn(newEnemy(1, 5, Vec2{}, newAnimatedTexture("./art/enemies/fliehead.png")))
		enemy_spawned = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyC) && !enemy_spawned {
		current_level.Spawn(newEnemy(2, 10, Vec2{}, newTexture("./art/enemies/crooked.png")))
		enemy_spawned = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyZ) && !enemy_spawned {
		current_level.Spawn(newEnemy(3, 5, Vec2{}, newTexture("./art/enemies/shrimp.png")))
		enemy_spawned = true
	}
	if !ebiten.IsKeyPressed(ebiten.KeyX) && !ebiten.IsKeyPressed(ebiten.KeyC) && !ebiten.IsKeyPressed(ebiten.KeyZ) {
		enemy_spawned = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyG) {
		player = players["greg"]
	} else if ebiten.IsKeyPressed(ebiten.KeyJ) {
		player = players["gojo"]
	} else if ebiten.IsKeyPressed(ebiten.KeyM) {
		player = players["megumi"]
	} else if ebiten.IsKeyPressed(ebiten.KeyB) {
		player = players["boberto"]
	}

	camera.offset.x = player.pos.x
	camera.offset.y = player.pos.y
	current_level.Update(&player)

	return nil
}

var display_img = ebiten.NewImage(1280, 720)

func (g *Game) Draw(s *ebiten.Image) {
	display_img.Fill(color.RGBA{0, 115, 255, 255})
	current_level.Draw(display_img, &camera)
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
