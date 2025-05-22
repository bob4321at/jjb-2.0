package level

import (
	"fmt"
	"jjb/camera"
	"jjb/enemyai"
	"jjb/players"
	"jjb/utils"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var temp_tile_img, _, _ = ebitenutil.NewImageFromFile("./art/temp_tile.png")

type Tile struct {
	Pos  utils.Vec2
	Tile int
}

type Waves struct {
	Waves [][]int
}

type Level struct {
	Tile_map             [][]uint8
	Tiles                []Tile
	Level_Cached_Image   *ebiten.Image
	Player_Spawn         utils.Vec2
	Spawn_Points         []utils.Vec2
	Enemies              []enemyai.Enemy
	Background           Background
	Gnerated             bool
	Tileset              map[int]*ebiten.Image
	Waves                Waves
	Current_Wave         int
	Spawn_Timer          float64
	Origonal_Spawn_Timer float64
	Spawned              bool
	HitBox               []utils.HitBox
	Beaten               bool
}

var Levels = []Level{}

var Current_Level *Level
var Current_Level_Index int

var test *ebiten.Image

func (level *Level) GenCach() {
	if level.Level_Cached_Image == nil {
		level.Level_Cached_Image = ebiten.NewImage(128*32, 128*32)
		for _, tile := range level.Tiles {
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(tile.Pos.X, tile.Pos.Y)
			level.Level_Cached_Image.DrawImage(level.Tileset[tile.Tile], &op)
		}
		fmt.Println("working")
	}
}

func (level *Level) Draw(screen *ebiten.Image, cam *camera.Camera) {
	level.GenCach()

	op := ebiten.DrawImageOptions{}
	level.Background.Draw(screen, cam)

	op.GeoM.Reset()
	op.GeoM.Translate(-cam.Offset.X-cam.Manual_Offset.X+640, -cam.Offset.Y-cam.Manual_Offset.Y+360)
	screen.DrawImage(level.Level_Cached_Image, &op)

	op.GeoM.Reset()

	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(1000-camera.Cam.Offset.X-camera.Cam.Manual_Offset.X, -2500-camera.Cam.Offset.Y-camera.Cam.Manual_Offset.Y)
	screen.DrawImage(utils.Domain_Background, &op)

	op.GeoM.Reset()

	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(2000-camera.Cam.Offset.X-camera.Cam.Manual_Offset.X, -2000-camera.Cam.Offset.Y-camera.Cam.Manual_Offset.Y)

	players.Player_Ref.Domain.Img.Draw(screen, &op)

	for enemy_index := 0; enemy_index < len(level.Enemies); enemy_index++ {
		level.Enemies[enemy_index].Draw(screen, cam)
	}
}

func (level *Level) SpawnWave() {
	for enemy_index := 0; enemy_index < len(level.Waves.Waves[level.Current_Wave]); enemy_index += 1 - 1 {
		time.Sleep(20000)
		if level.Spawn_Timer < 0 {
			level.Spawn_Timer = level.Origonal_Spawn_Timer
			level.Spawn(enemyai.Enemy_Table[level.Waves.Waves[level.Current_Wave][enemy_index]])
			enemy_index += 1
		} else {
			level.Spawn_Timer -= 0.01
		}
	}

	for level.Spawned {
		if len(level.Enemies) == 0 {
			level.Spawned = false
			level.Current_Wave += 1
		}
	}
}

func (level *Level) Update(player *players.Player) {
	player.DamageCheck()
	players.Emergency_Level_Hitbox = level.HitBox

	player.Update(level.HitBox)

	enemyai.Enemies_In_World = []*enemyai.Enemy{}
	for enemy_index := 0; enemy_index < len(level.Enemies); enemy_index++ {
		enemy := &level.Enemies[enemy_index]
		enemyai.Enemies_In_World = append(enemyai.Enemies_In_World, enemy)

		enemy.Update(&level.Enemies[enemy_index], players.Player_Ref.Pos, level.HitBox)

		level.Enemies[enemy_index].Tex.Update()

		for projectile_index := 0; projectile_index < len(enemy.Projectiles); projectile_index++ {
			projectile := &enemy.Projectiles[projectile_index]
			projectile.Pos.X += projectile.Vel.X
			projectile.Pos.Y += projectile.Vel.Y
		}

		if enemy.Health <= 0 {
			enemy.Alive = false
			enemy_index := 0
			for i := 0; i < len(level.Enemies); i++ {
				if &level.Enemies[i] == enemy {
					enemy_index = i
					i = len(level.Enemies) + 1
				}
			}
			utils.RemoveArrayElement(enemy_index, &level.Enemies)
		}
	}

	for enemy_index := 0; enemy_index < len(enemyai.Enemies_To_Add); enemy_index++ {
		level.Spawn(enemyai.Enemies_To_Add[enemy_index])
		utils.RemoveArrayElement(enemy_index, &enemyai.Enemies_To_Add)
	}
	if !level.Spawned && level.Current_Wave < len(level.Waves.Waves) {
		go level.SpawnWave()
		level.Spawned = true
	}
	if level.Current_Wave >= len(level.Waves.Waves) && Current_Level_Index+1 < len(Levels) {
		go func() {
			time.Sleep(time.Second * 2)
			level.Beaten = true
		}()
	}
}

func (level *Level) Spawn(enemy enemyai.Enemy) {
	for spawn_point_index := 0; spawn_point_index < len(level.Spawn_Points); spawn_point_index++ {
		enemy.Pos = level.Spawn_Points[spawn_point_index]
		enemy.Pos.X += rand.Float64()
		enemy.Pos.Y += rand.Float64()
		level.Enemies = append(level.Enemies, enemy)
	}
}
