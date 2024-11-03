package main
import (
	"simulador/src/core/models"
	"simulador/src/core/services"
	"simulador/src/view"
	"sync"
)

func StartSimulation(){
	parkingLot := models.NewParkingLot(ParkingCapacity)
	gui := view.NewGUI(parkingLot, TotalCars)
	carChannel := make(chan *models.Car, CarChannelBuffer)
	var wg sync.WaitGroup

	go services.GenerateCars(TotalCars, carChannel, &wg)

	// creacion de carros y para la interfaz
	go func() {
		for car := range carChannel {
			gui.AddCar(car)
			wg.Done()
		}
	}()

	gui.Run()
	wg.Wait()
}
