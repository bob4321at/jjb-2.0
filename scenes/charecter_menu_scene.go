package scenes

import (
	"jjb/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var charecter_menu_img, _, _ = ebitenutil.NewImageFromFile("./art/ui/charecter_menu.png")
var Charecter_Menu_Scene = NewScene(0, Charecter_Menu_Scene_Draw, Charecter_Menu_Scene_Update, Charecter_Menu_Scene_Setup)

func Charecter_Menu_Scene_Setup() {}

func Charecter_Menu_Scene_Draw(display_img *ebiten.Image, screen *ebiten.Image) {
	display_img.DrawImage(charecter_menu_img, &ebiten.DrawImageOptions{})
	screen.DrawImage(display_img, &ebiten.DrawImageOptions{})
}

func Charecter_Menu_Scene_Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !utils.Clicked {
		if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 83, Y: 155}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "greg"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 269, Y: 155}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "gojo"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 455, Y: 155}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "megumi"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 641, Y: 155}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "boberto"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 827, Y: 155}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "jerry"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 1012, Y: 155}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "hermes"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 83, Y: 234}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "test_player"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 269, Y: 234}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "agent_21"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 455, Y: 234}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "tk"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 641, Y: 234}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "pyro"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 827, Y: 234}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "toothbrush_guy"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 1012, Y: 234}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "birdman"
			Game_Scene.Setup()
		} else if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 83, Y: 313}, utils.Vec2{X: 181, Y: 74}) {
			Current_Scene = 3
			selected_player = "tim"
			Game_Scene.Setup()
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) && ebiten.IsKeyPressed(ebiten.KeyR) {
			Current_Scene = 3
			selected_player = "sukuna"
			Game_Scene.Setup()
		}
	}
}
