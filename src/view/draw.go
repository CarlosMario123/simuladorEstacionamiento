// Draw se encarga de renderizar todos los componentes en pantalla
package view

import (
    "image/color"
    "simulador/src/view/elements"
     "simulador/src/core/models"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


func (gui *GUI) Draw(screen *ebiten.Image) {
    gui.Mutex.Lock()
    defer gui.Mutex.Unlock()

    screen.Fill(color.RGBA{R: 200, G: 200, B: 200, A: 255})


    elements.DrawCityscape(screen, gui.windowWidth)

    
    roadYEntry := 200
    roadYExit := 300

    roadHeight := 60

  
    ebitenutil.DrawRect(screen, 0, float64(roadYEntry), float64(gui.windowWidth), float64(roadHeight), color.RGBA{R: 50, G: 50, B: 50, A: 255})
    ebitenutil.DrawRect(screen, 0, float64(roadYEntry+roadHeight/2-1), float64(gui.windowWidth), 2, color.White)


    ebitenutil.DrawRect(screen, 0, float64(roadYExit), float64(gui.windowWidth), float64(roadHeight), color.RGBA{R: 50, G: 50, B: 50, A: 255})
    ebitenutil.DrawRect(screen, 0, float64(roadYExit+roadHeight/2-1), float64(gui.windowWidth), 2, color.White)

    // Variables para el dibujado de la cola de espera
    queueStartX := float64(gui.windowWidth - 200)
    queueStartY := 350.0                          
    queueWidth := 180.0                          
    queueHeight := 200.0                          
    ebitenutil.DrawRect(screen, queueStartX, queueStartY, queueWidth, queueHeight, color.RGBA{R: 50, G: 50, B: 50, A: 255})
    ebitenutil.DrawRect(screen, queueStartX, queueStartY+queueHeight/2-1, queueWidth, 2, color.White)

    //dibujo del estaciomaniento
    gui.drawParkingSpaces(screen)

    // carretera de entrada
    for _, car := range gui.CarsInMotion {
        if car.Estado == models.Searching {
            posX := 50 + float64(car.Position)*700
            posY := float64(roadYEntry + 10)
            elements.DrawCar(screen, int(posX), int(posY), car.ID, car.Color)
        }
    }

    // para la carretera de salida
    for _, car := range gui.CarsInMotion {
        if car.Estado == models.Exiting {
            posX := 50 + float64(car.Position)*700
            posY := float64(roadYExit + 10)
            elements.DrawCar(screen, int(posX), int(posY), car.ID, car.Color)
        }
    }

    // carros en la cola de espera para renderizado
    for i, car := range gui.CarsWaiting {
       
        slotHeight := 40
        spacing := 10.0 

     //posiciones entre carros
        posX := int(queueStartX + 10) 
        posY := int(queueStartY + spacing + float64(i)*(float64(slotHeight)+spacing)) 

        
        if float64(posY)+float64(slotHeight) > queueStartY+queueHeight {
            // Por si a futuro se desea realizar algo con la cola
            break
        }

        elements.DrawCar(screen, posX, posY, car.ID, car.Color)
    }

    // Logica para el renderizado de los carros estacionados
    for _, car := range gui.ParkedCars {
        row := (car.ParkingSpaceID - 1) / 5
        col := (car.ParkingSpaceID - 1) % 5
        posX := 50 + col*110
        posY := 350 + row*50 
        elements.DrawCar(screen, posX, posY, car.ID, car.Color)
    }
}

// tamaño de la pantalla
func (gui *GUI) Layout(outsideWidth, outsideHeight int) (int, int) {
    return gui.windowWidth, gui.windowHeight
}
