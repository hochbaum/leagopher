package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/lhochbaum/leagopher/gfx"
	"github.com/lhochbaum/leagopher/system"
	"github.com/lhochbaum/leagopher/ui"
)

const (
	windowName = "Leagopher"
)

func main() {
	// initialize window.
	rl.InitWindow(system.WindowWidth, system.WindowHeight, windowName)

	// initialize audio device for background music.
	rl.InitAudioDevice()

	// set up some background music. todo: use scenes for this instead.
	music := rl.LoadMusicStream("assets/sfx/main_menu.ogg")
	rl.PlayMusicStream(music)
	rl.SetMusicVolume(music, .1)

	defer func() {
		rl.UnloadMusicStream(music)
		rl.CloseAudioDevice()
	}()

	// initialize stage and main menu.
	stage := ui.NewStage()
	stage.PutScene(ui.NewMainMenu(stage))
	stage.PutScene(ui.NewTestMenu())
	stage.SetScene("MainMenu")

	// create the renderer and let it draw the UI.
	rend := gfx.NewRenderer()
	rend.Add(stage.Draw)

	// main game loop.
	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(music)

		// begin drawing.
		rl.BeginDrawing()

		// clear the background.
		rl.ClearBackground(rl.Black)

		// render the layers.
		rend.Render()

		// end drawing.
		rl.EndDrawing()
	}
}
