package main

import (
	"image/color"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	log "github.com/sirupsen/logrus"
)

const (
	gravity         = 0.01
	thrustPower     = 0.05
	fuelConsumption = 0.5
	rotationSpeed   = 2.0
)

type Game struct {
	lander   Lander
	groundY  float64
	gameOver bool
	success  bool
}

func (g *Game) Update() error {
	if g.gameOver {
		return nil
	}

	g.lander.vy += gravity

	// Handle controls
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.lander.angle -= rotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.lander.angle += rotationSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.lander.fuel > 0 {
		rad := g.lander.angle * math.Pi / 180
		g.lander.vx -= thrustPower * -math.Sin(rad)
		g.lander.vy -= thrustPower * math.Cos(rad)
		g.lander.fuel -= fuelConsumption
	}

	// Update position
	g.lander.x += g.lander.vx
	g.lander.y += g.lander.vy

	g.checkCollision()

	return nil
}

func (g *Game) checkCollision() {
	centerY := g.lander.y
	width := g.lander.sizeX
	height := g.lander.sizeY

	// Assuming Lander is a rectangle

	if g.lander.y >= g.groundY-20 {
		g.gameOver = true
		// Check if landing was successful (soft landing)
		if math.Abs(g.lander.vy) < 1.0 && math.Abs(g.lander.vx) < 1.0 && math.Abs(g.lander.angle) < 10 {
			log.Info("Lander landed successfully")
			g.success = true
		} else {
			log.Info("Lander crashed!")
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the lander
	op := &ebiten.DrawImageOptions{}

	// Move image center to upper-left corner
	op.GeoM.Translate(-g.lander.sizeX/2, -g.lander.sizeY/2)
	op.GeoM.Rotate(g.lander.angle * math.Pi / 180)
	op.GeoM.Translate(g.lander.x, g.lander.y)
	screen.DrawImage(g.lander.img, op)

	// Draw the ground
	vector.StrokeLine(screen, 0, float32(g.groundY), screenWidth, float32(g.groundY), 5, color.White, true)

	ebitenutil.DebugPrintAt(screen,
		"Fuel: "+formatFloat(g.lander.fuel, 1)+
			"\nVelocity X: "+formatFloat(g.lander.vx, 2)+
			"\nVelocity Y: "+formatFloat(g.lander.vy, 2),
		10, 10)

	if g.gameOver {
		msg := "CRASHED!"
		if g.success {
			msg = "SUCCESSFUL LANDING!"
		}
		ebitenutil.DebugPrintAt(screen, msg, screenWidth/2-100, screenHeight/2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func formatFloat(f float64, decimals int) string {
	return strconv.FormatFloat(f, 'f', decimals, 64)
}
