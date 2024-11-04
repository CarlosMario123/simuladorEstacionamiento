//src/view/elements/city/
package city
import (
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawCityscape(screen *ebiten.Image) {
    buildingColors := []color.Color{
        color.RGBA{R: 100, G: 100, B: 100, A: 255},
        color.RGBA{R: 150, G: 150, B: 150, A: 255},
        color.RGBA{R: 120, G: 120, B: 120, A: 255},
    }
    buildingWidths := []int{80, 100, 60, 90, 70}
    buildingPositions := []int{10, 120, 230, 340, 450}

    for i, w := range buildingWidths {
        buildingColor := buildingColors[i%len(buildingColors)]
        ebitenutil.DrawRect(screen, float64(buildingPositions[i]), 10, float64(w), 150, buildingColor)
        drawWindows(screen, buildingPositions[i], 10, w, 150)
    }
}

func drawWindows(screen *ebiten.Image, x, y, buildingWidth, buildingHeight int) {
    windowWidth, windowHeight := 10, 10
    spacingX, spacingY := 5, 15

    for wx := x + 10; wx < x+buildingWidth-10; wx += windowWidth + spacingX {
        for wy := y + 20; wy < y+buildingHeight-10; wy += windowHeight + spacingY {
            ebitenutil.DrawRect(screen, float64(wx), float64(wy), float64(windowWidth), float64(windowHeight), color.White)
        }
    }
}
