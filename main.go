package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/lhochbaum/leagopher/gfx"
	"github.com/lhochbaum/leagopher/ui"
	"github.com/lhochbaum/leagopher/system"
)

const (
	windowName = "Leagopher"
)

func main() {
	// initialize window.
	raylib.InitWindow(system.WindowWidth, system.WindowHeight, windowName)

	// initialize audio device for background music.
	raylib.InitAudioDevice()

	// set up some background music. todo: use scenes for this instead.
	music := raylib.LoadMusicStream("assets/sfx/main_menu.ogg")
	raylib.PlayMusicStream(music)
	raylib.SetMusicVolume(music, .1)

	defer func() {
		raylib.UnloadMusicStream(music)
		raylib.CloseAudioDevice()
	}()

	// initialize stage and main menu.
	stage := ui.NewStage()
	stage.PutScene(ui.NewMainMenu())
	stage.SetScene("MainMenu")

	// create the renderer and let it draw the UI.
	rend := gfx.NewRenderer()
	rend.Add(stage.Draw)

	// main game loop.
	for !raylib.WindowShouldClose() {
		raylib.UpdateMusicStream(music)

		// begin drawing.
		raylib.BeginDrawing()

		// clear the background.
		raylib.ClearBackground(raylib.Black)

		// render the layers.
		rend.Render()

		// end drawing.
		raylib.EndDrawing()
	}
}