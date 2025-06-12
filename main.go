package main

import (
	"math"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
)

const (
	screenWidth  = 1200
	screenHeight = 800
	landerSvg    = "img/lander.svg"
)

func main() {
	// Seed the random number generator for terrain generation
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Convert SVG to PNG if needed
	landerPng, err := CheckAndConvertSVGToPNG(landerSvg)
	if err != nil {
		log.Warnf("Failed to convert SVG to PNG: %v. Using SVG file path anyway.", err)
		landerPng = filepath.Join(filepath.Dir(landerSvg), "lander.png")
	}

	// Create terrain with 40 segments
	terrain := NewTerrain(screenWidth, screenHeight, 40)

	// Create obstacles
	obstacles := createRandomObstacles(3)

	game := &Game{
		lander:    NewLander(screenWidth/2, screenHeight/20, landerPng),
		terrain:   terrain,
		groundY:   screenHeight - 50,
		obstacles: obstacles,
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GoLander")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// createRandomObstacles creates a specified number of random obstacle rectangles
func createRandomObstacles(count int) []Rectangle {
	obstacles := make([]Rectangle, count)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define the vertical area where obstacles can appear
	minY := screenHeight / 3
	maxY := screenHeight - 150

	for i := 0; i < count; i++ {
		// Random position for the obstacle
		x := float64(r.Intn(screenWidth-200) + 100) // Avoid edges
		y := float64(r.Intn(maxY-minY) + minY)      // Middle to lower part of screen

		// Random size (not too large)
		width := float64(r.Intn(60) + 40)
		height := float64(r.Intn(60) + 40)

		// Random rotation
		angle := float64(r.Intn(360)) * (math.Pi / 180)

		obstacles[i] = Rectangle{
			Center: Point{X: x, Y: y},
			Width:  width,
			Height: height,
			Angle:  angle,
		}
	}

	return obstacles
}
