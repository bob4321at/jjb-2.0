package players

import (
	"jjb/camera"
	"jjb/enemyai"
	"jjb/shaders"
	"jjb/textures"
	"jjb/utils"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Pos               utils.Vec2
	Vel               utils.Vec2
	Img               textures.AnimatedTexture
	Health            int
	Dir               bool
	Attacks           []Attack
	Damage_Multiplier float64
	Projectiles       []Projectile
	Entities          []PlayerEntity
	Domain            Domain
	Domain_Timer      float64
	I_Frames          float64
	Player_Name       string
}

type Projectile struct {
	Pos      utils.Vec2
	Vel      utils.Vec2
	Damage   int
	Speed    float64
	Pierce   float64
	Lifetime float64
	Img      textures.RenderableTexture
}

type Domain struct {
	Img    textures.RenderableTexture
	Effect func(enemies []*enemyai.Enemy)
}

type Attack struct {
	Attack       func()
	Cooldown     float64
	Max_Cooldown float64
}

var attack_keys = map[int]ebiten.Key{
	0: ebiten.KeyE,
	2: ebiten.KeyF,
}

func (p *Player) NewProjectile(pos, vel utils.Vec2, damage int, speed float64, pierce float64, lifetime float64, img textures.RenderableTexture) {
	projectile := Projectile{pos, vel, damage, speed, pierce, lifetime, img}

	p.Projectiles = append(p.Projectiles, projectile)
}

func (p *Player) NewDomain(img textures.RenderableTexture, effect func(l []*enemyai.Enemy)) (d Domain) {
	d.Img = img
	d.Effect = effect

	return d
}

type DomainedEnemy struct {
	enemy     *enemyai.Enemy
	alive     bool
	start_pos utils.Vec2
}

func (p *Player) simpleDomain(enemies []*enemyai.Enemy) {
	affected := []DomainedEnemy{}
	player_start_pos := p.Pos

	for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
		e := enemyai.Enemies_In_World[enemy_index]
		affected = append(affected, DomainedEnemy{e, true, e.Pos})
		if utils.Collide(utils.Vec2{X: p.Pos.X - 1024, Y: p.Pos.Y - 1024}, utils.Vec2{X: 2048, Y: 2048}, e.Pos, utils.Vec2{X: float64(e.Tex.GetTexture().Bounds().Dx()), Y: float64(e.Tex.GetTexture().Bounds().Dy())}) {
			e.Pos.X = 1800 + (rand.Float64() * 1000)
			e.Pos.Y = -1700 - (rand.Float64() * 300)
		}
	}
	p.Pos.X = 2000
	p.Pos.Y = -1600

	start_time := utils.Game_Time

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		e := affected[enemy_index].enemy
		e.Can_Move = true
	}

	for start_time+1500 > utils.Game_Time {
		for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
			de := affected[enemy_index]
			if de.enemy.Health < 0 {
				de.alive = false
			}
		}
	}

	p.Pos = player_start_pos

	for enemy_index := 0; enemy_index < len(affected); enemy_index++ {
		de := affected[enemy_index]
		de.enemy.Pos = de.start_pos
	}
}

func newPlayer(pos utils.Vec2, img textures.AnimatedTexture, domain_img textures.RenderableTexture, domain_effect func(l []*enemyai.Enemy), attacks []Attack) (p Player) {
	p.Pos = pos
	p.Vel = utils.Vec2{X: 0, Y: 0}

	p.Img = img

	p.Health = 100

	p.Attacks = attacks
	p.Dir = false
	p.Damage_Multiplier = 1

	domain := p.NewDomain(domain_img, domain_effect)
	p.Domain = domain
	p.Domain_Timer = 0

	return p
}

func (p *Player) Punch() {
	for ie := 0; ie < len(enemyai.Enemies_In_World); ie++ {
		e := enemyai.Enemies_In_World[ie]
		if utils.Collide(utils.Vec2{X: p.Pos.X - 32, Y: p.Pos.Y}, utils.Vec2{X: 96, Y: 64}, e.Pos, utils.Vec2{X: float64(e.Tex.GetTexture().Bounds().Dx()), Y: float64(e.Tex.GetTexture().Bounds().Dy())}) {
			e.Health -= 1
		}
	}
}

