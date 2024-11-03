// src/view/elements/queue_view.go
package elements

import (
    "image/color"
    "simulador/src/core/models"
    "simulador/src/view/elements/vehicle"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type QueueView struct {
    queue       *models.Queue // Referencia al modelo de Queue
    startX      float64
    startY      float64
    width       float64
    height      float64
    slotHeight  int
    spacing     float64
}

// NewQueueView crea una nueva instancia de QueueView
func NewQueueView(queue *models.Queue, startX, startY, width, height float64) *QueueView {
    return &QueueView{
        queue:      queue,
        startX:     startX,
        startY:     startY,
        width:      width,
        height:     height,
        slotHeight: 40,
        spacing:    10.0,
    }
}

// Draw renderiza la cola de espera en la pantalla
func (qv *QueueView) Draw(screen *ebiten.Image) {
    // Dibujar el fondo de la cola de espera
    ebitenutil.DrawRect(screen, qv.startX, qv.startY, qv.width, qv.height, color.RGBA{R: 50, G: 50, B: 50, A: 255})
    ebitenutil.DrawRect(screen, qv.startX, qv.startY+qv.height/2-1, qv.width, 2, color.White)

    // Dibujar cada carro en la cola
    cars := qv.queue.GetCars()
    for i, car := range cars {
        posX := int(qv.startX + 10) // Margen dentro de la cola
        posY := int(qv.startY + qv.spacing + float64(i)*(float64(qv.slotHeight) + qv.spacing))

        // Detener dibujo si excede el espacio de la cola
        if float64(posY)+float64(qv.slotHeight) > qv.startY+qv.height {
            break
        }
        vehicle.DrawCar(screen, posX, posY, car.ID, car.Color)
    }
}
