// cmd/app/main.go
package main

import (
    "fmt"
    "math/rand"
    "os"
    "os/signal"
    "simulador/src/models"
    "simulador/src/view"
    "sync"
    "syscall"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    totalCars := 100 // Número total de carros a generar

    // Crear el estacionamiento con capacidad para 20 vehículos.
    parkingLot := models.NewParkingLot(20)

    // Crear la interfaz gráfica.
    gui := view.NewGUI(parkingLot, totalCars)

    // Crear un canal para recibir nuevos vehículos.
    carChannel := make(chan *models.Car, 200) // Aumentar el buffer para más carros en carretera

    var wg sync.WaitGroup

    // Iniciar la generación de vehículos en una goroutine separada.
    go generateCars(totalCars, carChannel, &wg)

    // Manejar los vehículos que llegan.
    go func() {
        for car := range carChannel {
            gui.AddCar(car)
            wg.Done() // Marcar como completado después de añadir
        }
    }()

    // Manejar señales de interrupción para cerrar la aplicación limpiamente.
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-quit
        fmt.Println("\nRecibida señal de interrupción. Cerrando simulación...")
        close(carChannel)
    }()

    // Iniciar la interfaz gráfica.
    gui.Run()

    // Esperar a que todas las goroutines terminen antes de salir.
    wg.Wait()

    fmt.Println("Simulación finalizada.")
}

// generateCars simula la creación de vehículos con llegadas según una distribución de Poisson.
func generateCars(totalCars int, ch chan<- *models.Car, wg *sync.WaitGroup) {
    for i := 1; i <= totalCars; i++ {
        interArrival := rand.ExpFloat64() / 1.0 // Tasa λ = 1 (ajustable)
        time.Sleep(time.Duration(interArrival * float64(time.Second)))

        car := models.NewCar(i)
        fmt.Printf("Carro %d generado.\n", car.ID)
        wg.Add(1) // Añadir al WaitGroup antes de enviar
        ch <- car
    }
}
