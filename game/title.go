package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Title struct {
}

func (t *Title) Update(game *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		Displayed_scene = &MainScene{}
	}

	Global_valiable += 1
}

func (t *Title) Draw(screen *ebiten.Image, game *Game) {
	ebitenutil.DebugPrint(screen, "Press Enter to Start")
}
