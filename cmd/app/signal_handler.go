// cmd/app/signal_handler.go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "simulador/src/core/models"
)

// HandleInterruptSignal maneja señales de interrupción para cerrar el canal de carros y finalizar el programa.
func HandleInterruptSignal(carChannel chan *models.Car) {
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-quit
        fmt.Println("\nRecibida señal de interrupción. Cerrando simulación...")
        close(carChannel) 
    }()
}
