package main

var players map[string]Player

func initPlayer() {
	players = map[string]Player{
		"greg":    newPlayer(current_level.player_spawn, *newAnimatedTexture("./art/players/greg.png"), newTexture("./art/domains/simple_domain.png"), func(l *Level) { player.simpleDomain(current_level) }, greg_attacks),
		"gojo":    newPlayer(current_level.player_spawn, *newAnimatedTexture("./art/players/gojo.png"), newTexture("./art/domains/gojo_domain.png"), func(l *Level) { player.gojoDomain(current_level) }, gojo_attacks),
		"megumi":  newPlayer(current_level.player_spawn, *newAnimatedTexture("./art/players/megumi.png"), newTexture("./art/domains/megumi_domain.png"), func(l *Level) { player.megumiDomain(current_level) }, megumi_attacks),
		"boberto": newPlayer(current_level.player_spawn, *newAnimatedTexture("./art/players/boberto.png"), newTexture("./art/domains/boberto_domain.png"), func(l *Level) { player.bobertoDomain(current_level) }, boberto_attacks),
	}
}
