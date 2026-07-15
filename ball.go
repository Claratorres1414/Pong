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

	if b.Y <= 0 {
		b.Y = 0
		b.VY *= -1
	}

	if b.Y+b.Height >= ScreenHeight {
		b.Y = ScreenHeight - b.Height
		b.VY *= -1
	}

	if b.X <= 0 {
		b.X = 0
		b.VX *= -1
	}

	if b.X+b.Width >= ScreenWidth {
		b.X = ScreenWidth - b.Width
		b.VX *= -1
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(b.Width, b.Height)
	op.GeoM.Translate(b.X, b.Y)

	screen.DrawImage(whiteImage, op)
}

func (b *Ball) CollidesWith(p *Paddle) bool {
	return b.X < p.X+p.Width &&
		b.X+b.Width > p.X &&
		b.Y < p.Y+p.Height &&
		b.Y+b.Height > p.Y
}
