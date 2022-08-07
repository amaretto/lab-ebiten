package main

import (
	"image/color"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var t int64
var img *ebiten.Image

type Player struct {
	pos int
}

type Game struct{}

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	t1 := time.Now().Unix()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.3, 0.3)
	if t1-t < 1 {
		ebitenutil.DebugPrint(screen, "Hello, World!")
	} else if 1 <= t1-t && t1-t <= 2 {
		screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	} else if 3 <= t1-t && t1-t <= 4 {
		screen.DrawImage(img, op)
	} else {
		op.GeoM.Translate(float64(t1-t), float64(t1-t))
		screen.DrawImage(img, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	t = time.Now().Unix()
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
