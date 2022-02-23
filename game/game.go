package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/krile136/breakout/internal/draw"
)

const (
	screenWidth  float64 = 240
	screenHeight float64 = 320
)

var (
	Displayed_scene Scene = &Title{} // 各シーンを持ち回る変数
	Global_valiable int   = 0
	is_just_changed bool  = false
)

type Game struct {
	resourceLoadedCh chan error
	scene            Scene
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 240, 320
}

func (g *Game) Update() error {
	// 実行するシーンを変数より取得
	if g.scene != Displayed_scene {
		is_just_changed = true
		g.scene = Displayed_scene
	}
	g.scene.Update(g)
	if is_just_changed {
		is_just_changed = false
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen, g)
}

func NewGame() (*Game, error) {

	game := &Game{
		resourceLoadedCh: make(chan error),
	}

	// 画像リソース読み込み
	go func() {
		err := draw.LoadImages()
		if err != nil {
			game.resourceLoadedCh <- err
			return
		}

		close(game.resourceLoadedCh)
	}()

	return game, nil
}

// 各シーンは必ず持たないと行けない関数
type Scene interface {
	Update(g *Game)
	Draw(screen *ebiten.Image, g *Game)
}
