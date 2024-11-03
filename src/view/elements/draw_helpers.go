package elements

import (
    "fmt"
    "image/color"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/text"
    "golang.org/x/image/font/basicfont"
)

func DrawCar(screen *ebiten.Image, x, y, id int, carColor color.Color) {
    ebitenutil.DrawRect(screen, float64(x), float64(y), 60, 20, carColor)

    // Ventanas del carro
    ebitenutil.DrawRect(screen, float64(x+10), float64(y+5), 15, 10, color.RGBA{R: 173, G: 216, B: 230, A: 255})
    ebitenutil.DrawRect(screen, float64(x+35), float64(y+5), 15, 10, color.RGBA{R: 173, G: 216, B: 230, A: 255})

    // Llantas redondas
    drawCircle(screen, x+10, y+18, 5, color.Black)
    drawCircle(screen, x+50, y+18, 5, color.Black)

    // Texto ID del carro
    drawText(screen, fmt.Sprintf("ID: %d", id), x+20, y-10)
}

func drawText(screen *ebiten.Image, label string, x, y int) {
    text.Draw(screen, label, basicfont.Face7x13, x, y, color.Black)
}

func drawCircle(screen *ebiten.Image, cx, cy, r int, clr color.Color) {
    for y := -r; y <= r; y++ {
        for x := -r; x <= r; x++ {
            if x*x+y*y <= r*r {
                screen.Set(cx+x, cy+y, clr)
            }
        }
    }
}
