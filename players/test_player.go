package players

import (
	"jjb/camera"
	"jjb/textures"
	"jjb/utils"
)

func (player *Player) test_playerGlitchProjectile() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 25, 20, 1, -1, textures.NewTexture("./art/projectiles/test_player/glitch.png", ""))
}

func (player *Player) test_playerDeathBar() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 256, Y: player.Pos.Y - 512}, utils.Vec2{X: 0, Y: -1}, 30, 5, 5, 100, textures.NewTexture("./art/projectiles/test_player/death_bar.png", ""))
}

func (player *Player) test_playerTp() {
	player.Pos.X += utils.Mouse_X - (1280 / 2)
	player.Pos.Y += utils.Mouse_Y - (720 / 2)
}

var test_player_attacks = []Attack{
	{Player_Ref.test_playerTp, 0, 3},
	{Player_Ref.test_playerGlitchProjectile, 0, 10},
	{Player_Ref.test_playerDeathBar, 0, 40},
}
