package vehicle

import (
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
)
func drawCircle(screen *ebiten.Image, cx, cy, r int, clr color.Color) {
    for y := -r; y <= r; y++ {
        for x := -r; x <= r; x++ {
            if x*x+y*y <= r*r {
                screen.Set(cx+x, cy+y, clr)
            }
        }
    }
}
