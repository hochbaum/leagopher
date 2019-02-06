package gfx

import "github.com/gen2brain/raylib-go/raylib"

var textureCache = make(map[string]raylib.Texture2D)

func GetTexture(path string) *raylib.Texture2D {
	texture, ok := textureCache[path]
	if !ok {
		texture = raylib.LoadTexture(path)
	}

	return &texture
}
