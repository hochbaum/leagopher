package gfx

type Renderer struct {
	Queue []func()
}

func NewRenderer() *Renderer {
	return &Renderer{
		Queue: make([]func(), 0),
	}
}

func (r *Renderer) Add(f func()) {
	r.Queue = append(r.Queue, f)
}

func (r *Renderer) Set(index int, f func()) {
	r.Queue[index] = f
}

func (r *Renderer) Render() {
	for _, f := range r.Queue {
		f()
	}
}
