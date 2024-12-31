package main

var players map[string]Player

func initPlayer() {
	players = map[string]Player{
		"temp": newPlayer(current_level.player_spawn, "./art/temp_player.png", megumi_attacks),
		"greg": newPlayer(current_level.player_spawn, "./art/players/greg.png", greg_attacks),
		"gojo": newPlayer(current_level.player_spawn, "./art/players/gojo.png", gojo_attacks),
	}
}
