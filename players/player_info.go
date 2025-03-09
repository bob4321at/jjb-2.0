package players

import (
	"jjb/enemyai"
	"jjb/textures"
	"jjb/utils"
)

var Players map[string]Player

func InitPlayer(spawn_point utils.Vec2) {
	Players = map[string]Player{
		"greg":    newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/greg.png"), textures.NewTexture("./art/domains/simple_domain.png"), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, greg_attacks),
		"gojo":    newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/gojo.png"), textures.NewTexture("./art/domains/gojo_domain.png"), func(enemies []*enemyai.Enemy) { Player_Ref.gojoDomain(enemyai.Enemies_In_World) }, gojo_attacks),
		"megumi":  newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/megumi.png"), textures.NewTexture("./art/domains/megumi_domain.png"), func(enemies []*enemyai.Enemy) { Player_Ref.megumiDomain(enemyai.Enemies_In_World) }, megumi_attacks),
		"boberto": newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/boberto.png"), textures.NewTexture("./art/domains/boberto_domain.png"), func(enemies []*enemyai.Enemy) { Player_Ref.bobertoDomain(enemyai.Enemies_In_World) }, boberto_attacks),
		"jerry":   newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/jerry.png"), textures.NewTexture("./art/domains/simple_domain.png"), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, jerry_attacks),
	}
}
