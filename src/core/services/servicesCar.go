package services

import (
    "fmt"
    "simulador/src/core/models"
    "simulador/src/utils"
    "sync"
    "time"
)

type CarGenerator struct {
    Velocity     []float64
    CarChan      chan<- *models.Car
    stop         chan struct{}
    mu           sync.Mutex
    active       bool
}

func NewCarGenerator(velocity []float64, carChan chan<- *models.Car) *CarGenerator {
    return &CarGenerator{
        Velocity: velocity,
        CarChan:  carChan,
        stop:     make(chan struct{}),
        active:   true,  // Asumir activo al inicio
    }
}

func (cg *CarGenerator) Generate() {
    defer close(cg.CarChan)

    for {
        select {
        case <-cg.stop:
            return
        default:
            if cg.isActive() {
                randomDelay := utils.RandomDelay(cg.Velocity[0], cg.Velocity[1])
                time.Sleep(randomDelay)
                car := models.NewCar()
                fmt.Printf("Carro %d generado.\n", car.ID)
                cg.CarChan <- car
            } else {
                time.Sleep(500 * time.Millisecond) // Espera activa menos intensa cuando está inactivo
            }
        }
    }
}

func (cg *CarGenerator) Stop() {
    cg.mu.Lock()
    cg.active = false
    cg.mu.Unlock()
}

func (cg *CarGenerator) Active() {
    cg.mu.Lock()
    cg.active = true
    cg.mu.Unlock()
}

func (cg *CarGenerator) isActive() bool {
    cg.mu.Lock()
    defer cg.mu.Unlock()
    return cg.active
}
