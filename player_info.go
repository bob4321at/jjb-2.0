package main

var players map[string]Player

func initPlayer() {
	players = map[string]Player{
		"greg":   newPlayer(current_level.player_spawn, newAnimatedTexture("./art/players/greg.png"), greg_attacks),
		"gojo":   newPlayer(current_level.player_spawn, newAnimatedTexture("./art/players/gojo.png"), gojo_attacks),
		"megumi": newPlayer(current_level.player_spawn, newAnimatedTexture("./art/players/megumi.png"), megumi_attacks),
	}
}
