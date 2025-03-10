package scenes

import (
	"jjb/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Start_Menu_Scene = NewScene(0, Start_Menu_Scene_Draw, Start_Menu_Scene_Update, Start_Menu_Scene_Setup)

var start_menu_img, _, _ = ebitenutil.NewImageFromFile("./art/ui/start_menu.png")

func Start_Menu_Scene_Setup() {}

func Start_Menu_Scene_Draw(display_img *ebiten.Image, screen *ebiten.Image) {
	display_img.DrawImage(start_menu_img, &ebiten.DrawImageOptions{})
	screen.DrawImage(display_img, &ebiten.DrawImageOptions{})
}

func Start_Menu_Scene_Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		if utils.Collide(utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}, utils.Vec2{X: 1, Y: 1}, utils.Vec2{X: 94, Y: 188}, utils.Vec2{X: 411, Y: 121}) {
			Current_Scene = 1
		}
	}
}
