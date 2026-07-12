package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Paddle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

type Game struct {
	Player Paddle
}

func (g *Game) Update() error {
	// Aqui ficará toda a lógica do jogo.
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 25,
		G: 25,
		B: 35,
		A: 255,
	})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.Player.Width, g.Player.Height)
	op.GeoM.Translate(g.Player.X, g.Player.Y)

	screen.DrawImage(whiteImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

var whiteImage *ebiten.Image

func init() {
	whiteImage = ebiten.NewImage(1, 1)
	whiteImage.Fill(color.White)
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Pong em Go")

	game := &Game{
		Player: Paddle{
			X:      30,
			Y:      250,
			Width:  20,
			Height: 100,
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
