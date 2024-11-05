package observer

import (
    "simulador/src/core/services"
)

type ObserveCar struct{
	Generators []*services.CarGenerator
}

func NewObserveCar(generators []*services.CarGenerator) *ObserveCar {
    return &ObserveCar{
        Generators: generators,
    }
}

func (oc *ObserveCar) Subscribe(generator *services.CarGenerator) {
    oc.Generators = append(oc.Generators, generator)
}

func (oc *ObserveCar) NotifyGenerate(){
	for _, generator := range oc.Generators {
		generator.Active()
	}
}

func (oc *ObserveCar) NotifyStop(){
	for _, generator := range oc.Generators {
        generator.Stop()
    }
}