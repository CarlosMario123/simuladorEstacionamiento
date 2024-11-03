package models

import (
    "errors"
    "fmt"
    "sync"
    "time"
)
// ParkingLot es muchos parking space
type ParkingLot struct {
    Capacity  int
    Spaces    []*ParkingSpace
    semaphore chan struct{} // Semaforo para controlar la capacidad
    mutex     sync.Mutex
}


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

// parte donde se aplica semafaros se uso para los espacios ocupados
func (p *ParkingLot) OccupySpace(car *Car) (int, error) {
    select {
    case p.semaphore <- struct{}{}:
        // espacio en el semaforo
    default:
        return -1, errors.New("no hay espacios disponibles")
    }

    // Logica para buscar espacios
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

    
    <-p.semaphore 
    return -1, errors.New("no hay espacios disponibles")
}

//liberamos un espacio de estacionamiento.
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
            <-p.semaphore //eres libre semaforo
            return nil
        }
    }

    return errors.New("espacio no encontrado para el carro")
}
