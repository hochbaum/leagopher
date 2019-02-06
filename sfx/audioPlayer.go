package sfx

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	tracks = make(map[string]*rl.Music)
	sounds = make(map[string]*rl.Sound)
)

func PlayMusic(name string) {
	music, ok := tracks[name]
	if !ok {
		path := fmt.Sprintf("assets/sfx/%s.ogg", name)
		stream := rl.LoadMusicStream(path)

		music = &stream
		tracks[name] = music
	}

	rl.PlayMusicStream(*music)
}

func PlaySound(name string) {
	sound, ok := sounds[name]
	if !ok {
		path := fmt.Sprintf("assets/sfx/%s.ogg", name)
		s := rl.LoadSound(path)

		sound = &s
		sounds[name] = sound
	}

	rl.PlaySound(*sound)
}
