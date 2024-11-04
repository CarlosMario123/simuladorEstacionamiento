package main

import (

	"simulador/src/core/models"
	"simulador/src/core/services"
	"simulador/src/view"

	"sync"
)


type StartApp struct {}
func NewStartApp() *StartApp {
	return &StartApp{}
}
func (app *StartApp) StartSimulation() {
	parkingLot := models.NewParkingLot(ParkingCapacity)
	gui := view.NewGUI(parkingLot, TotalCars)
	carChannel := make(chan *models.Car, CarChannelBuffer)
	
	var wg sync.WaitGroup
	go services.GenerateCars(TotalCars, carChannel, &wg,VelocityGenerationCar)
	go services.ProcessCarWorker(carChannel, gui, &wg)
	gui.Run()
	wg.Wait()
}
