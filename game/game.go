// Package game implements GoLander's game logic
package game

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
	screenWidth  = 1200
	screenHeight = 800
	landerPng    = "img/lander.png"
)

type Game struct {
	lander   Lander
	groundY  float64
	gameOver bool
	success  bool
}

func NewGame() *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GoLander")
	return &Game{
		lander:  NewLander(landerPng),
		groundY: screenHeight - 50,
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		return nil
	}

	g.lander.Update()
	g.checkCollision()

	return nil
}

func (g *Game) checkCollision() {
	// Assuming Lander is a rectangle

	if g.lander.position.Y >= g.groundY-20 {
		g.gameOver = true
		// Check if landing was successful (soft landing)
		if math.Abs(g.lander.velocity.Y) < 1.0 && math.Abs(g.lander.velocity.X) < 1.0 && math.Abs(g.lander.rotation) < 10 {
			log.Info("Lander landed successfully")
			g.success = true
		} else {
			log.Info("Lander crashed!")
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.lander.Draw(screen)

	// Draw the ground
	vector.StrokeLine(screen, 0, float32(g.groundY), screenWidth, float32(g.groundY), 5, color.White, true)

	ebitenutil.DebugPrintAt(screen,
		"Fuel: "+formatFloat(g.lander.fuel, 1)+
			"\nVelocity X: "+formatFloat(g.lander.velocity.X, 2)+
			"\nVelocity Y: "+formatFloat(g.lander.velocity.Y, 2),
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
