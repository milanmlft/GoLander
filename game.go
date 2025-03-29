package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	lander   Lander
	gameOver bool
}

func (g *Game) Update() error {
	if g.gameOver {
		return nil
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-10, 10)
	op.GeoM.Rotate(g.lander.angle * math.Pi / 180)
	op.GeoM.Translate(g.lander.x, g.lander.y)
	screen.DrawImage(g.lander.img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
