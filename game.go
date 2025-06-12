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
	lander    Lander
	terrain   *Terrain
	groundY   float64
	gameOver  bool
	success   bool
	obstacles []Rectangle
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

	// Boundary checks
	if g.lander.x < 0 {
		g.lander.x = 0
		g.lander.vx = 0
	} else if g.lander.x > screenWidth {
		g.lander.x = screenWidth
		g.lander.vx = 0
	}

	if g.lander.y < 0 {
		g.lander.y = 0
		g.lander.vy = 0
	}

	// Update lander rectangle for collision detection
	g.lander.UpdateRect()

	g.checkCollision()

	return nil
}

func (g *Game) checkCollision() {
	// Use the lander's rectangle for collision detection
	landerRect := g.lander.rect

	// Check collision with terrain
	if collision, isLandingPad := g.terrain.CheckCollision(landerRect); collision {
		g.gameOver = true

		// Check if landing was successful (soft landing on a landing pad)
		if isLandingPad && math.Abs(g.lander.vy) < 1.0 && math.Abs(g.lander.vx) < 1.0 && math.Abs(g.lander.angle) < 10 {
			log.Info("Lander landed successfully on landing pad")
			g.success = true
		} else {
			log.Info("Lander crashed!")
		}
	}

	// Check collision with obstacles
	for _, obstacle := range g.obstacles {
		corners := GetRectangleCorners(landerRect)

		// Create the four edges of the lander rectangle
		landerEdges := [4]Line{
			{corners[0], corners[1]}, // Top edge
			{corners[1], corners[2]}, // Right edge
			{corners[2], corners[3]}, // Bottom edge
			{corners[3], corners[0]}, // Left edge
		}

		// Check if any edge of the lander intersects with the obstacle
		for _, edge := range landerEdges {
			if LineRectCollision(edge, obstacle) {
				g.gameOver = true
				log.Info("Lander crashed into an obstacle!")
				return
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the terrain
	g.terrain.Draw(screen)

	// Draw obstacles
	for _, obstacle := range g.obstacles {
		corners := GetRectangleCorners(obstacle)

		// Draw the obstacle as a filled polygon
		vector.StrokeLine(screen, float32(corners[0].X), float32(corners[0].Y),
			float32(corners[1].X), float32(corners[1].Y), 2, color.RGBA{200, 100, 100, 255}, true)
		vector.StrokeLine(screen, float32(corners[1].X), float32(corners[1].Y),
			float32(corners[2].X), float32(corners[2].Y), 2, color.RGBA{200, 100, 100, 255}, true)
		vector.StrokeLine(screen, float32(corners[2].X), float32(corners[2].Y),
			float32(corners[3].X), float32(corners[3].Y), 2, color.RGBA{200, 100, 100, 255}, true)
		vector.StrokeLine(screen, float32(corners[3].X), float32(corners[3].Y),
			float32(corners[0].X), float32(corners[0].Y), 2, color.RGBA{200, 100, 100, 255}, true)
	}

	// Draw the lander
	op := &ebiten.DrawImageOptions{}

	// Move image center to upper-left corner
	op.GeoM.Translate(-g.lander.sizeX/2, -g.lander.sizeY/2)
	op.GeoM.Rotate(g.lander.angle * math.Pi / 180)
	op.GeoM.Translate(g.lander.x, g.lander.y)
	screen.DrawImage(g.lander.img, op)

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
