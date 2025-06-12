package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// TerrainSegment represents a line segment of the terrain
type TerrainSegment struct {
	Line      Line
	IsLanding bool // Whether this segment is a landing pad
}

// Terrain represents the lunar surface with landing pads and obstacles
type Terrain struct {
	Segments []TerrainSegment
	Color    color.RGBA
	PadColor color.RGBA
}

// NewTerrain creates a new terrain with the given number of segments
func NewTerrain(width, height int, numSegments int) *Terrain {
	segments := make([]TerrainSegment, numSegments)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Set colors
	terrainColor := color.RGBA{150, 150, 150, 255} // Gray color for regular terrain
	padColor := color.RGBA{0, 255, 0, 255}         // Green color for landing pads

	// Initialize variables for terrain generation
	segmentWidth := float64(width) / float64(numSegments)
	baseHeight := float64(height) - 50

	// Generate random landing pad location (somewhere in the middle segments)
	padSegment := numSegments/4 + r.Intn(numSegments/2)
	padWidth := 3 // Number of segments to use for the landing pad

	// Generate the terrain segments
	prevX := 0.0
	prevY := baseHeight + (r.Float64() * 50)

	for i := 0; i < numSegments; i++ {
		// Calculate the next point
		nextX := prevX + segmentWidth

		var nextY float64
		isLanding := false

		// If this is part of the landing pad, make it flat
		if i >= padSegment && i < padSegment+padWidth {
			nextY = prevY // Keep the same height for landing pad
			isLanding = true
		} else {
			// Random height with smoothing
			heightVariation := 30.0
			if i > 0 && segments[i-1].IsLanding {
				// Smoother transition from landing pad
				heightVariation = 15.0
			}
			nextY = prevY + (r.Float64()*heightVariation*2 - heightVariation)

			// Ensure the terrain stays within reasonable bounds
			if nextY < baseHeight-100 {
				nextY = baseHeight - 100
			} else if nextY > baseHeight+100 {
				nextY = baseHeight + 100
			}
		}

		// Create the segment
		segments[i] = TerrainSegment{
			Line: Line{
				Start: Point{X: prevX, Y: prevY},
				End:   Point{X: nextX, Y: nextY},
			},
			IsLanding: isLanding,
		}

		// Update for the next segment
		prevX = nextX
		prevY = nextY
	}

	return &Terrain{
		Segments: segments,
		Color:    terrainColor,
		PadColor: padColor,
	}
}

// Draw renders the terrain on the screen
func (t *Terrain) Draw(screen *ebiten.Image) {
	for _, segment := range t.Segments {
		color := t.Color
		lineWidth := float32(2.0)

		if segment.IsLanding {
			color = t.PadColor
			lineWidth = 3.0
		}

		vector.StrokeLine(
			screen,
			float32(segment.Line.Start.X),
			float32(segment.Line.Start.Y),
			float32(segment.Line.End.X),
			float32(segment.Line.End.Y),
			lineWidth,
			color,
			true,
		)
	}
}

// CheckCollision checks if a rectangle collides with any terrain segment
func (t *Terrain) CheckCollision(rect Rectangle) (bool, bool) {
	// Get the four corners of the rectangle
	corners := GetRectangleCorners(rect)

	// Create the four edges of the rectangle
	edges := [4]Line{
		{corners[0], corners[1]}, // Top edge
		{corners[1], corners[2]}, // Right edge
		{corners[2], corners[3]}, // Bottom edge
		{corners[3], corners[0]}, // Left edge
	}

	// Check collision with each terrain segment
	for _, segment := range t.Segments {
		// Check if any edge of the rectangle intersects with this terrain segment
		for _, edge := range edges {
			if _, intersects := LineLineIntersection(edge, segment.Line); intersects {
				return true, segment.IsLanding
			}
		}
	}

	return false, false
}

// IsPointBelowTerrain checks if a point is below the terrain
func (t *Terrain) IsPointBelowTerrain(p Point) bool {
	// Find the terrain segment that contains this x-coordinate
	for _, segment := range t.Segments {
		if p.X >= segment.Line.Start.X && p.X <= segment.Line.End.X {
			// Calculate the y-coordinate of the terrain at this x-coordinate
			// using linear interpolation
			ratio := (p.X - segment.Line.Start.X) / (segment.Line.End.X - segment.Line.Start.X)
			terrainY := segment.Line.Start.Y + ratio*(segment.Line.End.Y-segment.Line.Start.Y)

			// If the point's y-coordinate is greater than the terrain's y-coordinate,
			// it's below the terrain
			return p.Y >= terrainY
		}
	}

	return false
}
