package kitman

import (
	"image/color"

	"github.com/jakecoffman/cp"
	"github.com/yohamta/ganim8/v3"
)

// 線條基礎模型
type BaseLine struct {
	X1     float64
	Y1     float64
	X2     float64
	Y2     float64
	PointA cp.Vector
	PointB cp.Vector
}

// 繪製線條
func (bl *BaseLine) Draw(dst *Image, clr color.Color) {
	driver.DrawLine(dst, clr, bl.X1, bl.Y1, bl.X2, bl.Y2)
}

// 精靈基礎模型
type BaseSprite struct {
	Body      *cp.Body
	Image     *Image
	W         float64
	H         float64
	Grid      *Grid
	Animation *ganim8.Animation
}

// 精靈安裝
func (bs *BaseSprite) Setup(path string) error {
	var err error
	if bs.Image, err = driver.GetImage(path); err != nil {
		return err
	}
	w, h := bs.Image.Size()
	bs.W, bs.H = float64(w), float64(h)
	return nil
}

// 建立影像網格
func (bs *BaseSprite) SetGrid(fw int, fh int, args ...int) {
	bs.Grid = ganim8.NewGrid(fw, fh, int(bs.W), int(bs.H), args...)
}

// 建立動畫
func (bs *BaseSprite) SetAnimation(durations interface{}, args ...interface{}) {
	bs.Animation = ganim8.New(bs.Image, bs.Grid.Frames(args...), durations)
}
