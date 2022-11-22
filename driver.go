package kitman

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/davecgh/go-spew/spew"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

// 啟動器結構
type Driver struct {
	assets *embed.FS // 嵌入素材的檔案系統
	SW     float64   // 屏幕寬
	SH     float64   // 屏幕高
	HSW    float64   // 屏幕半寬
	HSH    float64   // 屏幕半高
	Font   struct {
		Path string    // 字型檔路徑
		Face font.Face // 字體
		DPI  float64   // 文字解析度
		Size float64   // 文字大小
		fixX float64
		fixY float64
	}
}

// 偵錯列印
func (d *Driver) Dump(a ...interface{}) {
	spew.Dump(a...)
}

// 偵錯列印至檔案
func (d *Driver) Fdump(w io.Writer, a ...interface{}) {
	spew.Fdump(w, a...)
}

// 設定字型
func (d *Driver) SetFont(path string, size int) error {
	var (
		err error
		b   []byte
		tt  *sfnt.Font
	)
	if b, err = d.assets.ReadFile(path); err != nil {
		return err
	}
	if tt, err = opentype.Parse(b); err != nil {
		return err
	}
	if d.Font.Face, err = opentype.NewFace(tt, &opentype.FaceOptions{
		DPI:     float64(size),
		Size:    float64(size),
		Hinting: font.HintingFull,
	}); err != nil {
		return err
	}
	d.Font.Path = path
	d.Font.DPI = float64(size)
	d.Font.Size = float64(size)
	d.SetFontFixPosition(2, d.Font.Size/2)
	return nil
}

// 設定字型原點座標補正
func (d *Driver) SetFontFixPosition(x, y float64) {
	d.Font.fixX, d.Font.fixY = x, y
}

// 繪製文字
func (d *Driver) DrawText(dst *Image, str string, clr color.Color, x, y float64) {
	text.Draw(dst, str, d.Font.Face, int(x)+int(d.Font.fixX), int(y)+int(d.Font.fixY), clr)
}

// 繪製線條
func (d *Driver) DrawLine(dst *Image, clr color.Color, x1, y1, x2, y2 float64) {
	ebitenutil.DrawLine(dst, x1, y1, x2, y2, clr)
}

// 繪製正方形
func (d *Driver) DrawRect(dst *Image, clr color.Color, w, h, x, y float64) {
	ebitenutil.DrawRect(dst, x, y, w, h, clr)
}

// 繪製圓形
func (d *Driver) DrawCircle(dst *Image, clr color.Color, cx, cy, r float64) {
	ebitenutil.DrawCircle(dst, cx, cy, r, clr)
}

// 新建圖片
func (d *Driver) NewImage(w, h int) *Image {
	return ebiten.NewImage(w, h)
}

// 取得圖片
func (d *Driver) GetImage(path string) (*Image, error) {
	var (
		err error
		b   []byte
		img image.Image
	)
	if b, err = d.assets.ReadFile(path); err != nil {
		return nil, err
	}
	buf := bytes.NewReader(b)
	if img, _, err = image.Decode(buf); err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

// 設定屏幕框線
func (d *Driver) SetBorders(padding float64) []BaseLine {
	res := make([]BaseLine, 4)
	// 依序 上下左右
	res = append(res, BaseLine{
		X1: padding,
		Y1: padding,
		X2: d.SW - padding,
		Y2: padding,
	})
	res = append(res, BaseLine{
		X1: padding,
		Y1: d.SH - padding,
		X2: d.SW - padding,
		Y2: d.SH - padding,
	})
	res = append(res, BaseLine{
		X1: padding,
		Y1: padding,
		X2: padding,
		Y2: d.SH - padding,
	})
	res = append(res, BaseLine{
		X1: d.SW - padding,
		Y1: padding,
		X2: d.SW - padding,
		Y2: d.SH - padding,
	})
	for k, v := range res {
		res[k].PointA = cp.Vector{X: v.X1, Y: v.Y1}
		res[k].PointB = cp.Vector{X: v.X2, Y: v.Y2}
	}
	return res
}