func (player *Player) DamageCheck() {
	if player.I_Frames <= 0 {
		for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
			enemy := enemyai.Enemies_In_World[enemy_index]
			if utils.Collide(player.Pos, utils.Vec2{X: 32, Y: 64}, enemy.Pos, utils.Vec2{X: float64(enemy.Tex.GetTexture().Bounds().Dx()), Y: float64(enemy.Tex.GetTexture().Bounds().Dy())}) {
				player.I_Frames = 2
				player.Health -= enemy.Damage
			}
			for projectile_index := 0; projectile_index < len(enemy.Projectiles); projectile_index++ {
				projectile := &enemy.Projectiles[projectile_index]

				projectile.Pos.X += projectile.Vel.X
				projectile.Pos.Y += projectile.Vel.Y

				if utils.Collide(player.Pos, utils.Vec2{X: 32, Y: 64}, projectile.Pos, utils.Vec2{X: float64(projectile.Img.GetTexture().Bounds().Dx()), Y: float64(projectile.Img.GetTexture().Bounds().Dy())}) {
					player.Health -= projectile.Damage
					utils.RemoveArrayElement(projectile_index, &enemy.Projectiles)
					projectile_index -= 1
				}

				projectile.Lifetime -= 0.1
				if projectile.Lifetime < 0 {
					utils.RemoveArrayElement(projectile_index, &enemy.Projectiles)
					projectile_index -= 1
				}
			}
		}
	} else {
		player.I_Frames -= 0.1
	}
}

func (p *Player) Draw(s *ebiten.Image) {
	op := ebiten.DrawImageOptions{}

	op.GeoM.Reset()

	p.Img.SetUniforms(map[string]any{
		"I_Frames": p.I_Frames,
	})

	if !p.Dir {
		op.GeoM.Translate(640-camera.Cam.Manual_Offset.X, 360-camera.Cam.Manual_Offset.Y)
		p.Img.Draw(s, &op)
	} else {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(640+32-camera.Cam.Manual_Offset.X, 360-camera.Cam.Manual_Offset.Y)
		p.Img.Draw(s, &op)
	}

	op.GeoM.Reset()

	for entity_index := 0; entity_index < len(p.Entities); entity_index++ {
		e := &p.Entities[entity_index]
		op := ebiten.DrawImageOptions{}
		if !e.Dir {
			op.GeoM.Translate(-(float64(e.Img.GetTexture().Bounds().Dx()))/2, -(float64(e.Img.GetTexture().Bounds().Dy()))/2)
			op.GeoM.Rotate(utils.Deg2Rad(e.Rotation))
			op.GeoM.Translate((float64(e.Img.GetTexture().Bounds().Dx()))/2, (float64(e.Img.GetTexture().Bounds().Dy()))/2)
			op.GeoM.Translate(e.Pos.X-camera.Cam.Offset.X+640-camera.Cam.Manual_Offset.X, e.Pos.Y-camera.Cam.Offset.Y+360-camera.Cam.Manual_Offset.Y)
		} else {
			op.GeoM.Translate(-(float64(e.Img.GetTexture().Bounds().Dx()))/2, -(float64(e.Img.GetTexture().Bounds().Dy()))/2)
			op.GeoM.Rotate(utils.Deg2Rad(e.Rotation))
			op.GeoM.Translate((float64(e.Img.GetTexture().Bounds().Dx()))/2, (float64(e.Img.GetTexture().Bounds().Dy()))/2)
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(e.Pos.X-camera.Cam.Offset.X-camera.Cam.Manual_Offset.X+640+float64(e.Img.GetTexture().Bounds().Dx()), e.Pos.Y-camera.Cam.Offset.Y-camera.Cam.Manual_Offset.Y+360)
		}
		s.DrawImage(e.Img.GetTexture(), &op)
	}

	for projectile_index := 0; projectile_index < len(p.Projectiles); projectile_index++ {
		op.GeoM.Reset()
		op.GeoM.Translate(p.Projectiles[projectile_index].Pos.X-camera.Cam.Offset.X-camera.Cam.Manual_Offset.X+650, p.Projectiles[projectile_index].Pos.Y-camera.Cam.Offset.Y-camera.Cam.Manual_Offset.Y+380)
		p.Projectiles[projectile_index].Img.Draw(s, &op)
	}
}

