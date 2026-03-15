package game

import "math"

// Rectangle represents a rotatable rectangle
type Rectangle struct {
	Center  Point    // Center of the rectangle
	Width   float64  // Width of the rectangle
	Height  float64  // Height of the rectangle
	Angle   float64  // Rotation angle in radians
	Corners [4]Point // Cached corner positions
}

// Line represents a line segment between two points
type Line struct {
	Start, End Point
}

// Point represents a 2D point
type Point struct {
	X, Y float64
}

func LineRectCollision(line Line, rect Rectangle) bool {
	return false
}

// LineLineIntersection calculates the intersection point of two line segments
// and returns the intersection point and whether an intersection exists
func LineLineIntersection(line1, line2 Line) (Point, bool) {
	// Line 1 represented as a1x + b1y = c1
	a1 := line1.End.Y - line1.Start.Y
	b1 := line1.Start.X - line1.End.X
	c1 := a1*line1.Start.X + b1*line1.Start.Y

	// Line 2 represented as a2x + b2y = c2
	a2 := line2.End.Y - line2.Start.Y
	b2 := line2.Start.X - line2.End.X
	c2 := a2*line2.Start.X + b2*line2.Start.Y

	// Determinant
	determinant := a1*b2 - a2*b1
	if math.Abs(determinant) < 1e-10 {
		return Point{}, false
	}

	// Find intersection point
	x := (b2*c1 - b1*c2) / determinant
	y := (a1*c2 - a2*c1) / determinant

	// Check if the intersection point is on both line segments
	intersection := Point{X: x, Y: y}
	onLine1 := IsPointOnLine(intersection, line1.Start, line1.End, 1e-10)
	onLine2 := IsPointOnLine(intersection, line2.Start, line2.End, 1e-10)

	return intersection, onLine1 && onLine2
}

// IsPointOnLine checks if a point lies on a line segment
func IsPointOnLine(p, lineStart, lineEnd Point, tolerance float64) bool {
	// Calculate distances
	d1 := distance(p, lineStart)
	d2 := distance(p, lineEnd)
	lineLen := distance(lineStart, lineEnd)

	// Check if point is on line (with small tolerance for floating point errors)
	return math.Abs(d1+d2-lineLen) < tolerance
}

// distance calculates Euclidean distance between two points
func distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}
