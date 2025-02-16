package main

import (
	"encoding/json"
	"image"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type RenderableTexture interface {
	draw(s *ebiten.Image, op *ebiten.DrawImageOptions)
	update()
	getTexture() *ebiten.Image
}

type Texture struct {
	path string
	img  *ebiten.Image
}

func newTexture(img_path string) *Texture {
	t := Texture{}

	t.path = img_path

	timg, _, err := ebitenutil.NewImageFromFile(img_path)
	if err != nil {
		panic(err)
	}
	t.img = timg

	return &t
}

func (t *Texture) draw(s *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.DrawImage(t.img, op)
}

func (t *Texture) getTexture() *ebiten.Image {
	return t.img
}

func (t *Texture) update() {}

type Animation struct {
	frames             []*ebiten.Image
	animation_progress int
	speed              float64
	timer              float64
}

type AnimatedTexture struct {
	path              string
	animations        []Animation
	current_animation int
}

type SpriteSheetData struct {
	Frames [][][]int
	Speed  float64
}

func newAnimatedTexture(path string) *AnimatedTexture {
	sprite_sheet, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}

	at := AnimatedTexture{}

	at.path = path

	temp := SpriteSheetData{}

	json_path := strings.Replace(path, "png", "json", -1)

	temp_data, _ := os.ReadFile(json_path)

	json.Unmarshal(temp_data, &temp)

	animations := []Animation{}

	for anim := 0; anim < len(temp.Frames); anim++ {
		animations = append(animations, Animation{})
		animations[anim].speed = float64(temp.Speed)
		animations[anim].timer = 0
		for fram := 0; fram < len(temp.Frames[anim]); fram++ {
			frame := []float64{float64(int(temp.Frames[anim][fram][0])), float64(int(temp.Frames[anim][fram][1])), float64(int(temp.Frames[anim][fram][2])), float64(int(temp.Frames[anim][fram][3]))}
			animations[anim].frames = append(animations[anim].frames, ebiten.NewImageFromImage(sprite_sheet.SubImage(image.Rect(int(frame[0]), int(frame[1]), int(frame[2]), int(frame[3])))))
		}
	}
	at.animations = animations

	return &at
}

func newFunction(animations []Animation, anim int, frame []float64) {
	animations[anim].speed = frame[4]
}

func (t *AnimatedTexture) draw(s *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.DrawImage(t.animations[t.current_animation].frames[t.animations[t.current_animation].animation_progress], op)
}

func (t *AnimatedTexture) update() {
	t.animations[t.current_animation].timer -= t.animations[t.current_animation].speed

	if t.animations[t.current_animation].timer < 0 {
		t.animations[t.current_animation].animation_progress += 1
		if t.animations[t.current_animation].animation_progress >= len(t.animations[t.current_animation].frames) {
			t.animations[t.current_animation].animation_progress = 0
		}
		t.animations[t.current_animation].timer = 1
	}
}

func (t *AnimatedTexture) getTexture() *ebiten.Image {
	return t.animations[t.current_animation].frames[t.animations[t.current_animation].animation_progress]
}
