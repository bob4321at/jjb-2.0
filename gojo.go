package main

func (p *Player) gojoRed() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x - mouse_x - camera.offset.x + 640 + (float64(p.img.Bounds().Dx())), p.pos.y - mouse_y - camera.offset.y + 320 + (float64(p.img.Bounds().Dy()))}, 5, 5, -1, "./art/projectiles/gojo/red.png")
}

func (p *Player) gojoPurple() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x - mouse_x - camera.offset.x + 640 + (float64(p.img.Bounds().Dx())), p.pos.y - mouse_y - camera.offset.y + 320 + (float64(p.img.Bounds().Dy()))}, 1, 7, 40, "./art/projectiles/gojo/purple.png")
}

func (p *Player) gojoBlue() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x - mouse_x - camera.offset.x + 640 + (float64(p.img.Bounds().Dx())), p.pos.y - mouse_y - camera.offset.y + 320 + (float64(p.img.Bounds().Dy()))}, 1, 10, 4, "./art/projectiles/gojo/blue.png")
}

var gojo_attacks = []Attack{
	{player.gojoBlue, 0, 3},
	{player.gojoRed, 0, 5},
	{player.gojoPurple, 0, 20},
}
