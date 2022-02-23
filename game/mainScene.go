package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"gonum.org/v1/gonum/mat"

	"github.com/krile136/breakout/internal/draw"
)

const (
	ball_coefficient  float64 = 0.3
	block_coefficient float64 = 0.5
	radius            float64 = 32 * ball_coefficient / 2

	bar_y = 290

	blockWidth  int = 64
	blockHeight int = 24
)

var (
	ballCenterX   float64 = screenWidth / 2
	ballCenterY   float64 = screenHeight / 2
	velAngle      float64 = math.Pi / 4
	velocity      float64 = 5
	mouse_x       int
	mouse_y       int
	fixed_mouse_x float64
	blocks        [5][5]int
)

type MainScene struct {
}

func (m *MainScene) Update(game *Game) {
	if is_just_changed {
		// シーンが切り替わったときにブロックのライフを初期化
		fillBlocksLife()
	}

	mouse_x, mouse_y = ebiten.CursorPosition()
	fixed_mouse_x = math.Max(0, math.Min(float64(mouse_x), screenWidth))

	// 進行方向の角度による回転行列を生成
	basicPostureArray := []float64{math.Cos(velAngle), -math.Sin(velAngle), math.Sin(velAngle), math.Cos(velAngle)}
	postureRotateMatrix := mat.NewDense(2, 2, basicPostureArray)

	// 速度ベクトルを生成(Y軸方向に正の方向がvelocityとなるようなベクトル)
	basicVelocityArray := []float64{0, -velocity}
	velocityVector := mat.NewDense(2, 1, basicVelocityArray)

	// 移動ベクトルを生成
	moveVector := mat.NewDense(2, 1, nil)
	moveVector.Product(postureRotateMatrix, velocityVector)

	// ボールの中心位置を移動させる
	// 壁の橋にあたったときにX,Yそれぞれ角度を反転させる
	// barに触れたときにYを反転させる
	prevBallCenterX := ballCenterX + moveVector.At(0, 0)
	if prevBallCenterX-radius < 0 || prevBallCenterX+radius > screenWidth {
		ballCenterX -= moveVector.At(0, 0)
		velAngle = math.Pi*2 - velAngle
	} else {
		ballCenterX += moveVector.At(0, 0)
	}

	prevBallCenterY := ballCenterY + moveVector.At(1, 0)
	if prevBallCenterY-radius < 0 || prevBallCenterY+radius > screenHeight || (prevBallCenterY+radius > float64(bar_y) && prevBallCenterX >= (fixed_mouse_x-float64(blockWidth)*ball_coefficient) && prevBallCenterX <= (fixed_mouse_x+float64(blockWidth)*ball_coefficient)) {
		ballCenterY -= moveVector.At(1, 0)
		velAngle = math.Pi - velAngle
	} else {
		ballCenterY += moveVector.At(1, 0)
	}

reflect_label:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if blocks[i][j] > 0 {
				bx := 40 + 16 + float64(blockWidth*j)*block_coefficient
				by := 30 + float64(blockHeight*i)*block_coefficient
				var is_hit bool = reflectBallWithBlock(prevBallCenterX, prevBallCenterY, bx, by, moveVector)
				if is_hit {
					blocks[i][j] -= 1
					break reflect_label
				}
			}
		}
	}
}

func (M *MainScene) Draw(screen *ebiten.Image, game *Game) {
	// ボールを描画
	draw.DrawWithoutRect(screen, "ball", ball_coefficient, ballCenterX, ballCenterY, 0)

	// バーを描画
	draw.DrawWithoutRect(screen, "bar", ball_coefficient, fixed_mouse_x, bar_y, 0)

	// ブロックを描画
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if blocks[i][j] > 0 {
				bx := 40 + 16 + float64(blockWidth*j)*block_coefficient
				by := 30 + float64(blockHeight*i)*block_coefficient
				draw.Draw(screen, "blocks", block_coefficient, bx, by, 0, 64*blocks[i][j], 0, blockWidth, blockHeight)
			}
		}
	}

}

func reflectBallWithBlock(nb_x, nb_y, bl_x, bl_y float64, mv *mat.Dense) bool {
	// 移動後のボールがブロックの中にいるとき
	var is_hit bool = false
	if nb_x+radius >= (bl_x-block_coefficient*float64(blockWidth/2)) && nb_x-radius <= (bl_x+block_coefficient*float64(blockWidth/2)) && nb_y+radius >= (bl_y-block_coefficient*float64(blockHeight/2)) && nb_y-radius <= (bl_y+block_coefficient*float64(blockHeight/2)) {
		is_hit = true
		if ballCenterX >= (bl_x-block_coefficient*float64(blockWidth/2)) && ballCenterX <= (bl_x+block_coefficient*float64(blockWidth/2)) {
			//  上下反転
			ballCenterY -= mv.At(1, 0)
			velAngle = math.Pi - velAngle
		} else if ballCenterY+radius >= (bl_y-block_coefficient*float64(blockHeight/2)) && ballCenterY-radius <= (bl_y+block_coefficient*float64(blockHeight/2)) {
			// 左右反転
			ballCenterX -= mv.At(0, 0)
			velAngle = math.Pi*2 - velAngle
		} else {
			// 上下左右反転
			ballCenterY -= mv.At(1, 0)
			ballCenterX -= mv.At(0, 0)
			velAngle = math.Pi + velAngle
		}
	}

	return is_hit
}

func fillBlocksLife() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			blocks[i][j] = 5 - i
		}
	}
}
