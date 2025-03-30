package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	log "github.com/sirupsen/logrus"
)

const (
	gravity       = 0.1 // Gravity acceleration
	thrustPower   = 0.15
	rotationSpeed = 2.0
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
		// Apply thrust (opposite to angle)
		rad := g.lander.angle * math.Pi / 180
		g.lander.vx -= thrustPower * math.Sin(rad)
		g.lander.vy -= thrustPower * math.Cos(rad)
		g.lander.fuel -= 0.1
	}

	// Update position
	g.lander.x += g.lander.vx
	g.lander.y += g.lander.vy

	// Check collision with ground
	if g.lander.y >= g.groundY-20 { // Assuming lander is ~20px tall
		log.Info("Lander crashed!")
		g.gameOver = true
		// Check if landing was successful (soft landing)
		if math.Abs(g.lander.vy) < 1.0 && math.Abs(g.lander.vx) < 1.0 && math.Abs(g.lander.angle) < 10 {
			log.Info("Lander landed successfully")
			g.success = true
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the lander
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-10, 10)
	op.GeoM.Rotate(g.lander.angle * math.Pi / 180)
	op.GeoM.Translate(g.lander.x, g.lander.y)
	screen.DrawImage(g.lander.img, op)

	// Draw the ground
	// ebitenutil.DrawLine(screen, 0, g.groundY, screenWidth, g.groundY, color.White)
	vector.StrokeLine(screen, 0, float32(g.groundY), screenWidth, float32(g.groundY), 1, color.White, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
