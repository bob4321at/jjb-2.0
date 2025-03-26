package players

import (
	"jjb/camera"
	"jjb/textures"
	"jjb/utils"
)

func (p *Player) hermesJump() {
	p.Vel.Y = -7
	p.NewProjectile(utils.Vec2{X: p.Pos.X - 64, Y: p.Pos.Y + 32}, utils.Vec2{X: 0, Y: 0}, 10, 0, 1, 3, textures.NewAnimatedTexture("./art/projectiles/hermes/cloud_jump.png"))
}

func (p *Player) hermesTornado() {
	p.NewProjectile(utils.Vec2{X: p.Pos.X - 64, Y: p.Pos.Y - 64}, utils.Vec2{X: 0, Y: 1}, 1, 10, 100, 100, textures.NewAnimatedTexture("./art/projectiles/hermes/tornado.png"))
}

func (player *Player) hermesWindBall() {
	player.NewProjectile(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y}, utils.Vec2{X: player.Pos.X + (player.Vel.X * 2) - utils.Mouse_X - camera.Cam.Offset.X + 640 + (float64(player.Img.GetTexture().Bounds().Dx())), Y: player.Pos.Y + (player.Vel.Y * 2) - utils.Mouse_Y - camera.Cam.Offset.Y + 320 + (float64(player.Img.GetTexture().Bounds().Dy()))}, 2, 10, 1, -1, textures.NewTexture("./art/projectiles/hermes/wind_ball.png", ""))
}

var hermes_attacks = []Attack{
	{Player_Ref.hermesJump, 0, 10},
	{Player_Ref.hermesWindBall, 0, 1},
	{Player_Ref.hermesTornado, 0, 20},
}
