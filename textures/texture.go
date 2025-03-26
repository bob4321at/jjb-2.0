package textures

import (
	"encoding/json"
	"image"
	"jjb/utils"
	"math"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type RenderableTexture interface {
	Draw(s *ebiten.Image, op *ebiten.DrawImageOptions)
	Update()
	GetTexture() *ebiten.Image
	RefreshTexture()
}

type Texture struct {
	Path     string
	Img      *ebiten.Image
	Shader   *ebiten.Shader
	Uniforms map[string]any
}

func NewTexture(img_path string, shader string) *Texture {
	t := Texture{}

	t.Path = img_path

	timg, _, err := ebitenutil.NewImageFromFile(img_path)
	if err != nil {
		panic(err)
	}
	t.Img = timg

	// if shader == "" {
	shader = `//kage:unit pixels
			package main

			var U float

			func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
				col := imageSrc0At(srcPos.xy)
				return vec4(col.x, col.y + U, col.z, col.w)
			}
		`
	// }

	t.Shader, err = ebiten.NewShader([]byte(shader))
	if err != nil {
		panic(err)
	}

	return &t
}

func (t *Texture) Draw(s *ebiten.Image, op *ebiten.DrawImageOptions) {
	opts := &ebiten.DrawRectShaderOptions{}
	opts.Images[0] = t.Img
	opts.Uniforms = map[string]any{
		"U": math.Cos(utils.Game_Time / 40),
	}
	opts.GeoM = op.GeoM
	s.DrawRectShader(t.Img.Bounds().Dx(), t.Img.Bounds().Dy(), t.Shader, opts)
	// s.DrawImage(t.Img, op)
}

func (t *Texture) GetTexture() *ebiten.Image {
	return t.Img
}

func (t *Texture) Update() {}

func (t *Texture) RefreshTexture() {
	t.Path = t.Path

	timg, _, err := ebitenutil.NewImageFromFile(t.Path)
	if err != nil {
		panic(err)
	}
	t.Img = timg
}

type Animation struct {
	Frames             []*ebiten.Image
	Animation_Progress int
	Speed              float64
	Timer              float64
}

type SpriteSheetData struct {
	Frames [][][]int
	Speed  float64
}

type AnimatedTexture struct {
	Path              string
	Animations        []Animation
	Modified          bool
	Current_Animation int
}

func NewAnimatedTexture(path string) *AnimatedTexture {
	sprite_sheet, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}

	at := AnimatedTexture{}

	at.Path = path

	temp := SpriteSheetData{}

	json_path := strings.Replace(path, "png", "json", -1)

	temp_data, _ := os.ReadFile(json_path)

	json.Unmarshal(temp_data, &temp)

	animations := []Animation{}

	for anim := 0; anim < len(temp.Frames); anim++ {
		animations = append(animations, Animation{})
		animations[anim].Speed = float64(temp.Speed)
		animations[anim].Timer = 0
		for fram := 0; fram < len(temp.Frames[anim]); fram++ {
			frame := []float64{float64(int(temp.Frames[anim][fram][0])), float64(int(temp.Frames[anim][fram][1])), float64(int(temp.Frames[anim][fram][2])), float64(int(temp.Frames[anim][fram][3]))}
			animations[anim].Frames = append(animations[anim].Frames, ebiten.NewImageFromImage(sprite_sheet.SubImage(image.Rect(int(frame[0]), int(frame[1]), int(frame[2]), int(frame[3])))))
		}
	}
	at.Animations = animations

	return &at
}

func (t *AnimatedTexture) Draw(s *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.DrawImage(t.Animations[t.Current_Animation].Frames[t.Animations[t.Current_Animation].Animation_Progress], op)
}

func (t *AnimatedTexture) RefreshTexture() {
	sprite_sheet, _, err := ebitenutil.NewImageFromFile(t.Path)
	if err != nil {
		panic(err)
	}

	temp := SpriteSheetData{}

	json_path := strings.Replace(t.Path, "png", "json", -1)

	temp_data, _ := os.ReadFile(json_path)

	json.Unmarshal(temp_data, &temp)

	animations := []Animation{}

	for anim := 0; anim < len(temp.Frames); anim++ {
		animations = append(animations, Animation{})
		animations[anim].Speed = float64(temp.Speed)
		animations[anim].Timer = t.Animations[anim].Timer
		animations[anim].Animation_Progress = t.Animations[anim].Animation_Progress
		for fram := 0; fram < len(temp.Frames[anim]); fram++ {
			frame := []float64{float64(int(temp.Frames[anim][fram][0])), float64(int(temp.Frames[anim][fram][1])), float64(int(temp.Frames[anim][fram][2])), float64(int(temp.Frames[anim][fram][3]))}
			animations[anim].Frames = append(animations[anim].Frames, ebiten.NewImageFromImage(sprite_sheet.SubImage(image.Rect(int(frame[0]), int(frame[1]), int(frame[2]), int(frame[3])))))
		}
	}
	t.Animations = animations
}

func (t *AnimatedTexture) Update() {
	t.Animations[t.Current_Animation].Timer -= t.Animations[t.Current_Animation].Speed

	if t.Animations[t.Current_Animation].Timer < 0 {
		t.Animations[t.Current_Animation].Animation_Progress += 1
		if t.Animations[t.Current_Animation].Animation_Progress >= len(t.Animations[t.Current_Animation].Frames) {
			t.Animations[t.Current_Animation].Animation_Progress = 0
		}
		t.Animations[t.Current_Animation].Timer = 1
	}
}

func (t *AnimatedTexture) GetTexture() *ebiten.Image {
	return t.Animations[t.Current_Animation].Frames[t.Animations[t.Current_Animation].Animation_Progress]
}
