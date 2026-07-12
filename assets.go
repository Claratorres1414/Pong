package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var whiteImage *ebiten.Image

func init() {
	whiteImage = ebiten.NewImage(1, 1)
	whiteImage.Fill(color.White)
}
