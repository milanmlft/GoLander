package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Rect represents a rotatable rectangle
type Rect struct {
	X      float64 // X coordinate
	Y      float64 // Y coordinate
	Width  float64 // Width of the rectangle
	Height float64 // Height of the rectangle
}

func NewRectangle(x, y, width, height float64) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r Rect) MaxX() float64 {
	return r.X + r.Width
}

func (r Rect) MaxY() float64 {
	return r.Y + r.Height
}

func (r Rect) Intersects(other Rect) bool {
	return r.X <= other.MaxX() &&
		other.X <= r.MaxX() &&
		r.Y <= other.MaxY() &&
		other.Y <= r.MaxY()
}

// Draw draws the rectangle to the screen, mainly for debugging
func (r Rect) Draw(screen *ebiten.Image) {
	vector.StrokeRect(screen,
		float32(r.X), float32(r.Y),
		float32(r.Width), float32(r.Height),
		1, color.White, true)
}
