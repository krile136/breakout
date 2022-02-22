package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/breakout/game"
)

func main() {

	game, err := game.NewGame()
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(240, 320)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
