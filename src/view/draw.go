
package view

import (
    "image/color"
    "simulador/src/core/models"
    "simulador/src/view/elements/city"
    "simulador/src/view/elements/parking"
    "simulador/src/view/elements/vehicle"
    "github.com/hajimehoshi/ebiten/v2"
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

func (gui *GUI) Draw(screen *ebiten.Image) {
    gui.Mutex.Lock()
    defer gui.Mutex.Unlock()


    screen.Fill(color.RGBA{R: 200, G: 200, B: 200, A: 255})


    city.DrawCityscape(screen, gui.windowWidth)


    city.DrawRoad(screen, 0, float64(RoadYEntry), float64(gui.windowWidth), float64(RoadHeight), color.RGBA{R: 50, G: 50, B: 50, A: 255}, color.White)

  
    city.DrawRoad(screen, 0, float64(RoadYExit), float64(gui.windowWidth), float64(RoadHeight), color.RGBA{R: 50, G: 50, B: 50, A: 255}, color.White)
    gui.QueueView.Draw(screen)

    parking.DrawParkingSpaces(screen, gui.ParkingLot, ParkingStartX, ParkingStartY, SpaceWidth, SpaceHeight, Spacing, SpacesPerRow)

    // Dibujar carros en la carretera de entrada 
    for _, car := range gui.CarsInMotion {
        if car.Estado == models.Searching {
            posX := 50 + float64(car.Position)*700
            posY := float64(RoadYEntry + 10)
            vehicle.DrawCar(screen, int(posX), int(posY), car.ID, car.Color)
        }
    }

  
    // Dibujar carros en la carretera de salida
    for _, car := range gui.CarsInMotion {
        if car.Estado == models.Exiting {
            posX := 50 + float64(car.Position)*700
            posY := float64(RoadYExit + 10)
            vehicle.DrawCar(screen, int(posX), int(posY), car.ID, car.Color)
        }
    }

    
    for _, car := range gui.ParkedCars {
        row := (car.ParkingSpaceID - 1) / SpacesPerRow
        col := (car.ParkingSpaceID - 1) % SpacesPerRow
        posX := ParkingStartX + col*(SpaceWidth + Spacing)
        posY := ParkingStartY + row*(SpaceHeight + Spacing)
        vehicle.DrawCar(screen, posX, posY, car.ID, car.Color)
    }
}