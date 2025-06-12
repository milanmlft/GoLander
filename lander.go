package main

import (
	"image"
	_ "image/png"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
)

type Lander struct {
	x, y   float64 // Position of the lander
	vx, vy float64 // Velocity of the lander
	fuel   float64 // Amount of fuel left
	angle  float64 // Rotation angle of the lander (degrees)
	img    *ebiten.Image
	sizeX  float64
	sizeY  float64
	rect   Rectangle // For collision detection
}

func NewLander(x float64, y float64, landerPngFile string) Lander {
	img, err := loadImageFromFile(landerPngFile)
	if err != nil {
		log.Fatalf("failed to load image: %v", err)
	}
	landerImage := ebiten.NewImageFromImage(img)
	sizeX := float64(landerImage.Bounds().Size().X)
	sizeY := float64(landerImage.Bounds().Size().Y)

	return Lander{
		x:     x,
		y:     y,
		fuel:  100,
		img:   landerImage,
		sizeX: sizeX,
		sizeY: sizeY,
		rect: Rectangle{
			Center: Point{X: x, Y: y},
			Width:  sizeX,
			Height: sizeY,
			Angle:  0,
		},
	}
}

func loadImageFromFile(filepath string) (image.Image, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	return img, err
}

// UpdateRect updates the rectangle used for collision detection
func (l *Lander) UpdateRect() {
	l.rect = Rectangle{
		Center: Point{X: l.x, Y: l.y},
		Width:  l.sizeX,
		Height: l.sizeY,
		Angle:  l.angle * math.Pi / 180, // Convert degrees to radians
	}
}
