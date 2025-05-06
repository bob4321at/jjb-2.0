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
	Domain_Active     bool
	Domain_Start_Time float64
	Player_Return_Pos utils.Vec2
	Activate_Domain   bool
	Domained_Enemies  []DomainedEnemy
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

func (player *Player) NewProjectile(pos, vel utils.Vec2, damage int, speed float64, pierce float64, lifetime float64, img textures.RenderableTexture) {
	projectile := Projectile{pos, vel, damage, speed, pierce, lifetime, img}

	player.Projectiles = append(player.Projectiles, projectile)
}

func (player *Player) NewDomain(img textures.RenderableTexture, effect func(l []*enemyai.Enemy)) (d Domain) {
	d.Img = img
	d.Effect = effect

	return d
}

type DomainedEnemy struct {
	enemy     *enemyai.Enemy
	alive     bool
	start_pos utils.Vec2
}

func (player *Player) simpleDomain(enemies []*enemyai.Enemy) {
	if player.Activate_Domain {
		player.Domained_Enemies = []DomainedEnemy{}
		player.Player_Return_Pos = player.Pos

		for enemy_index := 0; enemy_index < len(enemyai.Enemies_In_World); enemy_index++ {
			e := enemyai.Enemies_In_World[enemy_index]
			player.Domained_Enemies = append(player.Domained_Enemies, DomainedEnemy{e, true, e.Pos})
			if utils.Collide(utils.Vec2{X: player.Pos.X - 1024, Y: player.Pos.Y - 1024}, utils.Vec2{X: 2048, Y: 2048}, e.Pos, utils.Vec2{X: float64(e.Tex.GetTexture().Bounds().Dx()), Y: float64(e.Tex.GetTexture().Bounds().Dy())}) {
				e.Return_To_Pos = e.Pos
				e.Pos.X = 1800 + (rand.Float64() * 1000)
				e.Pos.Y = -1700 - (rand.Float64() * 300)
			}
		}
		player.Pos.X = 2000
		player.Pos.Y = -1600

		player.Domain_Start_Time = utils.Game_Time

		for enemy_index := 0; enemy_index < len(player.Domained_Enemies); enemy_index++ {
			e := player.Domained_Enemies[enemy_index].enemy
			e.Can_Move = true
		}
		player.Activate_Domain = false
		player.Domain_Active = true
	}

	if player.Domain_Active && player.Domain_Start_Time+1499 < utils.Game_Time {
		player.Pos = player.Player_Return_Pos
		for _, e := range enemies {
			empty_vec2 := utils.Vec2{X: 0, Y: 0}
			if e.Return_To_Pos != empty_vec2 {
				e.Pos = e.Return_To_Pos
				e.Return_To_Pos = utils.Vec2{X: 0, Y: 0}
			}
		}
		player.Domain_Active = false
	}
}

func newPlayer(pos utils.Vec2, img textures.AnimatedTexture, domain_img textures.RenderableTexture, domain_effect func(level_enemies []*enemyai.Enemy), attacks []Attack) (player Player) {
	player.Pos = pos
	player.Vel = utils.Vec2{X: 0, Y: 0}

	player.Img = img

	player.Health = 100

	player.Attacks = attacks
	player.Dir = false
	player.Damage_Multiplier = 1

	domain := player.NewDomain(domain_img, domain_effect)
	player.Domain = domain
	player.Domain_Timer = 0

	return player
}

