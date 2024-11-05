//src/view/elements/city/
package city

import (
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


func DrawRoad(screen *ebiten.Image, x, y, width, height float64, roadColor, lineColor color.Color) {
    ebitenutil.DrawRect(screen, x, y, width, height, roadColor)
    ebitenutil.DrawRect(screen, x, y+height/2-1, width, 2, lineColor)
}
