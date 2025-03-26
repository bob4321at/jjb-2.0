package players

import (
	"jjb/enemyai"
	"jjb/shaders"
	"jjb/textures"
	"jjb/utils"
)

var Players map[string]Player

func InitPlayer(spawn_point utils.Vec2) {
	Players = map[string]Player{
		"greg":        newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/greg.png", shaders.Player_Shader), textures.NewTexture("./art/domains/simple_domain.png", ""), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, greg_attacks),
		"gojo":        newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/gojo.png", shaders.Player_Shader), textures.NewTexture("./art/domains/gojo_domain.png", ""), func(enemies []*enemyai.Enemy) { Player_Ref.gojoDomain(enemyai.Enemies_In_World) }, gojo_attacks),
		"megumi":      newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/megumi.png", shaders.Player_Shader), textures.NewTexture("./art/domains/megumi_domain.png", ""), func(enemies []*enemyai.Enemy) { Player_Ref.megumiDomain(enemyai.Enemies_In_World) }, megumi_attacks),
		"boberto":     newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/boberto.png", shaders.Player_Shader), textures.NewTexture("./art/domains/boberto_domain.png", ""), func(enemies []*enemyai.Enemy) { Player_Ref.bobertoDomain(enemyai.Enemies_In_World) }, boberto_attacks),
		"jerry":       newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/jerry.png", shaders.Player_Shader), textures.NewTexture("./art/domains/simple_domain.png", ""), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, jerry_attacks),
		"sukuna":      newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/sukuna_playable.png", shaders.Player_Shader), textures.NewTexture("./art/domains/sukuna_domain.png", ""), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, sukuna_attacks),
		"hermes":      newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/hermes.png", shaders.Player_Shader), textures.NewTexture("./art/domains/simple_domain.png", ""), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, hermes_attacks),
		"test_player": newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/test_player.png", shaders.Player_Shader), textures.NewTexture("./art/domains/simple_domain.png", ""), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, test_player_attacks),
		"agent_21":    newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/agent_21.png", shaders.Player_Shader), textures.NewTexture("./art/domains/simple_domain.png", ""), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, agent_21_attacks),

		"tk": newPlayer(spawn_point, *textures.NewAnimatedTexture("./art/players/tk.png", shaders.Player_Shader), textures.NewTexture("./art/domains/simple_domain.png", ""), func(enemies []*enemyai.Enemy) { Player_Ref.simpleDomain(enemyai.Enemies_In_World) }, tk_attacks),
	}
}
