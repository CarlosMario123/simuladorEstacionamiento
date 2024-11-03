package view
// drawParkingSpaces dibuja los espacios de estacionamiento
import (
    "image/color"
    "simulador/src/core/models" 
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ models.ParkingLot // Referencia explícita para evitar importacion no usada


func (gui *GUI) drawParkingSpaces(screen *ebiten.Image) {
    startX, startY := 50, 350
    spaceWidth, spaceHeight := 100, 40
    spacing, spacesPerRow := 10, 5

    for i, space := range gui.ParkingLot.Spaces {
        x := startX + (i%spacesPerRow)*(spaceWidth+spacing)
        y := startY + (i/spacesPerRow)*(spaceHeight+spacing)
        clr := color.RGBA{G: 255, B: 0, A: 255}
        if space.IsOccupied {
            clr = color.RGBA{R: 255, G: 0, B: 0, A: 255}
        }
        ebitenutil.DrawRect(screen, float64(x), float64(y), float64(spaceWidth), float64(spaceHeight), clr)
    }
}
