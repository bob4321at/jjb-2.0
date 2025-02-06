package main

func (p *Player) gojoRed() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x + p.vel.x - mouse_x - camera.offset.x + 640 + (float64(p.img.getTexture().Bounds().Dx())), p.pos.y + p.vel.y - mouse_y - camera.offset.y + 320 + (float64(p.img.getTexture().Bounds().Dy()))}, 5, 5, 1, -1, "./art/projectiles/gojo/red.png")
}

func (p *Player) gojoBlue() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x + p.vel.x - mouse_x - camera.offset.x + 640 + (float64(p.img.getTexture().Bounds().Dx())), p.pos.y + p.vel.y - mouse_y - camera.offset.y + 320 + (float64(p.img.getTexture().Bounds().Dy()))}, 1, 10, 5, 5, "./art/projectiles/gojo/blue.png")
}

func (p *Player) gojoPurple() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x + p.vel.x - mouse_x - camera.offset.x + 640 + (float64(p.img.getTexture().Bounds().Dx())), p.pos.y + p.vel.y - mouse_y - camera.offset.y + 320 + (float64(p.img.getTexture().Bounds().Dy()))}, 1, 7, 40, 10, "./art/projectiles/gojo/purple.png")
}

var gojo_attacks = []Attack{
	{player.gojoBlue, 0, 3},
	{player.gojoRed, 0, 5},
	{player.gojoPurple, 0, 20},
}
