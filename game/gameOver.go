package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameOver struct {
}

func (g *GameOver) Update(game *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		Displayed_scene = &MainScene{}
	}
}

func (g *GameOver) Draw(screen *ebiten.Image, game *Game) {
	ebitenutil.DebugPrint(screen, "Game Over \n Press Enter to retry")
}
