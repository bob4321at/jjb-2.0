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
var selected_player = "greg"

var tab_key_hit = false

var amount_of_inits = 0

func Game_Scene_Setup() {
	level.Levels = append(level.Levels, level.LoadLevel("./maps/level01/"))
	level.Current_Level_Index = 0
	level.Current_Level = &level.Levels[level.Current_Level_Index]
	level.Current_Level.Spawned = false
	go level.LoadAllLevels("./maps/", &level.Levels)
	players.InitPlayer(level.Current_Level.Player_Spawn)
	players.Player_Ref = players.Players[selected_player]
	players.Player_Ref.Player_Name = selected_player
}

func Game_Scene_Draw(display_img *ebiten.Image, screen *ebiten.Image) {
	display_img.Fill(color.RGBA{0, 0, 0, 255})
	level.Current_Level.Draw(display_img, &camera.Cam)
	players.Player_Ref.Draw(display_img)
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.5, 1.5)
	op.GeoM.Translate(-float64(screen.Bounds().Dx()/4), -float64(screen.Bounds().Dy()/4))
	screen.DrawImage(display_img, &op)
	ui.DrawUi(screen)
}

func Game_Scene_Update() {
	camera.Cam.Update()

	if !ebiten.IsKeyPressed(ebiten.KeyTab) {
		tab_key_hit = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyTab) && !tab_key_hit {
		tab_key_hit = true
		level.Current_Level_Index += 1
		if &level.Levels[level.Current_Level_Index] != level.Current_Level {
			level.Current_Level = &level.Levels[level.Current_Level_Index]
			players.Player_Ref = players.Players[selected_player]
			players.InitPlayer(level.Current_Level.Player_Spawn)
			players.Player_Ref = players.Players[selected_player]
			players.Player_Ref.Attacks[0].Cooldown = 0
			players.Player_Ref.Attacks[1].Cooldown = 0
			players.Player_Ref.Attacks[2].Cooldown = 0
		}
	}

	if level.Current_Level.Beaten == true {
		level.Current_Level_Index += 1
		if &level.Levels[level.Current_Level_Index] != level.Current_Level {
			level.Current_Level = &level.Levels[level.Current_Level_Index]
			players.Player_Ref = players.Players[selected_player]
			players.InitPlayer(level.Current_Level.Player_Spawn)
			players.Player_Ref = players.Players[selected_player]
			players.Player_Ref.Attacks[0].Cooldown = 0
			players.Player_Ref.Attacks[1].Cooldown = 0
			players.Player_Ref.Attacks[2].Cooldown = 0
		}
	}

	if players.Player_Ref.Health <= 0 {
		Current_Scene = 4
		players.Player_Ref.Health = 100
	}

	utils.Game_Time += 1

	camera.Cam.Offset.X = players.Player_Ref.Pos.X
	camera.Cam.Offset.Y = players.Player_Ref.Pos.Y
	level.Current_Level.Update(&players.Player_Ref)

}
