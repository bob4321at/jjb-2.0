package enemyai

import (
	"jjb/camera"
	"jjb/textures"
	"jjb/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

var Enemies_In_World = []*Enemy{}
var Enemies_To_Add = []Enemy{}

type EnemyProjectile struct {
	Pos      utils.Vec2
	Vel      utils.Vec2
	Img      textures.RenderableTexture
	Damage   int
	Lifetime float64
}

func (e *Enemy) NewProjectile(pos, vel utils.Vec2, img textures.RenderableTexture, damage int, lifetime float64) {
	projectile := EnemyProjectile{}

	projectile.Pos = pos
	projectile.Vel = vel

	projectile.Img = img

	projectile.Damage = damage
	projectile.Lifetime = lifetime

	e.Projectiles = append(e.Projectiles, projectile)
}

type Enemy struct {
	Id          int
	Health      int
	MaxHealth   int
	Damage      int
	Projectiles []EnemyProjectile
	Alive       bool
	Pos         utils.Vec2
	Vel         utils.Vec2
	Can_Move    bool
	Tex         textures.RenderableTexture
	Dir         bool
	Update      func(e *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox)
	Coll_Shape  utils.HitBox
}

func NewEnemy(id int, health int, damage int, pos utils.Vec2, img textures.RenderableTexture, update func(enemy *Enemy, player_pos utils.Vec2, level_hitbox []utils.HitBox)) (enemy Enemy) {
	enemy.Id = id
	enemy.Pos = pos
	enemy.Vel = utils.Vec2{}
	enemy.Can_Move = true
	enemy.Health = health
	enemy.MaxHealth = health
	enemy.Alive = true

	enemy.Tex = img
	enemy.Damage = damage

	enemy.Update = update

	enemy.Coll_Shape = utils.HitBox{X: pos.X, Y: pos.Y, W: float64(enemy.Tex.GetTexture().Bounds().Dx()), H: float64(enemy.Tex.GetTexture().Bounds().Dy())}

	return enemy
}

func (enemy *Enemy) Draw(screen *ebiten.Image, cam *camera.Camera) {
	if enemy.Pos.X-cam.Offset.X+640+float64(enemy.Tex.GetTexture().Bounds().Dx()) > 0 && enemy.Pos.X-cam.Offset.X+640-float64(enemy.Tex.GetTexture().Bounds().Dx()) < 1280 {
		if enemy.Pos.Y-cam.Offset.Y+360+float64(enemy.Tex.GetTexture().Bounds().Dy()) > 0 && enemy.Pos.Y-cam.Offset.Y+360-float64(enemy.Tex.GetTexture().Bounds().Dy()) < 720 {
			op := ebiten.DrawImageOptions{}

			if !enemy.Dir {
				op.GeoM.Translate(enemy.Pos.X-cam.Offset.X+640, enemy.Pos.Y-cam.Offset.Y+360)
			} else {
				op.GeoM.Scale(-1, 1)
				op.GeoM.Translate(enemy.Pos.X-cam.Offset.X+640+float64(enemy.Tex.GetTexture().Bounds().Dx()), enemy.Pos.Y-cam.Offset.Y+360)
			}

			enemy.Tex.Draw(screen, &op)
		}
	}

	for projectile_index := 0; projectile_index < len(enemy.Projectiles); projectile_index++ {
		projectile := &enemy.Projectiles[projectile_index]

		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(projectile.Pos.X-cam.Offset.X+640, projectile.Pos.Y-cam.Offset.Y+360)

		projectile.Img.Draw(screen, &op)
	}
}

func (enemy *Enemy) CheckRemove() {
}

var Enemy_Table = map[int]Enemy{}

func init() {
	Enemy_Table = map[int]Enemy{
		1:  NewEnemy(1, 10, 1, utils.Vec2{}, textures.NewAnimatedTexture("./art/enemies/fliehead.png"), flieHeadUpdate),
		2:  NewEnemy(2, 20, 2, utils.Vec2{}, textures.NewTexture("./art/enemies/crooked.png"), crookedUpdate),
		3:  NewEnemy(3, 10, 2, utils.Vec2{}, textures.NewTexture("./art/enemies/shrimp.png"), shrimpUpdate),
		4:  NewEnemy(4, 100, 5, utils.Vec2{}, textures.NewAnimatedTexture("./art/enemies/bosshead.png"), bossHeadUpdate),
		5:  NewEnemy(5, 10, 5, utils.Vec2{}, textures.NewTexture("./art/enemies/cloudhead.png"), cloudHeadUpdate),
		6:  NewEnemy(6, 20, 3, utils.Vec2{}, textures.NewTexture("./art/enemies/balloon.png"), balloonUpdate),
		7:  NewEnemy(7, 10, 3, utils.Vec2{}, textures.NewAnimatedTexture("./art/enemies/bunny.png"), bunnyUpdate),
		8:  NewEnemy(8, 20, 3, utils.Vec2{}, textures.NewAnimatedTexture("./art/enemies/fuzzball.png"), fuzzBallUpdate),
		9:  NewEnemy(9, 200, 5, utils.Vec2{}, textures.NewTexture("./art/enemies/simple_grade_curse.png"), simpleGradeCurseUpdate),
		10: NewEnemy(9, 200, 10, utils.Vec2{}, textures.NewTexture("./art/enemies/sukuna.png"), sukunaUpdate),
		11: NewEnemy(6, 50, 3, utils.Vec2{}, textures.NewTexture("./art/enemies/green_balloon.png"), greenBalloonUpdate),
		12: NewEnemy(6, 100, 3, utils.Vec2{}, textures.NewTexture("./art/enemies/purple_balloon.png"), purpleBalloonUpdate),
	}
}
