package ui

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/lhochbaum/leagopher/math"
	"github.com/lhochbaum/leagopher/system"
	"time"
	"github.com/lhochbaum/leagopher/gfx"
	"github.com/lhochbaum/leagopher/sfx"
)

type Stage struct {
	Scene  *Scene
	Scenes map[string]*Scene
}

func NewStage() *Stage {
	return &Stage{
		Scene:  nil,
		Scenes: make(map[string]*Scene),
	}
}

func (s *Stage) Draw() {
	for _, scene := range s.Scenes {
		scene.Update(raylib.GetTime())
	}

	if s.Scene == nil || s.Scene.Background == nil {
		return
	}

	raylib.DrawTextureV(*s.Scene.Background, raylib.NewVector2(0, 0), raylib.White)

	for i, option := range s.Scene.Options {
		var color raylib.Color
		if s.Scene.Selected == i {
			color = raylib.Yellow
		} else {
			color = raylib.White
		}

		raylib.DrawText(option.Text, option.Position.X, option.Position.Y, option.TextSize, color)
	}
}

func (s *Stage) SetScene(scene string) {
	s.Scene = s.Scenes[scene]
}

func (s *Stage) PutScene(scene *Scene) {
	s.Scenes[scene.ID] = scene
}

type Scene struct {
	ID         string
	Background *raylib.Texture2D
	Music      *raylib.Music
	Options    []*Option
	Selected int
}

func NewScene(id string, background string, music *raylib.Music) *Scene {
	return &Scene{
		ID:         id,
		Background: gfx.GetTexture(background),
		Music:      music,
	}
}

func (s *Scene) AddOption(opt *Option) *Scene {
	s.Options = append(s.Options, opt)
	return s
}

func (s *Scene) SelectOption(index int) {
	s.Selected = index
}

// todo: redesign!
var optionCooldown = time.Now()

func (s *Scene) Update(dt float32) {
	if time.Since(optionCooldown) <= 200 * time.Millisecond {
		return
	}

	if raylib.IsKeyDown(system.KeyBinds["down"]) {
		if s.Selected == len(s.Options) - 1 {
			s.Selected = 0
		} else {
			s.Selected += 1
		}

		optionCooldown = time.Now()
		sfx.PlaySound("ui_tap")
	} else if raylib.IsKeyDown(system.KeyBinds["up"]) {
		if s.Selected == 0 {
			s.Selected = len(s.Options) - 1
		} else {
			s.Selected -= 1
		}

		optionCooldown = time.Now()
		sfx.PlaySound("ui_tap")
	}
}

type Option struct {
	Text     string
	TextSize int32
	Position math.Vector2
	Callback func()
}
