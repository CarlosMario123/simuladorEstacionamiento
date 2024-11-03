// src/view/gui.go
package view

import (
    "fmt"
    "simulador/src/core/models"
    "simulador/src/view/elements"
    "sync"
    "github.com/hajimehoshi/ebiten/v2"
)

type GUI struct {
    ParkingLot    *models.ParkingLot
    CarsInMotion  []*models.Car
    CarsWaiting   []*models.Car
    ParkedCars    []*models.Car
    Queue         *models.Queue          
    QueueView     *elements.QueueView    
    Mutex         sync.Mutex
    windowWidth   int
    windowHeight  int
    completed     bool
    totalCars     int
    processedCars int
    quit          chan struct{}
    shouldExit    bool
}

const (
    MaxQueueSize = 30 
)

func NewGUI(parkingLot *models.ParkingLot, totalCars int) *GUI {
    queue := models.NewQueue(MaxQueueSize) 
    queueView := elements.NewQueueView(queue, QueueStartX, QueueStartY, QueueWidth, QueueHeight) // Inicializa `QueueView` con `Queue`

    return &GUI{
        ParkingLot:    parkingLot,
        CarsInMotion:  []*models.Car{},
        CarsWaiting:   []*models.Car{},
        ParkedCars:    []*models.Car{},
        Queue:         queue,        
        QueueView:     queueView,     
        windowWidth:   1000,
        windowHeight:  600,
        totalCars:     totalCars,
        processedCars: 0,
        quit:          make(chan struct{}),
        shouldExit:    false,
    }
}


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


func (gui *GUI) Stop() {
    close(gui.quit)
}

func (gui *GUI) Layout(outsideWidth, outsideHeight int) (int, int) {
    return gui.windowWidth, gui.windowHeight
}
