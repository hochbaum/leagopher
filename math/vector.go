package math

type Vector2 struct {
	X int32
	Y int32
}

type Vector2f struct {
	X float32
	Y float32
}

func NewVector2(x, y int32) *Vector2 {
	return &Vector2{
		X: x,
		Y: y,
	}
}

func NewVector2f(x, y float32) *Vector2f {
	return &Vector2f{
		X: x,
		Y: y,
	}
}
