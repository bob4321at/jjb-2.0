package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func drawUi(s *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, 10)
	op.GeoM.Scale((float64(player.health) / 100), 1)
	s.DrawImage(healthbar_img, &op)

	op.GeoM.Reset()

	op.GeoM.Translate(15, 45)
	s.DrawImage(keybinds, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (player.attacks[0].cooldown / player.attacks[0].max_cooldown))
	op.GeoM.Translate(16, 46)
	s.DrawImage(keybind_cover, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (player.attacks[1].cooldown / player.attacks[1].max_cooldown))
	op.GeoM.Translate(48, 46)
	s.DrawImage(keybind_cover, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (player.attacks[2].cooldown / player.attacks[2].max_cooldown))
	op.GeoM.Translate(16, 78)
	s.DrawImage(keybind_cover, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (player.domain_timer / 240))
	op.GeoM.Translate(16, 110)
	s.DrawImage(keybind_cover, &op)
}
