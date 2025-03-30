package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const landerPng = "img/lunar-lander.png"

type Lander struct {
	x, y   float64 // Position of the lander
	vx, vy float64 // Velocity of the lander
	fuel   float64 // Amount of fuel left
	angle  float64 // Rotation angle of the lander (degrees)
	img    *ebiten.Image
}

func NewLander() Lander {
	img, err := loadImageFromFile(landerPng)
	if err != nil {
		log.Fatalf("failed to load image: %v", err)
	}

	landerImage := ebiten.NewImageFromImage(img)
	return Lander{
		x:    screenWidth / 2,
		y:    50,
		fuel: 100,
		img:  landerImage,
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
