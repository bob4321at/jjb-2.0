package players

import (
	"jjb/camera"
	"jjb/textures"
	"jjb/utils"
)

func (player *Player) gregLaunch() {
	player.Vel.Y = -10
	if !player.Dir {
		player.Vel.X = 10
	} else {
		player.Vel.X = -10
	}
}

func (player *Player) gregThrow() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 5, 10, 1, -1, textures.NewTexture("./art/projectiles/greg/rock.png"))
}

func (player *Player) gregNuke() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X - 128, Y: player.Pos.Y - 128}, utils.Vec2{X: 0, Y: 0}, 1, 0, 5, 10, textures.NewTexture("./art/projectiles/greg/explosion.png"))
}

var greg_attacks = []Attack{
	{Player_Ref.gregLaunch, 0, 20},
	{Player_Ref.gregThrow, 0, 5},
	{Player_Ref.gregNuke, 0, 1},
}