func (p *Player) Update(level_hitbox []utils.HitBox) {
	p.Img.Update()
	p.Img.RefreshTexture()

	if p.I_Frames > 0 {
	}

	p.Vel.Y += 0.1
	if p.Vel.X != 0 {
		p.Img.Current_Animation = 1
	} else {
		p.Img.Current_Animation = 0
	}

	if p.Vel.X > 5 {
		p.Vel.X -= 0.1
		if p.Vel.X > 10 {
			p.Vel.X -= 0.2
		}
	} else if p.Vel.X < -5 {
		p.Vel.X += 0.1
		if p.Vel.X < -10 {
			p.Vel.X += 0.2
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Vel.X -= 0.1
		p.Dir = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Vel.X += 0.1
		p.Dir = false
	} else {
		if p.Vel.X > 0 {
			p.Vel.X -= 0.2
			if p.Vel.X > -0.6 && p.Vel.X < 0.6 {
				p.Vel.X = 0
			}
		} else if p.Vel.X < 0 {
			p.Vel.X += 0.2
			if p.Vel.X > -0.6 && p.Vel.X < 0.6 {
				p.Vel.X = 0
			}
		}
	}

	for b := 0; b < len(p.Attacks); b++ {
		p.Attacks[b].Cooldown -= 0.1
		if p.Attacks[b].Cooldown < 0 {
			p.Attacks[b].Cooldown = 0
		}
		if ebiten.IsKeyPressed(attack_keys[b]) && p.Attacks[b].Cooldown <= 0 && attack_keys[b] != utils.Empty_Key {
			p.Attacks[b].Attack()
			p.Attacks[b].Cooldown = p.Attacks[b].Max_Cooldown
		} else if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) && p.Attacks[1].Cooldown <= 0 {
			p.Attacks[1].Attack()
			p.Attacks[1].Cooldown = p.Attacks[1].Max_Cooldown
		}
	}

	if p.Domain_Timer < 0 {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			go p.Domain.Effect(enemyai.Enemies_In_World)
			p.Domain_Timer = 360
		}
	} else {
		p.Domain_Timer -= 0.1
	}

	if utils.Collide(utils.Vec2{X: p.Pos.X, Y: p.Pos.Y + p.Vel.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		p.Vel.Y = 0
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
			if utils.Collide(utils.Vec2{X: p.Pos.X, Y: p.Pos.Y + p.Vel.Y - 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000, Y: -2000 + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
				p.Vel.Y = 0
			} else {
				p.Vel.Y = -5.1
			}
		}
	}
	if utils.Collide(utils.Vec2{X: p.Pos.X, Y: p.Pos.Y + p.Vel.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) - (250)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		p.Vel.Y = 0
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
			if utils.Collide(utils.Vec2{X: p.Pos.X, Y: p.Pos.Y + p.Vel.Y - 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000, Y: -2000 + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
				p.Vel.Y = 0
			} else {
				p.Vel.Y = -5.1
			}
		}
	}

	if utils.Collide(utils.Vec2{X: p.Pos.X + p.Vel.X, Y: p.Pos.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		p.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: p.Pos.X + p.Vel.X, Y: p.Pos.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		p.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: p.Pos.X + p.Vel.X, Y: p.Pos.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000 + 2048 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		p.Vel.X = 0
	}

	for tile_index := 0; tile_index < len(level_hitbox); tile_index++ {
		t := level_hitbox[tile_index]
		if utils.Collide(utils.Vec2{X: p.Pos.X, Y: p.Pos.Y + p.Vel.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			p.Vel.Y = 0
			if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
				if utils.Collide(utils.Vec2{X: p.Pos.X, Y: p.Pos.Y + p.Vel.Y - 2}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
					p.Vel.Y = 0
				} else {
					p.Vel.Y = -5.1
				}
			}
		}
		if utils.Collide(utils.Vec2{X: p.Pos.X + p.Vel.X, Y: p.Pos.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: t.X, Y: t.Y}, utils.Vec2{X: 32, Y: 32}) {
			p.Vel.X = 0
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !utils.Clicked {
		p.Punch()
		utils.Clicked = true
	}

	for projectile_index := 0; projectile_index < len(p.Projectiles); projectile_index++ {
		projectile := &p.Projectiles[projectile_index]
		projectile_move_dir := math.Atan2(projectile.Vel.Y, projectile.Vel.X)
		projectile.Pos.X -= math.Cos(projectile_move_dir) * projectile.Speed
		projectile.Pos.Y -= math.Sin(projectile_move_dir) * projectile.Speed

		projectile.Img.Update()

		for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
			e := enemyai.Enemies_In_World[enemy_index]
			if e.I_Frames == 0 {
				if utils.Collide(projectile.Pos, utils.Vec2{X: float64(projectile.Img.GetTexture().Bounds().Dx()), Y: float64(projectile.Img.GetTexture().Bounds().Dy())}, e.Pos, utils.Vec2{X: float64(e.Tex.GetTexture().Bounds().Dx()), Y: float64(e.Tex.GetTexture().Bounds().Dy())}) {
					e.DoDamage(projectile.Damage)
					if projectile.Pierce == -1 {
						utils.RemoveArrayElement(projectile_index, &p.Projectiles)
						break
					} else {
						projectile.Pierce -= 1.1
						if projectile.Pierce <= 0 {
							projectile.Damage = 0
						}
					}
				}
			}
		}
		if projectile.Lifetime != -1 {
			projectile.Lifetime -= 0.1
			if projectile.Lifetime < 0 {
				utils.RemoveArrayElement(projectile_index, &p.Projectiles)
			}
		}
	}

	for entity_index := 0; entity_index < len(p.Entities); entity_index++ {
		e := &p.Entities[entity_index]
		e.Update(e, level_hitbox)

		if e.Lifespan < 0 {
			utils.RemoveArrayElement(entity_index, &p.Entities)
		}
	}

	p.Pos.Y += p.Vel.Y
	p.Pos.X += p.Vel.X
}

var Player_Ref Player

func init() {
	Player_Ref = newPlayer(utils.Vec2{X: 0, Y: 0}, *textures.NewAnimatedTexture("./art/players/greg.png", ""), textures.NewTexture("./art/domains/simple_domain.png", shaders.Player_Shader), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, greg_attacks)
}
