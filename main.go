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

func (p *Paddle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.Width, p.Height)
	op.GeoM.Translate(p.X, p.Y)

	screen.DrawImage(whiteImage, op)
}

type Game struct {
	Player Paddle
}

func (g *Game) Update() error {
	const speed = 5

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.Y -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.Y += speed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 25,
		G: 25,
		B: 35,
		A: 255,
	})

	g.Player.Draw(screen)
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
