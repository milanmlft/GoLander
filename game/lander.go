package game

import (
	"image"
	_ "image/png"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
)

const (
	gravity           = 0.01
	thrustPower       = 0.05
	fuelConsumption   = 0.05
	rotationPerSecond = math.Pi
)

type Lander struct {
	x, y     float64       // Position of the lander
	vx, vy   float64       // Velocity of the lander
	fuel     float64       // Amount of fuel left
	rotation float64       // Rotation angle of the lander (radian)
	sprite   *ebiten.Image // Sprite
	sizeX    float64       // Size along X-axis
	sizeY    float64       // Size along Y-axis
}

func NewLander(x float64, y float64, landerPngFile string) Lander {
	img, err := loadImageFromFile(landerPngFile)
	if err != nil {
		log.Fatalf("failed to load image: %v", err)
	}
	landerImage := ebiten.NewImageFromImage(img)
	return Lander{
		x:        x,
		y:        y,
		fuel:     100,
		rotation: 0,
		sprite:   landerImage,
		sizeX:    float64(landerImage.Bounds().Dx()),
		sizeY:    float64(landerImage.Bounds().Dy()),
	}
}

func (lander *Lander) Update() {
	lander.vy += gravity
	rotationSpeed := rotationPerSecond / float64(ebiten.TPS())

	// Handle controls
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		lander.rotation -= rotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		lander.rotation += rotationSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && lander.fuel > 0 {
		lander.vx -= thrustPower * -math.Sin(lander.rotation)
		lander.vy -= thrustPower * math.Cos(lander.rotation)
		lander.fuel -= fuelConsumption
	}

	// Update position
	lander.x += lander.vx
	lander.y += lander.vy
}

func (lander *Lander) Draw(screen *ebiten.Image) {
	halfW := lander.sizeX / 2
	halfH := lander.sizeY / 2

	op := &ebiten.DrawImageOptions{}

	// Move image center to upper-left corner
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(lander.rotation)
	op.GeoM.Translate(lander.x, lander.y)
	screen.DrawImage(lander.sprite, op)
}

func loadImageFromFile(filepath string) (image.Image, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	return img, err
}
