package main

func (p *Player) megumiBird() {
	if p.dir {
		p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{-3, 0}, 1, 1, 5, 100, "./art/projectiles/megumi/birdright.png")
	} else {
		p.newProjectile(Vec2{p.pos.x, p.pos.y}, Vec2{3, 0}, 1, 1, 5, 100, "./art/projectiles/megumi/birdleft.png")
	}
}

var megumi_attacks = []Attack{
	{player.megumiBird, 0, 4},
	{player.megumiBird, 0, 4},
	{player.megumiBird, 0, 4},
}
