package main

import (
	"jjb/players"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var healthbar_img, _, _ = ebitenutil.NewImageFromFile("./art/ui/healthbar.png")

var keybinds, _, _ = ebitenutil.NewImageFromFile("./art/ui/keybinds.png")

var keybind_cover, _, _ = ebitenutil.NewImageFromFile("./art/ui/keybing_cover.png")

func drawUi(s *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, 10)
	op.GeoM.Scale((float64(players.Player_Ref.Health) / 100), 1)
	s.DrawImage(healthbar_img, &op)

	op.GeoM.Reset()

	op.GeoM.Translate(15, 45)
	s.DrawImage(keybinds, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (players.Player_Ref.Attacks[0].Cooldown / players.Player_Ref.Attacks[0].Max_Cooldown))
	op.GeoM.Translate(16, 46)
	s.DrawImage(keybind_cover, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (players.Player_Ref.Attacks[1].Cooldown / players.Player_Ref.Attacks[1].Max_Cooldown))
	op.GeoM.Translate(48, 46)
	s.DrawImage(keybind_cover, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (players.Player_Ref.Attacks[2].Cooldown / players.Player_Ref.Attacks[2].Max_Cooldown))
	op.GeoM.Translate(16, 78)
	s.DrawImage(keybind_cover, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (players.Player_Ref.Domain_Timer / 360))
	op.GeoM.Translate(16, 110)
	s.DrawImage(keybind_cover, &op)
}
