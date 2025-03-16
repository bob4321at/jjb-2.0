package scenes

import (
	"jjb/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Death_Menu_Scene = NewScene(1, Death_Menu_Scene_Draw, Death_Menu_Scene_Update, func() {})

var death_menu_img, _, _ = ebitenutil.NewImageFromFile("./art/ui/death_menu.png")

func Death_Menu_Scene_Draw(display_img *ebiten.Image, screen *ebiten.Image) {
	display_img.DrawImage(death_menu_img, &ebiten.DrawImageOptions{})
	screen.DrawImage(display_img, &ebiten.DrawImageOptions{})
}

func Death_Menu_Scene_Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !utils.Clicked {
		Current_Scene = 0
		utils.Clicked = true
	}
}
