package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	lander   Lander
	groundY  float64
	gameOver bool
}

func (g *Game) Update() error {
	if g.gameOver {
		return nil
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