func (player *Player) Punch() {
	for ie := 0; ie < len(enemyai.Enemies_In_World); ie++ {
		e := enemyai.Enemies_In_World[ie]
		if utils.Collide(utils.Vec2{X: player.Pos.X - 32, Y: player.Pos.Y}, utils.Vec2{X: 96, Y: 64}, e.Pos, utils.Vec2{X: float64(e.Tex.GetTexture().Bounds().Dx()), Y: float64(e.Tex.GetTexture().Bounds().Dy())}) {
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

func (player *Player) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}

	op.GeoM.Reset()

	player.Img.SetUniforms(map[string]any{
		"I_Frames": player.I_Frames,
	})

	if !player.Dir {
		op.GeoM.Translate(640-camera.Cam.Manual_Offset.X, 360-camera.Cam.Manual_Offset.Y)
		player.Img.Draw(screen, &op)
	} else {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(640+32-camera.Cam.Manual_Offset.X, 360-camera.Cam.Manual_Offset.Y)
		player.Img.Draw(screen, &op)
	}

	op.GeoM.Reset()

	for entity_index := 0; entity_index < len(player.Entities); entity_index++ {
		entity := &player.Entities[entity_index]
		op := ebiten.DrawImageOptions{}
		if !entity.Dir {
			op.GeoM.Translate(-(float64(entity.Img.GetTexture().Bounds().Dx()))/2, -(float64(entity.Img.GetTexture().Bounds().Dy()))/2)
			op.GeoM.Rotate(utils.Deg2Rad(entity.Rotation))
			op.GeoM.Translate((float64(entity.Img.GetTexture().Bounds().Dx()))/2, (float64(entity.Img.GetTexture().Bounds().Dy()))/2)
			op.GeoM.Translate(entity.Pos.X-camera.Cam.Offset.X+640-camera.Cam.Manual_Offset.X, entity.Pos.Y-camera.Cam.Offset.Y+360-camera.Cam.Manual_Offset.Y)
		} else {
			op.GeoM.Translate(-(float64(entity.Img.GetTexture().Bounds().Dx()))/2, -(float64(entity.Img.GetTexture().Bounds().Dy()))/2)
			op.GeoM.Rotate(utils.Deg2Rad(entity.Rotation))
			op.GeoM.Translate((float64(entity.Img.GetTexture().Bounds().Dx()))/2, (float64(entity.Img.GetTexture().Bounds().Dy()))/2)
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(entity.Pos.X-camera.Cam.Offset.X-camera.Cam.Manual_Offset.X+640+float64(entity.Img.GetTexture().Bounds().Dx()), entity.Pos.Y-camera.Cam.Offset.Y-camera.Cam.Manual_Offset.Y+360)
		}
		entity.Img.Draw(screen, &op)
	}

	for projectile_index := 0; projectile_index < len(player.Projectiles); projectile_index++ {
		op.GeoM.Reset()
		op.GeoM.Translate(player.Projectiles[projectile_index].Pos.X-camera.Cam.Offset.X-camera.Cam.Manual_Offset.X+650, player.Projectiles[projectile_index].Pos.Y-camera.Cam.Offset.Y-camera.Cam.Manual_Offset.Y+380)
		player.Projectiles[projectile_index].Img.Draw(screen, &op)
	}
}

func (player *Player) Update(level_hitbox []utils.HitBox) {
	player.Img.Update()
	player.Img.RefreshTexture()

	if player.I_Frames > 0 {
	}

	player.Vel.Y += 0.1
	if player.Vel.X != 0 {
		player.Img.Current_Animation = 1
	} else {
		player.Img.Current_Animation = 0
	}

	if player.Vel.X > 5 {
		player.Vel.X -= 0.1
		if player.Vel.X > 10 {
			player.Vel.X -= 0.2
		}
	} else if player.Vel.X < -5 {
		player.Vel.X += 0.1
		if player.Vel.X < -10 {
			player.Vel.X += 0.2
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.Vel.X -= 0.1
		player.Dir = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.Vel.X += 0.1
		player.Dir = false
	} else {
		if player.Vel.X > 0 {
			player.Vel.X -= 0.2
			if player.Vel.X > -0.6 && player.Vel.X < 0.6 {
				player.Vel.X = 0
			}
		} else if player.Vel.X < 0 {
			player.Vel.X += 0.2
			if player.Vel.X > -0.6 && player.Vel.X < 0.6 {
				player.Vel.X = 0
			}
		}
	}

	for button_index := 0; button_index < len(player.Attacks); button_index++ {
		player.Attacks[button_index].Cooldown -= 0.1
		if player.Attacks[button_index].Cooldown < 0 {
			player.Attacks[button_index].Cooldown = 0
		}
		if ebiten.IsKeyPressed(attack_keys[button_index]) && player.Attacks[button_index].Cooldown <= 0 && attack_keys[button_index] != utils.Empty_Key {
			player.Attacks[button_index].Attack()
			player.Attacks[button_index].Cooldown = player.Attacks[button_index].Max_Cooldown
		} else if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) && player.Attacks[1].Cooldown <= 0 {
			player.Attacks[1].Attack()
			player.Attacks[1].Cooldown = player.Attacks[1].Max_Cooldown
		}
	}

	player.Domain.Effect(enemyai.Enemies_In_World)

	if player.Domain_Timer < 0 {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			player.Activate_Domain = true
			player.Domain_Timer = 360
		}
	} else {
		player.Domain_Timer -= 0.1
	}

	if utils.Collide(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y + player.Vel.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		player.Vel.Y = 0
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
			if utils.Collide(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y + player.Vel.Y - 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000, Y: -2000 + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
				player.Vel.Y = 0
			} else {
				player.Vel.Y = -5.1
			}
		}
	}
	if utils.Collide(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y + player.Vel.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) - (250)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		player.Vel.Y = 0
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
			if utils.Collide(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y + player.Vel.Y - 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000, Y: -2000 + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
				player.Vel.Y = 0
			} else {
				player.Vel.Y = -5.1
			}
		}
	}

	if utils.Collide(utils.Vec2{X: player.Pos.X + player.Vel.X, Y: player.Pos.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000 - (1280 / 2), Y: -2000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 2048, Y: (126 * 2)}) {
		player.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: player.Pos.X + player.Vel.X, Y: player.Pos.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		player.Vel.X = 0
	}

	if utils.Collide(utils.Vec2{X: player.Pos.X + player.Vel.X, Y: player.Pos.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: 2000 + 2048 - (1280 / 2), Y: -3000 - (720 / 2) + (449 * 2)}, utils.Vec2{X: 1, Y: 1000}) {
		player.Vel.X = 0
	}

	for tile_index := 0; tile_index < len(level_hitbox); tile_index++ {
		tile := level_hitbox[tile_index]
		if utils.Collide(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y + player.Vel.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: tile.X, Y: tile.Y}, utils.Vec2{X: 32, Y: 32}) {
			player.Vel.Y = 0
			if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
				if utils.Collide(utils.Vec2{X: player.Pos.X, Y: player.Pos.Y + player.Vel.Y}, utils.Vec2{X: 32, Y: 64}, utils.Vec2{X: tile.X, Y: tile.Y}, utils.Vec2{X: 32, Y: 32}) {
					player.Vel.Y = 0
				} else {
					player.Vel.Y = -5.1
				}
			}
		}
		if utils.Collide(utils.Vec2{X: player.Pos.X + player.Vel.X, Y: player.Pos.Y + 2}, utils.Vec2{X: 32, Y: 62}, utils.Vec2{X: tile.X, Y: tile.Y}, utils.Vec2{X: 32, Y: 32}) {
			player.Vel.X = 0
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !utils.Clicked {
		player.Punch()
		utils.Clicked = true
	}

	for projectile_index := 0; projectile_index < len(player.Projectiles); projectile_index++ {
		projectile := &player.Projectiles[projectile_index]
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
						utils.RemoveArrayElement(projectile_index, &player.Projectiles)
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
				utils.RemoveArrayElement(projectile_index, &player.Projectiles)
			}
		}
	}

	for entity_index := 0; entity_index < len(player.Entities); entity_index++ {
		entity := &player.Entities[entity_index]
		entity.Update(entity, level_hitbox)

		if entity.Lifespan < 0 {
			utils.RemoveArrayElement(entity_index, &player.Entities)
		}
	}

	player.Pos.Y += player.Vel.Y
	player.Pos.X += player.Vel.X
}

var Player_Ref Player

func init() {
	Player_Ref = newPlayer(utils.Vec2{X: 0, Y: 0}, *textures.NewAnimatedTexture("./art/players/greg.png", ""), textures.NewTexture("./art/domains/simple_domain.png", shaders.Player_Shader), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, greg_attacks)
}
