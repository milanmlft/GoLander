package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	game := &Game{
		lander: NewLander(),
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GoLander")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
