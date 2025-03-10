package scenes

import "github.com/hajimehoshi/ebiten/v2"

type Scene struct {
	Id        int
	Draw      func(s *ebiten.Image, screen_img *ebiten.Image)
	Update    func()
	Setup     func()
	Setup_run bool
}

func NewScene(id int, draw func(display_img *ebiten.Image, screen_img *ebiten.Image), update func(), setup func()) (scene Scene) {
	scene.Id = id

	scene.Draw = draw
	scene.Update = update
	scene.Setup = setup

	scene.Setup_run = false

	return scene
}
