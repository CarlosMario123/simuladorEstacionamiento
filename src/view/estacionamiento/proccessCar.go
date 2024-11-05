package estacionamiento

import (
	"fmt"
	"simulador/src/core/models"
)
func ProcessCarWorker(carChannel <-chan *models.Car, gui *GUI) {
    for car := range carChannel {
        gui.AddCar(car)
        fmt.Println("Carro ", car.ID, " procesado.")
    }
  
}
