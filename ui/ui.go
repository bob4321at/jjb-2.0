package ui

import (
	"image/color"
	"jjb/level"
	"jjb/players"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var healthbar_img, _, _ = ebitenutil.NewImageFromFile("./art/ui/healthbar.png")

var keybinds, _, _ = ebitenutil.NewImageFromFile("./art/ui/keybinds.png")

var keybind_cover, _, _ = ebitenutil.NewImageFromFile("./art/ui/keybing_cover.png")

var wave_progressbar, _, _ = ebitenutil.NewImageFromFile("./art/ui/wave_progressbar.png")
var wave_progress_indicator, _, _ = ebitenutil.NewImageFromFile("./art/ui/wave_progress_indicator.png")
var yellaw_bar = ebiten.NewImage(1, 1)

func init() {
	yellaw_bar.Set(0, 0, color.RGBA{255, 255, 0, 255})
}

func DrawUi(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, 10)
	op.GeoM.Scale((float64(players.Player_Ref.Health) / 100), 1)
	screen.DrawImage(healthbar_img, &op)

	op.GeoM.Reset()

	op.GeoM.Translate(15, 45)
	screen.DrawImage(keybinds, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (players.Player_Ref.Attacks[0].Cooldown / players.Player_Ref.Attacks[0].Max_Cooldown))
	op.GeoM.Translate(16, 46)
	screen.DrawImage(keybind_cover, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (players.Player_Ref.Attacks[1].Cooldown / players.Player_Ref.Attacks[1].Max_Cooldown))
	op.GeoM.Translate(48, 46)
	screen.DrawImage(keybind_cover, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (players.Player_Ref.Attacks[2].Cooldown / players.Player_Ref.Attacks[2].Max_Cooldown))
	op.GeoM.Translate(16, 78)
	screen.DrawImage(keybind_cover, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(1, (players.Player_Ref.Domain_Timer / 360))
	op.GeoM.Translate(16, 110)
	screen.DrawImage(keybind_cover, &op)

	op.GeoM.Reset()
	op.GeoM.Translate(1280-524, 12)
	screen.DrawImage(wave_progressbar, &op)

	op.GeoM.Reset()
	op.GeoM.Scale(float64(level.Current_Level.Current_Wave)*float64(512/len(level.Current_Level.Waves.Waves)), 44)
	op.GeoM.Translate(1280-522, 14)
	screen.DrawImage(yellaw_bar, &op)

	for i := 1; i < len(level.Current_Level.Waves.Waves); i++ {
		op.GeoM.Reset()
		op.GeoM.Translate(1280-524+float64(i*(512/len(level.Current_Level.Waves.Waves))), 12)
		screen.DrawImage(wave_progress_indicator, &op)
	}
}
