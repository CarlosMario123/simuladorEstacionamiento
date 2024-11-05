package view

import (
    "fmt"
    "simulador/src/core/models"
    "time"
    "simulador/src/core/observer"
)

func (gui *GUI) Update() error {
    gui.Mutex.Lock()
    defer gui.Mutex.Unlock()


    observer := observer.GetInstance()

    updatedCarsInMotion := []*models.Car{}

   
    for _, car := range gui.CarsInMotion {
        switch car.Estado {
        case models.Searching:
            car.Move() // Mueve el auto hacia adelante
            if car.Position >= 1.0 { // Si alcanza la posición de estacionamiento
                spaceID, err := gui.ParkingLot.OccupySpace(car) // Intenta ocupar un espacio
                if err != nil { 
                    observer.NotifyStop()
                    fmt.Println("No se pudo estacionar el carro. Esperando espacio...")
                    updatedCarsInMotion = append(updatedCarsInMotion, car) 
                } else { 
                    car.Estado = models.Parked
                    car.ParkingSpaceID = spaceID
                    car.ParkingEndTime = time.Now().Add(car.ParkingDuration)
                    gui.ParkedCars = append(gui.ParkedCars, car)
                    fmt.Printf("Carro %d estacionado en el espacio %d.\n", car.ID, spaceID)

                 
                    observer.NotifyGenerate()
                }
            } else {
                updatedCarsInMotion = append(updatedCarsInMotion, car)
            }
        case models.Exiting:
            car.MoveExit()
            if car.Position <= -0.1 {
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
    }

    return nil
}

func (gui *GUI) checkParkedCars() {
    currentTime := time.Now()
    updatedParked := []*models.Car{}
    observer := observer.GetInstance()

   
    for _, car := range gui.ParkedCars {
        if currentTime.After(car.ParkingEndTime) {
            err := gui.ParkingLot.ReleaseSpace(car) // Libera el espacio de estacionamiento
            if err == nil {
                car.Estado = models.Exiting
                car.Position = 1.0
                gui.CarsInMotion = append(gui.CarsInMotion, car)
                fmt.Printf("Carro %d marcado para salir.\n", car.ID)

        
                observer.NotifyGenerate()

                // Espera 
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
