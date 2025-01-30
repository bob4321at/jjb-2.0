package main

func (p *Player) gregLaunch() {
	p.vel.y = -10
	if !p.dir {
		p.vel.x = 10
	} else {
		p.vel.x = -10
	}
}

func (p *Player) gregThrow() {
	p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{p.pos.x + (p.vel.x * 2) - mouse_x - camera.offset.x + 640 + (float64(p.img.Bounds().Dx())), p.pos.y + (p.vel.y * 2) - mouse_y - camera.offset.y + 320 + (float64(p.img.Bounds().Dy()))}, 5, 10, 1, -1, "./art/projectiles/greg/rock.png")
}

func (p *Player) gregNuke() {
	p.newProjectile(Vec2{p.pos.x - 128, p.pos.y - 128}, Vec2{0, 0}, 1, 0, 5, 10, "./art/projectiles/greg/explosion.png")
}

var greg_attacks = []Attack{
	{player.gregLaunch, 0, 20},
	{player.gregThrow, 0, 5},
	{player.gregNuke, 0, 1},
}
