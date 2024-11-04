package services

import (
	"fmt"
	"simulador/src/core/models"
	"simulador/src/view"
	"sync"
)
func ProcessCarWorker(carChannel <-chan *models.Car, gui *view.GUI, wg *sync.WaitGroup) {
    for car := range carChannel {
        gui.AddCar(car)
        fmt.Println("Carro ", car.ID, " procesado.")
        wg.Done()
    }
  
}
