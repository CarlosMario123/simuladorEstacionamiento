package main

import (
    "simulador/src/controller/simulator"
)
func main(){
        simulator := simulator.NewSimulatorController()
        simulator.Run()
        
}