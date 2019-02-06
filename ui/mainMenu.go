package ui

import (
	"github.com/lhochbaum/leagopher/math"
	"github.com/lhochbaum/leagopher/system"
)

func NewMainMenu() *Scene {
	menu := NewScene("MainMenu", "assets/gfx/back_main_menu.png", nil)

	menu.AddOption(&Option{
		Text: "PLAY",
		TextSize: 30,
		Position: *math.NewVector2(system.WindowWidth / 2 - 50, system.WindowHeight / 2 - 100),
	}).AddOption(&Option{
		Text: "SETTINGS",
		TextSize: 30,
		Position: *math.NewVector2(system.WindowWidth / 2 - 95, system.WindowHeight / 2 - 50),
	}).AddOption(&Option{
		Text: "REPORT A BUG",
		TextSize: 30,
		Position: *math.NewVector2(system.WindowWidth / 2 - 130, system.WindowHeight / 2),
	})

	return menu
}