// src/models/car.go
package models

import (
    "fmt"
    "math/rand"
    "time"
)

// EstadoCarro representa los diferentes estados de un carro.
type EstadoCarro int

const (
    Searching EstadoCarro = iota // Buscando estacionamiento
    Parked                       // Estacionado
    Exiting                      // Saliendo del estacionamiento
)

// Car representa un vehículo en el simulador.
type Car struct {
    ID              int
    ParkingDuration time.Duration
    ParkingEndTime  time.Time
    Position        float64 // Representa la posición en la carretera (0.0 a 1.0)
    Estado          EstadoCarro
    ParkingSpaceID  int
    AttemptCount    int
}

// NewCar crea una nueva instancia de Car con un tiempo de estacionamiento aleatorio entre 3 y 4 segundos.
func NewCar(id int) *Car {
    duration := time.Duration(3+rand.Intn(2)) * time.Second // Entre 3 y 4 segundos
    return &Car{
        ID:              id,
        ParkingDuration: duration,
        Position:        0.0,
        Estado:          Searching,
        ParkingSpaceID:  -1,
        AttemptCount:    0,
    }
}

// Move incrementa la posición del carro hacia el estacionamiento.
func (c *Car) Move() {
    c.Position += 0.005 // Incremento reducido para movimiento más lento
    if c.Position > 1.0 {
        c.Position = 1.0
    }
    fmt.Printf("Carro %d moviéndose a posición %.2f\n", c.ID, c.Position)
}

// MoveExit decrementa la posición del carro para salir del estacionamiento.
func (c *Car) MoveExit() {
    c.Position -= 0.005 // Movimiento hacia la izquierda para salir
    if c.Position < 0.0 {
        c.Position = 0.0
    }
    fmt.Printf("Carro %d saliendo a posición %.2f\n", c.ID, c.Position)
}
