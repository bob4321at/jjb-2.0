package main

func (p *Player) megumiTp() {
	p.pos.x = ((mouse_x - 640) / 2) + (p.pos.x - 16)
	p.pos.y = ((mouse_y - 480) / 2) + (p.pos.y + 32)
}

func (p *Player) megumiBird() {
	if p.pos.x-camera.offset.x+640 < mouse_x {
		p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{-1, 0}, 1, 3, 5, 100, "./art/projectiles/megumi/birdright.png")
	} else {
		p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{1, 0}, 1, 3, 5, 100, "./art/projectiles/megumi/birdleft.png")
	}
}

func (p *Player) megumiMahoraga() {
	p.newEntity(Vec2{p.pos.x - camera.offset.x + 640, p.pos.y - camera.offset.y + 320}, "./art/entities/megumi/mahoraga.png", mahoragaUpdate)
}

func mahoragaUpdate(e *PlayerEntity) {
	e.pos.y -= 1
}

var megumi_attacks = []Attack{
	{player.megumiTp, 0, 4},
	{player.megumiBird, 0, 4},
	{player.megumiMahoraga, 0, 4},
}
