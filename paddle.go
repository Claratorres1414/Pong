package main

import "github.com/hajimehoshi/ebiten/v2"

type Paddle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (p *Paddle) Update() {
	const speed = 5

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += speed
	}
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.Width, p.Height)
	op.GeoM.Translate(p.X, p.Y)

	screen.DrawImage(whiteImage, op)
}
