// src/view/gui.go
package view

import (
    "fmt"
    "simulador/src/core/models"
    "sync"

    "github.com/hajimehoshi/ebiten/v2"
)

const (
    MaxQueueSize = 30 // Capacidad máxima de la cola de espera
)

// GUI representa la interfaz gráfica del simulador
type GUI struct {
    ParkingLot    *models.ParkingLot
    CarsInMotion  []*models.Car // Carros que están moviéndose (Searching o Exiting)
    CarsWaiting   []*models.Car // Carros que están en la cola de espera
    ParkedCars    []*models.Car // Carros estacionados
    Mutex         sync.Mutex
    windowWidth   int
    windowHeight  int
    completed     bool
    totalCars     int
    processedCars int
    quit          chan struct{}
    shouldExit    bool
}

// NewGUI crea una nueva instancia de GUI
func NewGUI(parkingLot *models.ParkingLot, totalCars int) *GUI {
    return &GUI{
        ParkingLot:    parkingLot,
        CarsInMotion:  []*models.Car{},
        CarsWaiting:   []*models.Car{},
        ParkedCars:    []*models.Car{},
        windowWidth:   1000, // Aumentar el ancho para mayor espacio visual
        windowHeight:  600,
        totalCars:     totalCars,
        processedCars: 0,
        quit:          make(chan struct{}),
        shouldExit:    false,
    }
}

// AddCar añade un carro al GUI, poniéndolo en movimiento
func (gui *GUI) AddCar(car *models.Car) {
    gui.CarsInMotion = append(gui.CarsInMotion, car)
    car.Estado = models.Searching
    fmt.Printf("Carro %d agregado a motion.\n", car.ID)
}

// Run inicia el bucle del juego usando ebiten
func (gui *GUI) Run() {
    ebiten.SetWindowSize(gui.windowWidth, gui.windowHeight)
    ebiten.SetWindowTitle("Simulador de Estacionamiento")
    if err := ebiten.RunGame(gui); err != nil {
        panic(err)
    }
}

// Stop cierra el canal quit para detener el GUI
func (gui *GUI) Stop() {
    close(gui.quit)
}
