//src/view/elements/vehicle/
package vehicle

import (
    "fmt"
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
func DrawCar(screen *ebiten.Image, x, y, id int, carColor color.Color) {
    ebitenutil.DrawRect(screen, float64(x), float64(y), 60, 20, carColor)
    // Ventanas del carro
    DrawWindowCar(screen, x, y)
    drawCircle(screen, x+10, y+18, 5, color.Black)
    drawCircle(screen, x+50, y+18, 5, color.Black)
    drawText(screen, fmt.Sprintf("ID: %d", id), x+20, y-10)
}