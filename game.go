package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	Player1 Paddle
	Player2 Paddle

	Ball Ball

	Player1Score int
	Player2Score int
}

func (g *Game) Update() error {
	g.Player1.Update()
	g.Player2.Update()

	g.Ball.Update()

	g.CheckCollisions()
	g.CheckGoals()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 25,
		G: 25,
		B: 35,
		A: 255,
	})

	g.Player1.Draw(screen)
	g.Player2.Draw(screen)

	g.Ball.Draw(screen)

	score := fmt.Sprintf("%d   -   %d", g.Player1Score, g.Player2Score)

	text.Draw(
		screen,
		score,
		basicfont.Face7x13,
		380,
		40,
		color.White,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) CheckCollisions() {
	if g.Ball.CollidesWith(&g.Player1) {
		g.Ball.X = g.Player1.X + g.Player1.Width
		g.BounceBall(&g.Player1)
	}

	if g.Ball.CollidesWith(&g.Player2) {
		g.Ball.X = g.Player2.X - g.Ball.Width
		g.BounceBall(&g.Player2)
	}
}

func (g *Game) BounceBall(p *Paddle) {
	ballCenter := g.Ball.Y + g.Ball.Height/2
	paddleCenter := p.Y + p.Height/2

	difference := ballCenter - paddleCenter
	normalized := difference / (p.Height / 2)

	g.Ball.VY = normalized * g.Ball.Speed
	g.Ball.VX *= -1
}

func (g *Game) CheckGoals() {

	if g.Ball.X <= 0 {
		g.Player2Score++
		g.Ball.Reset()
	}

	if g.Ball.X+g.Ball.Width >= ScreenWidth {
		g.Player1Score++
		g.Ball.Reset()
	}
}
