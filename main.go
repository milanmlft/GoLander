package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
)

const (
	screenWidth  = 1200
	screenHeight = 800
	landerPng    = "img/lander.png"
)

func main() {
	game := &Game{
		lander:  NewLander(screenWidth/2, screenHeight/20, landerPng),
		groundY: screenHeight - 50,
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GoLander")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
