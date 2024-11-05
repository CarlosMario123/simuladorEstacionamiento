package observer

import (
    "sync"
    "simulador/src/core/services"
)


type SingletonGeneratorCar struct {
    observeCar *ObserveCar
}

var instance *SingletonGeneratorCar
var once sync.Once


func GetInstance() *SingletonGeneratorCar {
    once.Do(func() {
    
        instance = &SingletonGeneratorCar{
            observeCar: NewObserveCar([]*services.CarGenerator{}),
        }
    })
    return instance
}


func (s *SingletonGeneratorCar) Subscribe(generator *services.CarGenerator) {
    s.observeCar.Subscribe(generator)
}


func (s *SingletonGeneratorCar) NotifyGenerate() {
    s.observeCar.NotifyGenerate()
}

func (s *SingletonGeneratorCar) NotifyStop() {
    s.observeCar.NotifyStop()
}
