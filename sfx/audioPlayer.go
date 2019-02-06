package sfx

import (
	"github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

var (
	tracks = make(map[string]*raylib.Music)
	sounds = make(map[string]*raylib.Sound)
)

func PlayMusic(name string) {
	music, ok := tracks[name]
	if !ok {
		path := fmt.Sprintf("assets/sfx/%s.ogg", name)
		stream := raylib.LoadMusicStream(path)

		music = &stream
		tracks[name] = music
	}

	raylib.PlayMusicStream(*music)
}

func PlaySound(name string) {
	sound, ok := sounds[name]
	if !ok {
		path := fmt.Sprintf("assets/sfx/%s.ogg", name)
		s := raylib.LoadSound(path)

		sound = &s
		sounds[name] = sound
	}

	raylib.PlaySound(*sound)
}