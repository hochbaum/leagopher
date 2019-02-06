package game

import (
	"github.com/lhochbaum/leagopher/math"
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/lhochbaum/leagopher/gfx"
)

type Player struct {
	Position *math.Vector2
	Texture *raylib.Texture2D
}

func NewPlayer() *Player {
	return &Player{
		Position: math.NewVector2(0, 0),
		Texture: gfx.GetTexture("assets/gfx/player/player.png"),
	}
}

func (p *Player) Update(dt float32) {

}