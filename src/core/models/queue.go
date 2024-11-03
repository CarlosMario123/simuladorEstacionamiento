
package models


type Queue struct {
    cars    []*Car
    maxSize int
}

func NewQueue(maxSize int) *Queue {
    return &Queue{
        cars:    []*Car{},
        maxSize: maxSize,
    }
}

func (q *Queue) AddCar(car *Car) bool {
    if len(q.cars) < q.maxSize {
        q.cars = append(q.cars, car)
        return true
    }
    return false
}
func (q *Queue) RemoveCar() *Car {
    if len(q.cars) == 0 {
        return nil
    }
    car := q.cars[0]
    q.cars = q.cars[1:]
    return car
}

func (q *Queue) GetCars() []*Car {
    carsCopy := make([]*Car, len(q.cars))
    copy(carsCopy, q.cars)
    return carsCopy
}

func (q *Queue) IsEmpty() bool {
    return len(q.cars) == 0
}
func (q *Queue) IsFull() bool {
    return len(q.cars) >= q.maxSize
}
