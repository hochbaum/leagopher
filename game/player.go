package game

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/lhochbaum/leagopher/gfx"
	"github.com/lhochbaum/leagopher/math"
	"image"
)

type Player struct {
	Position *math.Vector2
	Texture  *rl.Texture2D
	Bounds *image.Rectangle
}

func NewPlayer() *Player {
	return &Player{
		Position: math.NewVector2(0, 0),
		Texture:  gfx.GetTexture("assets/gfx/player/player.png"),
	}
}

func (p *Player) Update(dt float32) {

}