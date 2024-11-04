package services

import (
    "fmt"
    "simulador/src/core/models"
    "sync"
    "time"
    "simulador/src/utils"
)

func GenerateCars(totalCars int, ch chan<- *models.Car, wg *sync.WaitGroup, velocity []float64) {
    defer close(ch) 



    for i := 1; i <= totalCars; i++ {
    
        randomDelay := utils.RandomDelay(velocity[0],velocity[1])
        time.Sleep(randomDelay)

        car := models.NewCar(i)
        fmt.Printf("Carro %d generado.\n", car.ID)
        wg.Add(1) 
        ch <- car
    }
}
