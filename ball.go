package main

import "github.com/hajimehoshi/ebiten/v2"

type Ball struct {
	X, Y float64

	Width, Height float64

	VX float64
	VY float64
}

func (b *Ball) Update() {
	b.X += b.VX
	b.Y += b.VY

	if b.Y <= 0 || b.Y+b.Height >= ScreenHeight {
		b.VY *= -1
	}

	if b.X <= 0 || b.X+b.Width >= ScreenWidth {
		b.VX *= -1
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(b.Width, b.Height)
	op.GeoM.Translate(b.X, b.Y)

	screen.DrawImage(whiteImage, op)
}
