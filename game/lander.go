package game

import (
	"image"
	_ "image/png"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
)

type Lander struct {
	x, y   float64       // Position of the lander
	vx, vy float64       // Velocity of the lander
	fuel   float64       // Amount of fuel left
	angle  float64       // Rotation angle of the lander (degrees)
	img    *ebiten.Image // Sprite
	sizeX  float64       // Size along X-axis
	sizeY  float64       // Size along Y-axis
}

func NewLander(x float64, y float64, landerPngFile string) Lander {
	img, err := loadImageFromFile(landerPngFile)
	if err != nil {
		log.Fatalf("failed to load image: %v", err)
	}
	landerImage := ebiten.NewImageFromImage(img)
	return Lander{
		x:     x,
		y:     y,
		fuel:  100,
		img:   landerImage,
		sizeX: float64(landerImage.Bounds().Size().X),
		sizeY: float64(landerImage.Bounds().Size().Y),
	}
}

func (lander *Lander) Update() {
	lander.vy += gravity

	// Handle controls
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		lander.angle -= rotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		lander.angle += rotationSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && lander.fuel > 0 {
		rad := lander.angle * math.Pi / 180
		lander.vx -= thrustPower * -math.Sin(rad)
		lander.vy -= thrustPower * math.Cos(rad)
		lander.fuel -= fuelConsumption
	}

	// Update position
	lander.x += lander.vx
	lander.y += lander.vy
}

func loadImageFromFile(filepath string) (image.Image, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	return img, err
}
