package models
import (
    "image/color"
    "math/rand"
    "time"
)

type CarState int

const (
    Searching CarState = iota
    Parked
    Exiting
    Waiting // Esto se se usa para la cola
)

type Car struct {
    ID              int
    Estado          CarState
    Position        float64 // 0.0 izquierda a 1.0 derecha
    ParkingSpaceID  int
    ParkingEndTime  time.Time
    ParkingDuration time.Duration
    AttemptCount    int
    Color           color.RGBA
    lastAttemptTime time.Time 
}

func NewCar(id int) *Car {
    return &Car{
        ID:              id,
        Estado:          Searching,
        Position:        0.0,
        ParkingDuration: time.Duration(rand.Intn(3)+3) * time.Second,

        Color: color.RGBA{
            R: uint8(rand.Intn(256)),
            G: uint8(rand.Intn(256)),
            B: uint8(rand.Intn(256)),
            A: 255,
        },
        lastAttemptTime: time.Now(),
    }
}
// avanza a derecha
func (c *Car) Move() {
    c.Position += 0.0090
    if c.Position > 1.0 {
        c.Position = 1.0
    }
}
func (c *Car) MoveExit() {
    c.Position -= 0.02 
    if c.Position < -0.1 { 
        c.Position = -0.1
    }
}

// determina si el carro debe intentar estacionarse
func (c *Car) ShouldAttemptParking() bool {

    return time.Since(c.lastAttemptTime) > 3*time.Second
}
// reinicia el tiempo del ultimo intento
func (c *Car) ResetParkingAttempt() {
    c.lastAttemptTime = time.Now()
}
