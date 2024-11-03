
package view

import (
    "image/color"
    "simulador/src/core/models"
    "simulador/src/view/elements/city"
    "simulador/src/view/elements/parking"
    
    "simulador/src/view/elements/vehicle"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
    RoadYEntry     = 200
    RoadYExit      = 300
    RoadHeight     = 60
    QueueStartX    = 800
    QueueStartY    = 350.0
    QueueWidth     = 180.0
    QueueHeight    = 200.0
    ParkingStartX  = 50
    ParkingStartY  = 400
    SpaceWidth     = 100
    SpaceHeight    = 40
    Spacing        = 10
    SpacesPerRow   = 5
)

// Draw renderiza todos los elementos gráficos en pantalla
func (gui *GUI) Draw(screen *ebiten.Image) {
    gui.Mutex.Lock()
    defer gui.Mutex.Unlock()

    // Fondo de pantalla
    screen.Fill(color.RGBA{R: 200, G: 200, B: 200, A: 255})

    // Dibujar la ciudad
    city.DrawCityscape(screen, gui.windowWidth)

    // Carretera de entrada
    ebitenutil.DrawRect(screen, 0, float64(RoadYEntry), float64(gui.windowWidth), float64(RoadHeight), color.RGBA{R: 50, G: 50, B: 50, A: 255})
    ebitenutil.DrawRect(screen, 0, float64(RoadYEntry+RoadHeight/2-1), float64(gui.windowWidth), 2, color.White)

    // Carretera de salida
    ebitenutil.DrawRect(screen, 0, float64(RoadYExit), float64(gui.windowWidth), float64(RoadHeight), color.RGBA{R: 50, G: 50, B: 50, A: 255})
    ebitenutil.DrawRect(screen, 0, float64(RoadYExit+RoadHeight/2-1), float64(gui.windowWidth), 2, color.White)

    // Dibujar la cola de espera
    gui.QueueView.Draw(screen)

    // Dibujar el estacionamiento usando el paquete `parking`
    parking.DrawParkingSpaces(screen, gui.ParkingLot, ParkingStartX, ParkingStartY, SpaceWidth, SpaceHeight, Spacing, SpacesPerRow)

    // Dibujar carros en la carretera de entrada (estado Searching)
    for _, car := range gui.CarsInMotion {
        if car.Estado == models.Searching {
            posX := 50 + float64(car.Position)*700
            posY := float64(RoadYEntry + 10)
            vehicle.DrawCar(screen, int(posX), int(posY), car.ID, car.Color)
        }
    }

    // Dibujar carros en la carretera de salida (estado Exiting)
    for _, car := range gui.CarsInMotion {
        if car.Estado == models.Exiting {
            posX := 50 + float64(car.Position)*700
            posY := float64(RoadYExit + 10)
            vehicle.DrawCar(screen, int(posX), int(posY), car.ID, car.Color)
        }
    }

    // Dibujar carros estacionados en el estacionamiento
    for _, car := range gui.ParkedCars {
        row := (car.ParkingSpaceID - 1) / SpacesPerRow
        col := (car.ParkingSpaceID - 1) % SpacesPerRow
        posX := ParkingStartX + col*(SpaceWidth + Spacing)
        posY := ParkingStartY + row*(SpaceHeight + Spacing)
        vehicle.DrawCar(screen, posX, posY, car.ID, car.Color)
    }
}
