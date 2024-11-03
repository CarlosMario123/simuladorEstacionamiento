package models
// ParkingSpace representa un espacio individual en el estacionamiento.
type ParkingSpace struct {
    ID           int
    IsOccupied   bool
    OccupyingCar *Car
}