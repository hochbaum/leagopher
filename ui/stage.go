package ui

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/lhochbaum/leagopher/gfx"
	"github.com/lhochbaum/leagopher/math"
	"github.com/lhochbaum/leagopher/sfx"
	"github.com/lhochbaum/leagopher/system"
	"time"
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
	if s.Scene == nil {
		return
	}

	s.Scene.Update(rl.GetFrameTime())

	if s.Scene.Background != nil {
		rl.DrawTextureV(*s.Scene.Background, rl.NewVector2(0, 0), rl.White)
	}

	for i, option := range s.Scene.Options {
		var color rl.Color
		if s.Scene.Selected == i {
			color = rl.Yellow
		} else {
			color = rl.White
		}

		rl.DrawText(option.Text, option.Position.X, option.Position.Y, option.TextSize, color)
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
	Background *rl.Texture2D
	Music      *rl.Music
	Options    []*Option
	Selected   int
}

func NewScene(id string, backPath string, music *rl.Music) *Scene {
	var background *rl.Texture2D
	if backPath == "" {
		background = nil
	} else {
		background = gfx.GetTexture(backPath)
	}

	return &Scene{
		ID:         id,
		Background: background,
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
	if time.Since(optionCooldown) <= 200*time.Millisecond {
		return
	}

	if len(s.Options) > 0 {
		if rl.IsKeyDown(system.KeyBinds["enter"]) {
			callback := s.Options[s.Selected].Callback
			if callback != nil {
				callback()
			}

			optionCooldown = time.Now()
			sfx.PlaySound("ui_tap")
			return
		}

		if rl.IsKeyDown(system.KeyBinds["down"]) {
			if s.Selected == len(s.Options)-1 {
				s.Selected = 0
			} else {
				s.Selected += 1
			}

			optionCooldown = time.Now()
			sfx.PlaySound("ui_tap")
		} else if rl.IsKeyDown(system.KeyBinds["up"]) {
			if s.Selected == 0 {
				s.Selected = len(s.Options) - 1
			} else {
				s.Selected -= 1
			}

			optionCooldown = time.Now()
			sfx.PlaySound("ui_tap")
		}
	}
}

type Option struct {
	Text     string
	TextSize int32
	Position math.Vector2
	Callback func()
}
