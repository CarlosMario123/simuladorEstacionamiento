//src/view/elements/parking/
package parking

import (
    "image/color"
    "simulador/src/core/models"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// DrawParkingSpaces dibuja los espacios de estacionamiento en la pantalla
func DrawParkingSpaces(screen *ebiten.Image, parkingLot *models.ParkingLot, startX, startY, spaceWidth, spaceHeight, spacing, spacesPerRow int) {
    for i, space := range parkingLot.Spaces {
        x := startX + (i%spacesPerRow)*(spaceWidth+spacing)
        y := startY + (i/spacesPerRow)*(spaceHeight+spacing)

        clr := color.RGBA{G: 255, B: 0, A: 255} // Color por defecto para espacios vacíos
        if space.IsOccupied {
            clr = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Color para espacios ocupados
        }
        
        ebitenutil.DrawRect(screen, float64(x), float64(y), float64(spaceWidth), float64(spaceHeight), clr)
    }
}
