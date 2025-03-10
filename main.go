package main

import (
	"jjb/scenes"
	"jjb/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (game *Game) Update() error {
	rmx, rmy := ebiten.CursorPosition()
	utils.Mouse_X, utils.Mouse_Y = float64(rmx), float64(rmy)

	if scenes.List_Of_Scenes[scenes.Current_Scene].Setup_run == false {
		scenes.List_Of_Scenes[scenes.Current_Scene].Setup()
		scenes.List_Of_Scenes[scenes.Current_Scene].Setup_run = true
	}

	scenes.List_Of_Scenes[scenes.Current_Scene].Update()

	return nil
}

var display_img = ebiten.NewImage(1280, 720)

func (game *Game) Draw(screen *ebiten.Image) {
	scenes.List_Of_Scenes[scenes.Current_Scene].Draw(display_img, screen)
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
