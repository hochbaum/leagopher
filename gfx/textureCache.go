package gfx

import "github.com/gen2brain/raylib-go/raylib"

var textureCache = make(map[string]rl.Texture2D)

func GetTexture(path string) *rl.Texture2D {
	texture, ok := textureCache[path]
	if !ok {
		texture = rl.LoadTexture(path)
	}

	return &texture
}
