package main

import "math"

// Rectangle represents a rotatable rectangle
type Rectangle struct {
	Center  Point    // Center of the rectangle
	Width   float64  // Width of the rectangle
	Height  float64  // Height of the rectangle
	Angle   float64  // Rotation angle in radians
	Corners [4]Point // Cached corner positions
}

// GetRectangleCorners calculates and returns the four corners of a rotated rectangle
func GetRectangleCorners(rect Rectangle) [4]Point {
	// Calculate the half-width and half-height
	halfWidth := rect.Width / 2
	halfHeight := rect.Height / 2

	// Calculate sin and cos of the angle
	sinAngle := math.Sin(rect.Angle)
	cosAngle := math.Cos(rect.Angle)

	// Calculate the four corners (top-left, top-right, bottom-right, bottom-left)
	corners := [4]Point{
		// Top-left
		{
			X: rect.Center.X + (-halfWidth*cosAngle - halfHeight*sinAngle),
			Y: rect.Center.Y + (-halfWidth*sinAngle + halfHeight*cosAngle),
		},
		// Top-right
		{
			X: rect.Center.X + (halfWidth*cosAngle - halfHeight*sinAngle),
			Y: rect.Center.Y + (halfWidth*sinAngle + halfHeight*cosAngle),
		},
		// Bottom-right
		{
			X: rect.Center.X + (halfWidth*cosAngle + halfHeight*sinAngle),
			Y: rect.Center.Y + (halfWidth*sinAngle - halfHeight*cosAngle),
		},
		// Bottom-left
		{
			X: rect.Center.X + (-halfWidth*cosAngle + halfHeight*sinAngle),
			Y: rect.Center.Y + (-halfWidth*sinAngle - halfHeight*cosAngle),
		},
	}

	return corners
}

// Line represents a line segment between two points
type Line struct {
	Start, End Point
}

// Point represents a 2D point
type Point struct {
	X, Y float64
}

// LineRectCollision checks if a line segment intersects with a rotated rectangle
func LineRectCollision(line Line, rect Rectangle) bool {
	// Get the four corners of the rectangle
	corners := GetRectangleCorners(rect)

	// Create the four edges of the rectangle
	edges := [4]Line{
		{corners[0], corners[1]}, // Top edge
		{corners[1], corners[2]}, // Right edge
		{corners[2], corners[3]}, // Bottom edge
		{corners[3], corners[0]}, // Left edge
	}

	// Check if the line intersects with any of the rectangle's edges
	for _, edge := range edges {
		if _, intersects := LineLineIntersection(line, edge); intersects {
			return true
		}
	}

	// Check if either endpoint of the line is inside the rectangle
	if IsPointInsideRectangle(line.Start, rect) || IsPointInsideRectangle(line.End, rect) {
		return true
	}

	return false
}

// LineIntersection calculates the intersection point of two line segments
// Returns the intersection point and whether an intersection exists
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

// IsPointInsideRectangle checks if a point is inside a rotated rectangle
func IsPointInsideRectangle(p Point, rect Rectangle) bool {
	// Get the corners of the rectangle
	corners := GetRectangleCorners(rect)

	// Create vectors from the point to each corner
	vectors := [4]struct{ x, y float64 }{
		{corners[0].X - p.X, corners[0].Y - p.Y},
		{corners[1].X - p.X, corners[1].Y - p.Y},
		{corners[2].X - p.X, corners[2].Y - p.Y},
		{corners[3].X - p.X, corners[3].Y - p.Y},
	}

	// Calculate cross products between adjacent vectors
	cross1 := vectors[0].x*vectors[1].y - vectors[0].y*vectors[1].x
	cross2 := vectors[1].x*vectors[2].y - vectors[1].y*vectors[2].x
	cross3 := vectors[2].x*vectors[3].y - vectors[2].y*vectors[3].x
	cross4 := vectors[3].x*vectors[0].y - vectors[3].y*vectors[0].x

	// If all cross products have the same sign, the point is inside the rectangle
	return (cross1 > 0 && cross2 > 0 && cross3 > 0 && cross4 > 0) ||
		(cross1 < 0 && cross2 < 0 && cross3 < 0 && cross4 < 0)
}

// distance calculates Euclidean distance between two points
func distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}
