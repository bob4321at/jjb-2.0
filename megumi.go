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
	p.newEntity(Vec2{p.pos.x, p.pos.y}, Vec2{1, 0}, "./art/entities/megumi/mahoraga.png", mahoragaUpdate)
}

func mahoragaUpdate(e *PlayerEntity) {
	e.vel.y += 0.1

	for tile_index := 0; tile_index < len(current_level.tiles); tile_index++ {
		t := &current_level.tiles[tile_index]
		if collide(Vec2{e.pos.x, e.pos.y + e.vel.y}, Vec2{float64(e.img.Bounds().Dx()), float64(e.img.Bounds().Dy())}, t.pos, Vec2{32, 32}) {
			if e.vel.y >= 0 {
				e.vel.y = -3
			} else {
				e.vel.y = 0
			}
		}
		if collide(Vec2{e.pos.x + e.vel.x, e.pos.y}, Vec2{float64(e.img.Bounds().Dx()), float64(e.img.Bounds().Dy())}, t.pos, Vec2{32, 32}) {
			e.vel.x = -e.vel.x
		}
	}

	e.pos.x += e.vel.x
	e.pos.y += e.vel.y
}

var megumi_attacks = []Attack{
	{player.megumiTp, 0, 4},
	{player.megumiBird, 0, 4},
	{player.megumiMahoraga, 0, 4},
}
