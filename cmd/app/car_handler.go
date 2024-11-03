package main
import (
    "sync"
    "simulador/src/core/models"
    "simulador/src/view"
)
// HandleCars procesa los carros que llegan del canal y los añade a la GUI.
func HandleCars(carChannel chan *models.Car, wg *sync.WaitGroup, gui *view.GUI) {
    for car := range carChannel {
        gui.AddCar(car)
        wg.Done() // Marcar como completado después de añadir
    }
}