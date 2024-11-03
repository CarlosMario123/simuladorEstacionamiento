package vehicle

import (
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/text"
    "golang.org/x/image/font/basicfont"
)
func drawText(screen *ebiten.Image, label string, x, y int) {
    text.Draw(screen, label, basicfont.Face7x13, x, y, color.Black)
}