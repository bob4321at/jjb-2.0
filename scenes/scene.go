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

var List_Of_Scenes = []Scene{Start_Menu_Scene, Story_Scene, Charecter_Menu_Scene, Game_Scene, Death_Menu_Scene}
var Current_Scene = 0
var Old_Scene = Current_Scene
