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
	go services.ProcessCarWorker(carChannel, gui, &wg)

	gui.Run()
	wg.Wait()
}
