package estacionamiento
import (

    "fmt"
    "simulador/src/core/models"
    "sync"
    "github.com/hajimehoshi/ebiten/v2"
)

type GUI struct {
    ParkingLot    *models.ParkingLot
    CarsInMotion  []*models.Car
    CarsWaiting   []*models.Car
    ParkedCars    []*models.Car   
    Mutex         sync.Mutex
    windowWidth   int
    windowHeight  int
    processedCars int
    quit          chan struct{}
    shouldExit    bool
}

func NewGUI(parkingLot *models.ParkingLot) *GUI {

    return &GUI{
        ParkingLot:    parkingLot,
        CarsInMotion:  []*models.Car{},
        CarsWaiting:   []*models.Car{},
        ParkedCars:    []*models.Car{},
        windowWidth:   1000,
        windowHeight:  600,
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


func (gui *GUI) Run() {
    ebiten.SetWindowSize(gui.windowWidth, gui.windowHeight)
    ebiten.SetWindowTitle("Simulador de Estacionamiento")
    if err := ebiten.RunGame(gui); err != nil {
        panic(err)
    }
}


