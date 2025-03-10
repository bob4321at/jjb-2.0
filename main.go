package main

import (
	"jjb/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var enemy_spawned bool = false

var started = false

var list_of_scenes = []scenes.Scene{scenes.Game_Scene}
var current_scene = 0

func (game *Game) Update() error {
	if list_of_scenes[current_scene].Setup_run == false {
		list_of_scenes[current_scene].Setup()
		list_of_scenes[current_scene].Setup_run = true
	}

	list_of_scenes[current_scene].Update()

	return nil
}

var display_img = ebiten.NewImage(1280, 720)

func (game *Game) Draw(screen *ebiten.Image) {
	list_of_scenes[current_scene].Draw(display_img, screen)
}

func (game *Game) Layout(origonal_width, origonal_height int) (screen_width, screen_height int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
