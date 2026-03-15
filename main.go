package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milanmlft/GoLander/game"
	log "github.com/sirupsen/logrus"
)

func main() {
	game := game.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
