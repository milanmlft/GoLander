package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
	landerPng    = "img/lander.png"
)

func main() {
	game := &Game{
		lander:  NewLander(screenWidth/2, screenHeight/2, landerPng),
		groundY: screenHeight - 50,
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GoLander")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
