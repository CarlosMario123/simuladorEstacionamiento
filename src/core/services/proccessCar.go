package services
import (

	"simulador/src/core/models"
	"simulador/src/view"
	"sync"
)
func ProcessCarWorker(carChannel <-chan *models.Car, gui *view.GUI, wg *sync.WaitGroup) {
    for car := range carChannel {
        gui.AddCar(car)
        wg.Done()
    }
}
