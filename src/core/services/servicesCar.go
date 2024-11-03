package services

import (
    "time"
    "simulador/src/core/models"
    "math/rand"
    "fmt"
    "sync"
)

/*
Este nos permite generar carros con llegadas utilizando la distribucion 
de poison

para realizar eventos aleatorios en tiempo discreto.
*/

func GenerateCars(totalCars int, ch chan<- *models.Car, wg *sync.WaitGroup) {
    defer close(ch) 

    for i := 1; i <= totalCars; i++ {
        interArrival := rand.ExpFloat64() / 1 
        time.Sleep(time.Duration(interArrival * float64(time.Second)))

        car := models.NewCar(i)
        fmt.Printf("Carro %d generado.\n", car.ID)
        wg.Add(1) 
        ch <- car
    }
}
