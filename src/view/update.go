package view
// Update maneja la actualización del estado del GUI
import (
    "fmt"
    "simulador/src/core/models"
    "time"

    "github.com/hajimehoshi/ebiten/v2"
)


func (gui *GUI) Update() error {
    select {
    case <-gui.quit:
        return nil
    default:
    }

    gui.Mutex.Lock()
    defer gui.Mutex.Unlock()

    // X cierra el programa
    if ebiten.IsKeyPressed(ebiten.KeyX) {
        fmt.Println("Salida manual activada por el usuario")
        close(gui.quit)
        gui.shouldExit = true
        return nil
    }

    updatedCarsInMotion := []*models.Car{}

    // Procesamiento de  carros en movimiento tanto salida como entrada
    for _, car := range gui.CarsInMotion {
        if car.Estado == models.Searching {
            car.Move()
            if car.Position >= 1.0 {
                spaceID, err := gui.ParkingLot.OccupySpace(car)
                if err != nil {
                    if len(gui.CarsWaiting) < MaxQueueSize {
                    
                        car.Estado = models.Waiting
                        gui.CarsWaiting = append(gui.CarsWaiting, car)
                        fmt.Printf("Carro %d no pudo estacionarse. Movido a la cola de espera.\n", car.ID)
                    } else {
                    
                        fmt.Printf("Carro %d no pudo estacionarse y la cola está llena. Carro descartado.\n", car.ID)
                        gui.processedCars++
                    }
                } else {
                   
                    car.Estado = models.Parked
                    car.ParkingSpaceID = spaceID
                    car.ParkingEndTime = time.Now().Add(car.ParkingDuration)
                    gui.ParkedCars = append(gui.ParkedCars, car)
                    fmt.Printf("Carro %d estacionado en el espacio %d.\n", car.ID, spaceID)
                }
            } else {
             
                updatedCarsInMotion = append(updatedCarsInMotion, car)
            }
        } else if car.Estado == models.Exiting {
            car.MoveExit()
            if car.Position < -0.1 { 
                fmt.Printf("Carro %d ha salido de la pantalla.\n", car.ID)
                gui.processedCars++
            } else {
           
                updatedCarsInMotion = append(updatedCarsInMotion, car)
            }
        }
    }

    gui.CarsInMotion = updatedCarsInMotion

   
    gui.checkParkedCars()

    
    if gui.processedCars >= gui.totalCars && !gui.completed {
        fmt.Println("Todos los carros han sido procesados. Saliendo.")
        gui.completed = true
        close(gui.quit)
        gui.shouldExit = true
        return nil
    }

    return nil
}

func (gui *GUI) checkParkedCars() {
    currentTime := time.Now()
    updatedParked := []*models.Car{}
    for _, car := range gui.ParkedCars {
        if currentTime.After(car.ParkingEndTime) {
            err := gui.ParkingLot.ReleaseSpace(car)
            if err == nil {
                car.Estado = models.Exiting
                car.Position = 1.0 
                gui.CarsInMotion = append(gui.CarsInMotion, car)
                fmt.Printf("Carro %d marcado para salir.\n", car.ID)


                if len(gui.CarsWaiting) > 0 {
                    waitingCar := gui.CarsWaiting[0]
                    gui.CarsWaiting = gui.CarsWaiting[1:]
                    gui.CarsInMotion = append(gui.CarsInMotion, waitingCar)
                    waitingCar.Estado = models.Searching
                    waitingCar.ResetParkingAttempt()
                    fmt.Printf("Carro %d reintentando estacionarse desde la cola de espera.\n", waitingCar.ID)
                }
            } else {
                updatedParked = append(updatedParked, car)
            }
        } else {
            updatedParked = append(updatedParked, car)
        }
    }
    gui.ParkedCars = updatedParked
}
