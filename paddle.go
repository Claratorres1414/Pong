package main

import "github.com/hajimehoshi/ebiten/v2"

type Paddle struct {
	X, Y           float64
	Width, Height  float64
	UpKey, DownKey ebiten.Key
}

func (p *Paddle) Update() {
	const speed = 5

	if ebiten.IsKeyPressed(p.DownKey) {
		p.Y += speed
	}

	if ebiten.IsKeyPressed(p.UpKey) {
		p.Y -= speed
	}
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.Width, p.Height)
	op.GeoM.Translate(p.X, p.Y)

	screen.DrawImage(whiteImage, op)
}
