package main

import (

	"simulador/src/core/models"
	"simulador/src/core/services"
	"simulador/src/view"
     "simulador/src/core/observer"
)


type StartApp struct {}
func NewStartApp() *StartApp {
	return &StartApp{}
}
func (app *StartApp) StartSimulation() {
	parkingLot := models.NewParkingLot(ParkingCapacity)
	gui := view.NewGUI(parkingLot)
	carChannel := make(chan *models.Car, CarChannelBuffer)
	observerGenerator := observer.GetInstance()

    generators := make([]*services.CarGenerator, 3)
     
	for i := 0; i < 3; i++ {
	  generators[i] = services.NewCarGenerator(VelocityGenerationCar, carChannel)
	  observerGenerator.Subscribe(generators[i])
	} 

	for i := 0; i < 3; i++ {
		go generators[i].Generate()
	}

  
	go view.ProcessCarWorker(carChannel, gui)
	gui.Run()
}
