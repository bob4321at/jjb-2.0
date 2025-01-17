package main

func (p *Player) megumiTp() {
	p.pos.x += mouse_x - (1280 / 2)
	p.pos.y += mouse_y - (720 / 2)
}

func (p *Player) megumiBird() {
	if p.pos.x-camera.offset.x+640 < mouse_x {
		p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{-1, 0.5}, 1, 3, 5, 100, "./art/projectiles/megumi/birdright.png")
	} else {
		p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{1, 0.5}, 1, 3, 5, 100, "./art/projectiles/megumi/birdleft.png")
	}
}

func (p *Player) megumiMahoraga() {
	if !p.dir {
		p.newEntity(Vec2{p.pos.x - 16, p.pos.y - 32}, Vec2{1, 0}, 1, "./art/entities/megumi/mahoraga.png", mahoragaUpdate)
	} else {
		p.newEntity(Vec2{p.pos.x - 16, p.pos.y - 32}, Vec2{-1, 0}, 1, "./art/entities/megumi/mahoraga.png", mahoragaUpdate)
	}
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

	if e.vel.x > 0 {
		e.dir = false
	} else {
		e.dir = true
	}

	if e.cooldown < 0 {
		for enemy_index := 0; enemy_index < len(current_level.enemies); enemy_index++ {
			le := &current_level.enemies[enemy_index]
			if collide(e.pos, Vec2{float64(e.img.Bounds().Dx()), float64(e.img.Bounds().Dy())}, le.pos, Vec2{float64(le.img.Bounds().Dx()), float64(le.img.Bounds().Dy())}) {
				le.health -= 1
				e.cooldown = 1
			}
		}
	} else {
		e.cooldown -= 0.1
	}

	e.pos.x += e.vel.x
	e.pos.y += e.vel.y
}

var megumi_attacks = []Attack{
	{player.megumiTp, 0, 4},
	{player.megumiBird, 0, 4},
	{player.megumiMahoraga, 0, 4},
}
