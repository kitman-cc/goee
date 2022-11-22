package kitman

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v3"
)

// 類別別名
type (
	Image            = ebiten.Image
	DrawImageOptions = ebiten.DrawImageOptions
	Grid             = ganim8.Grid
)

// 常數別名
const (
	FilterNearest = ebiten.FilterNearest
	FilterLinear  = ebiten.FilterLinear
)

// 啟動器指針
var driver *Driver

// 界面
type Game interface {
	Setup(driver *Driver) error
	Layout(ow, oh int) (sw, sh int)
	Update() error
	Draw(screen *Image)
}

// 啟動參數
type RunGameOptions struct {
	Assets *embed.FS // 嵌入素材的檔案系統
	Title  string    // 標題
	Width  int       // 屏幕寬
	Height int       // 屏幕高
}

// 啟動遊戲
func RunGame(game Game, opts *RunGameOptions) error {
	driver = &Driver{
		assets: opts.Assets,
		SW:     float64(opts.Width),
		SH:     float64(opts.Height),
		HSW:    float64(opts.Width) / 2,
		HSH:    float64(opts.Height) / 2,
	}
	if err := game.Setup(driver); err != nil {
		log.Fatalln(err.Error())
	}

	ebiten.SetWindowTitle(opts.Title)
	ebiten.SetWindowSize(opts.Width, opts.Height)

	return ebiten.RunGame(game)
}

// 設定動畫繪製參數
func DrawAnimationOpts(x float64, y float64, args ...float64) *ganim8.DrawOptions {
	return ganim8.DrawOpts(x, y, args...)
}
