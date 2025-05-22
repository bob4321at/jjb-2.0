package scenes

import (
	"jjb/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

var Story_Scene = NewScene(3, Story_Scene_Draw, Story_Scene_Update, Story_Scene_Setup)

func Story_Scene_Setup() {}

var Story_Tex = textures.NewAnimatedTexture("./art/ui/story.png", "")

func Story_Scene_Draw(display_image *ebiten.Image, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)

	Story_Tex.Draw(display_image, op)

	screen.DrawImage(display_image, &ebiten.DrawImageOptions{})
}

func Story_Scene_Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !utils.Clicked {
		Story_Tex.Current_Animation += 1
		if Story_Tex.Current_Animation+1 > len(Story_Tex.Animations) {
			Current_Scene += 1
		}
		utils.Clicked = true
	}
}
