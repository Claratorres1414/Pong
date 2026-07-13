package main

import "github.com/hajimehoshi/ebiten/v2"

type Paddle struct {
	X, Y           float64
	Width, Height  float64
	UpKey, DownKey ebiten.Key
	speed          float64
}

func (p *Paddle) Update() {
	if ebiten.IsKeyPressed(p.DownKey) {
		p.Y += p.speed
	}

	if ebiten.IsKeyPressed(p.UpKey) {
		p.Y -= p.speed
	}

	if p.Y < 0 {
		p.Y = 0
	}

	if p.Y+p.Height > ScreenHeight {
		p.Y = ScreenHeight - p.Height
	}
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.Width, p.Height)
	op.GeoM.Translate(p.X, p.Y)

	screen.DrawImage(whiteImage, op)
}
