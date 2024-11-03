package main
import (
	"simulador/src/core/models"
	"simulador/src/core/services"
	"simulador/src/view"
	"sync"
	"fmt"
)

func StartSimulation(){
	
	parkingLot := models.NewParkingLot(ParkingCapacity)
	gui := view.NewGUI(parkingLot, TotalCars)

	// Crear canal para los carros y WaitGroup para la sincronización
	carChannel := make(chan *models.Car, CarChannelBuffer)
	var wg sync.WaitGroup

	
	go services.GenerateCars(TotalCars, carChannel, &wg)
   
	
	for i := 0; i < NumWorkers; i++ {
		go func(workerID int) {
			for car := range carChannel {
				gui.AddCar(car)
				wg.Done()      
			}
			fmt.Printf("Worker %d finalizó su procesamiento\n", workerID)
		}(i) 
	}

	
	gui.Run()

	wg.Wait()
}
