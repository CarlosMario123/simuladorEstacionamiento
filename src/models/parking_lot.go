// src/models/parking_lot.go
package models

import (
    "errors"
    "fmt"
    "sync"
    "time"
)

// ParkingSpace representa un espacio individual en el estacionamiento.
type ParkingSpace struct {
    ID           int
    IsOccupied   bool
    OccupyingCar *Car
}

// ParkingLot gestiona múltiples espacios de estacionamiento.
type ParkingLot struct {
    Capacity  int
    Spaces    []*ParkingSpace
    semaphore chan struct{} // Semáforo para controlar la capacidad
    mutex     sync.Mutex
}

// NewParkingLot crea un nuevo estacionamiento con una capacidad dada.
func NewParkingLot(capacity int) *ParkingLot {
    spaces := make([]*ParkingSpace, capacity)
    for i := 0; i < capacity; i++ {
        spaces[i] = &ParkingSpace{
            ID:           i + 1,
            IsOccupied:   false,
            OccupyingCar: nil,
        }
    }
    return &ParkingLot{
        Capacity:  capacity,
        Spaces:    spaces,
        semaphore: make(chan struct{}, capacity),
    }
}

// OccupySpace intenta ocupar un espacio de estacionamiento.
func (p *ParkingLot) OccupySpace(car *Car) (int, error) {
    select {
    case p.semaphore <- struct{}{}:
        // Adquirido un espacio en el semáforo
    default:
        // No hay espacio disponible
        return -1, errors.New("no hay espacios disponibles")
    }

    // Buscar un espacio disponible
    p.mutex.Lock()
    defer p.mutex.Unlock()

    for _, space := range p.Spaces {
        if !space.IsOccupied {
            space.IsOccupied = true
            space.OccupyingCar = car
            car.Estado = Parked
            car.ParkingSpaceID = space.ID
            car.ParkingEndTime = time.Now().Add(car.ParkingDuration)
            fmt.Printf("Carro %d ha ocupado el espacio %d\n", car.ID, space.ID)
            return space.ID, nil
        }
    }

    // Debería ser improbable llegar aquí debido al semáforo
    <-p.semaphore // Liberar el semáforo
    return -1, errors.New("no hay espacios disponibles")
}

// ReleaseSpace libera un espacio de estacionamiento.
func (p *ParkingLot) ReleaseSpace(car *Car) error {
    if car.Estado != Parked || car.ParkingSpaceID == -1 {
        return errors.New("el carro no está estacionado")
    }

    p.mutex.Lock()
    defer p.mutex.Unlock()

    for _, space := range p.Spaces {
        if space.ID == car.ParkingSpaceID && space.OccupyingCar.ID == car.ID {
            space.IsOccupied = false
            space.OccupyingCar = nil
            car.Estado = Exiting
            car.ParkingSpaceID = -1
            fmt.Printf("Carro %d ha liberado el espacio %d y está saliendo.\n", car.ID, space.ID)
            <-p.semaphore // Liberar el semáforo
            return nil
        }
    }

    return errors.New("espacio no encontrado para el carro")
}
