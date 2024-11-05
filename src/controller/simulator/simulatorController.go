
package simulator

import (
	"simulador/src/core/models"
	"simulador/src/core/observer"
	"simulador/src/core/services"
	"simulador/src/view/estacionamiento"
)
type SimulatorController struct{}


func NewSimulatorController() *SimulatorController {
    return &SimulatorController{}
}
func (sc *SimulatorController) Run() error {
	parkingLot := models.NewParkingLot(ParkingCapacity)
	gui := estacionamiento.NewGUI(parkingLot)
	carChannel := make(chan *models.Car, CarChannelBuffer)
	observerGenerator := observer.GetInstance()

    generators := make([]*services.CarGenerator, numGenerators)
    
	for i := 0; i < numGenerators; i++ {
	  generators[i] = services.NewCarGenerator(VelocityGenerationCar, carChannel)
	  observerGenerator.Subscribe(generators[i])
	} 

	for i := 0; i < numGenerators; i++ {
		go generators[i].Generate()
	}

  
	go estacionamiento.ProcessCarWorker(carChannel, gui)
	gui.Run()

	return nil
}
