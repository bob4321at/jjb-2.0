package scenes

import (
	"image/color"
	"jjb/camera"
	"jjb/level"
	"jjb/players"
	"jjb/ui"
	"jjb/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

var Game_Scene = NewScene(1, Game_Scene_Draw, Game_Scene_Update, Game_Scene_Setup)

func Game_Scene_Setup() {
	level.Levels = level.LoadAllLevels("./maps/")
	level.Current_Level_Index = 0
	level.Current_Level = &level.Levels[level.Current_Level_Index]
	players.InitPlayer(level.Current_Level.Player_Spawn)

	players.Player_Ref = players.Players["greg"]
}

func Game_Scene_Draw(display_img *ebiten.Image, screen *ebiten.Image) {
	display_img.Fill(color.RGBA{0, 115, 255, 255})
	level.Current_Level.Draw(display_img, &camera.Cam)
	players.Player_Ref.Draw(display_img)
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.5, 1.5)
	op.GeoM.Translate(-float64(screen.Bounds().Dx()/4), -float64(screen.Bounds().Dy()/4))
	screen.DrawImage(display_img, &op)
	ui.DrawUi(screen)
}

func Game_Scene_Update() {
	if &level.Levels[level.Current_Level_Index] != level.Current_Level {
		level.Current_Level = &level.Levels[level.Current_Level_Index]
		players.InitPlayer(level.Current_Level.Player_Spawn)
		players.Player_Ref = players.Players["greg"]
	}

	utils.Game_Time += 1

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		utils.Clicked = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyG) {
		players.Player_Ref = players.Players["greg"]
	} else if ebiten.IsKeyPressed(ebiten.KeyJ) {
		players.Player_Ref = players.Players["gojo"]
	} else if ebiten.IsKeyPressed(ebiten.KeyM) {
		players.Player_Ref = players.Players["megumi"]
	} else if ebiten.IsKeyPressed(ebiten.KeyB) {
		players.Player_Ref = players.Players["boberto"]
	} else if ebiten.IsKeyPressed(ebiten.KeyN) {
		players.Player_Ref = players.Players["jerry"]
	}

	camera.Cam.Offset.X = players.Player_Ref.Pos.X
	camera.Cam.Offset.Y = players.Player_Ref.Pos.Y
	level.Current_Level.Update(&players.Player_Ref)
}
