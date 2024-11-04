//src/view/elements/vehicle/
package vehicle

import (
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawWindowCar(screen *ebiten.Image, x, y int) {
    // Ventana izquierda
    ebitenutil.DrawRect(screen, float64(x+10), float64(y+5), 15, 10, color.RGBA{R: 173, G: 216, B: 230, A: 255})
    
    // Ventana derecha
    ebitenutil.DrawRect(screen, float64(x+35), float64(y+5), 15, 10, color.RGBA{R: 173, G: 216, B: 230, A: 255})
}
