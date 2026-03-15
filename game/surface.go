package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Surface struct {
	segments []Segment
}

type Segment struct {
	position Vector
	angle    float64
	length   float64
}

func NewSurface() Surface {
	groundSegment := Segment{
		position: Vector{0, screenHeight - 50},
		angle:    0,
		length:   screenWidth,
	}
	return Surface{
		segments: []Segment{groundSegment},
	}
}

func (s *Surface) Draw(screen *ebiten.Image) {
	for _, seg := range s.segments {
		seg.Draw(screen)
	}
}

func (s *Segment) Draw(screen *ebiten.Image) {
	x0 := float32(s.position.X)
	y0 := float32(s.position.Y)
	x1 := x0 + float32(s.length*math.Cos(s.angle))
	y1 := y0 + float32(s.length*math.Sin(s.angle))
	vector.StrokeLine(screen, x0, y0, x1, y1, 5, color.White, true)
}
