package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Pong")

	game := &Game{
		Player1: Paddle{
			X:      30,
			Y:      250,
			Width:  20,
			Height: 100,

			UpKey:   ebiten.KeyW,
			DownKey: ebiten.KeyS,

			speed: 5.0,
		},

		Player2: Paddle{
			X:      ScreenWidth - 30 - 20,
			Y:      250,
			Width:  20,
			Height: 100,

			UpKey:   ebiten.KeyArrowUp,
			DownKey: ebiten.KeyArrowDown,

			speed: 5.0,
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
