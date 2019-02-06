package ui

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/lhochbaum/leagopher/math"
	"github.com/lhochbaum/leagopher/system"
)

const fontSize = 30

var stage *Stage

func NewMainMenu(s *Stage) *Scene {
	menu := NewScene("MainMenu", "assets/gfx/back_main_menu.png", nil)
	stage = s

	menu.AddOption(&Option{
		Text:     "PLAY",
		TextSize: fontSize,
		Position: *math.NewVector2(offset("PLAY"), system.WindowHeight/2-100),
		Callback: onPlay,
	}).AddOption(&Option{
		Text:     "SETTINGS",
		TextSize: fontSize,
		Position: *math.NewVector2(offset("SETTINGS"), system.WindowHeight/2-50),
	}).AddOption(&Option{
		Text:     "REPORT A BUG",
		TextSize: fontSize,
		Position: *math.NewVector2(offset("REPORT A BUG"), system.WindowHeight/2),
	})

	return menu
}

func offset(text string) int32 {
	textWidth := rl.MeasureText(text, fontSize)
	return (system.WindowWidth+textWidth)/2 - textWidth
}

func onPlay() {
	stage.SetScene("TestMenu")
}
