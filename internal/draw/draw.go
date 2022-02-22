package draw

import (
	"image"
	_ "image/png"
	"math"
	"path"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/krile136/breakout/internal/assets"
)

var (
	images = map[string]*ebiten.Image{}
)

// 画像を描画する (Rectによる切り抜きあり)
func Draw(screen *ebiten.Image, key string, coefficient, x_coodinates, y_coodinates, angle float64, rect_base_x, rect_base_y, rect_area_x, rect_area_y int) {
	var image *ebiten.Image = images[key].SubImage(image.Rect(64, 0, 128, 24)).(*ebiten.Image)
	calcAndDisplay(screen, image, coefficient, x_coodinates, y_coodinates, angle)
}

func DrawWithoutRect(screen *ebiten.Image, key string, coefficient, x_coodinates, y_coodinates, angle float64) {
	var image *ebiten.Image = images[key]
	calcAndDisplay(screen, image, coefficient, x_coodinates, y_coodinates, angle)
}

func calcAndDisplay(screen *ebiten.Image, image *ebiten.Image, coefficient, x_coodinates, y_coodinates, angle float64) {
	// 画像のサイズを取得
	w, h := image.Size()

	// 係数で画像を拡大/縮小したときの大きさを計算しておく
	var sw, sh float64 = float64(w) * coefficient, float64(h) * coefficient

	// オプションを宣言
	op := &ebiten.DrawImageOptions{}

	// 画像を拡大/縮小する
	op.GeoM.Scale(coefficient, coefficient)

	// 縮小したサイズに合わせて、画面の左上に縦横半分めり込む形にする
	op.GeoM.Translate(-sw/2, -sh/2)

	// 画像を画面の左上を中心に回転させる（縦横半分めり込んでいるので、中心で回転することになる)
	op.GeoM.Rotate(angle / 180 * math.Pi)

	// 好きな位置へ移動させる
	op.GeoM.Translate(x_coodinates, y_coodinates)

	// 画像を描画する
	screen.DrawImage(image, op)

}

// 画像リソースを読み込む
func LoadImages() error {
	// asset.go から読み込みたい画像のディレクトリパス
	const dir = "images"

	// ディレクトリの内容を読み取る
	ents, err := assets.Assets.ReadDir(dir)
	if err != nil {
		return err
	}

	// ディレクトリの中身を取り出し、pngでなければ画像として登録
	for _, ent := range ents {
		name := ent.Name()

		// 拡張子(png)のチェック
		ext := filepath.Ext(name)
		if ext != ".png" {
			continue
		}

		// 画像を読み込む
		f, err := assets.Assets.Open(path.Join(dir, name))
		if err != nil {
			return err
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			return err
		}

		// imagesの配列に画像を登録
		key := name[:len(name)-len(ext)]
		images[key] = ebiten.NewImageFromImage(img)
	}
	return nil
}
