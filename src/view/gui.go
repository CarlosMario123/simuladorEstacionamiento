// src/view/gui.go
package view

import (
    "fmt"
    "image/color"
    "simulador/src/models"
    "sync"
    "time"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/text"
    "golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
)

// GUI implementa la interfaz ebiten.Game.
type GUI struct {
    ParkingLot    *models.ParkingLot
    CarsInMotion  []*models.Car
    CarsPassing   []*models.Car
    ParkedCars    []*models.Car
    Mutex         sync.Mutex
    fontFace      font.Face
    windowWidth   int
    windowHeight  int
    completed     bool
    totalCars     int
    processedCars int
}

// NewGUI crea una nueva instancia de GUI.
func NewGUI(parkingLot *models.ParkingLot, totalCars int) *GUI {
    gui := &GUI{
        ParkingLot:    parkingLot,
        CarsInMotion:  []*models.Car{},
        CarsPassing:   []*models.Car{},
        ParkedCars:    []*models.Car{},
        fontFace:      basicfont.Face7x13,
        windowWidth:   800,
        windowHeight:  600,
        totalCars:     totalCars,
        processedCars: 0,
    }
    return gui
}

// Update actualiza el estado del juego.
func (gui *GUI) Update() error {
    gui.Mutex.Lock()
    defer gui.Mutex.Unlock()

    for _, car := range gui.CarsInMotion {
        if car.Estado == models.Searching {
            car.Move()
            if car.Position >= 1.0 {
                spaceID, err := gui.ParkingLot.OccupySpace(car)
                if err != nil {
                    gui.CarsPassing = append(gui.CarsPassing, car)
                } else {
                    car.Estado = models.Parked
                    car.ParkingSpaceID = spaceID
                    car.ParkingEndTime = time.Now().Add(car.ParkingDuration)
                    gui.ParkedCars = append(gui.ParkedCars, car)
                    gui.processedCars++
                }
            }
        } else if car.Estado == models.Exiting {
            car.MoveExit()
            if car.Position <= 0.0 {
                gui.removeExitedCar(car.ID)
                gui.processedCars++
            }
        }
    }

    gui.removeParkedFromMotion()
    gui.handlePassingCars()
    gui.checkParkedCars()

    if gui.processedCars >= gui.totalCars && !gui.completed {
        gui.completed = true
    }

    return nil
}

// checkParkedCars verifica si los carros estacionados deben salir.
func (gui *GUI) checkParkedCars() {
    currentTime := time.Now()
    updatedParked := []*models.Car{}
    for _, car := range gui.ParkedCars {
        if currentTime.After(car.ParkingEndTime) {
            err := gui.ParkingLot.ReleaseSpace(car)
            if err == nil {
                car.Estado = models.Exiting
            }
        } else {
            updatedParked = append(updatedParked, car)
        }
    }
    gui.ParkedCars = updatedParked
}

// removeParkedFromMotion elimina carros estacionados de la lista de carros en movimiento.
func (gui *GUI) removeParkedFromMotion() {
    updatedCars := []*models.Car{}
    for _, car := range gui.CarsInMotion {
        if car.Estado != models.Parked {
            updatedCars = append(updatedCars, car)
        }
    }
    gui.CarsInMotion = updatedCars
}

// handlePassingCars maneja la lógica de reintentos para carros pasando.
func (gui *GUI) handlePassingCars() {
    updatedPassing := []*models.Car{}
    for _, car := range gui.CarsPassing {
        go func(c *models.Car) {
            time.Sleep(2 * time.Second)
            gui.Mutex.Lock()
            c.AttemptCount++
            gui.Mutex.Unlock()

            if c.AttemptCount <= 3 {
                gui.AddCar(c)
            } else {
                gui.Mutex.Lock()
                gui.processedCars++
                gui.Mutex.Unlock()
            }
        }(car)
    }
    gui.CarsPassing = updatedPassing
}

// AddCar añade un carro a la lista de carros en movimiento.
func (gui *GUI) AddCar(car *models.Car) {
    gui.Mutex.Lock()
    defer gui.Mutex.Unlock()
    gui.CarsInMotion = append(gui.CarsInMotion, car)
    car.Estado = models.Searching
}

// removeExitedCar remueve un carro de la lista de carros estacionados después de salir.
func (gui *GUI) removeExitedCar(carID int) {
    for i, car := range gui.ParkedCars {
        if car.ID == carID {
            gui.ParkedCars = append(gui.ParkedCars[:i], gui.ParkedCars[i+1:]...)
            break
        }
    }
}

// Draw renderiza la interfaz gráfica.
func (gui *GUI) Draw(screen *ebiten.Image) {
    gui.Mutex.Lock()
    defer gui.Mutex.Unlock()

    screen.Fill(color.RGBA{R: 200, G: 200, B: 200, A: 255})

    startX := 50
    startY := 50
    spaceWidth := 50
    spaceHeight := 50
    spacing := 10
    spacesPerRow := 5

    for _, space := range gui.ParkingLot.Spaces {
        x := startX + (space.ID-1)%spacesPerRow*(spaceWidth+spacing)
        y := startY + (space.ID-1)/spacesPerRow*(spaceHeight+spacing)

        var clr color.Color
        if space.IsOccupied {
            clr = color.RGBA{R: 255, G: 0, B: 0, A: 255}
        } else {
            clr = color.RGBA{G: 255, B: 0, A: 255}
        }

        ebitenutil.DrawRect(screen, float64(x), float64(y), float64(spaceWidth), float64(spaceHeight), clr)
    }

    for _, car := range gui.CarsInMotion {
        posX := 50 + float64(car.Position)*700
        posY := 320.0
        if posX > float64(gui.windowWidth)-30 {
            posX = float64(gui.windowWidth) - 30
        }
        ebitenutil.DrawRect(screen, posX, posY, 20, 10, color.RGBA{R: 255, G: 255, B: 0, A: 255})
        drawText(screen, fmt.Sprintf("ID: %d", car.ID), int(posX), int(posY-10))
    }

    for _, car := range gui.CarsPassing {
        posX := 50 + float64(car.Position)*700
        posY := 420.0
        ebitenutil.DrawRect(screen, posX, posY, 20, 10, color.RGBA{R: 0, G: 255, B: 255, A: 255})
        drawText(screen, fmt.Sprintf("ID: %d", car.ID), int(posX), int(posY-10))
    }

    for _, car := range gui.ParkedCars {
        row := (car.ParkingSpaceID - 1) / spacesPerRow
        col := (car.ParkingSpaceID - 1) % spacesPerRow
        posX := 50 + float64(col)*(float64(spaceWidth)+float64(spacing))
        posY := 500 + float64(row)*(float64(spaceHeight)+float64(spacing))
        var clr color.Color
        if car.Estado == models.Parked {
            clr = color.RGBA{G: 255, B: 0, A: 255}
        } else if car.Estado == models.Exiting {
            clr = color.RGBA{R: 255, G: 0, B: 0, A: 255}
        }
        ebitenutil.DrawRect(screen, posX, posY, 20, 10, clr)
        drawText(screen, fmt.Sprintf("ID: %d", car.ID), int(posX), int(posY-10))
    }

    if gui.completed {
        finalMsg := "¡Todos los carros han sido procesados!"
        textX := gui.windowWidth/2 - len(finalMsg)*3
        textY := gui.windowHeight / 2
        drawText(screen, finalMsg, textX, textY)
    }
}

// drawText es una función auxiliar para dibujar texto en la pantalla.
func drawText(screen *ebiten.Image, label string, x, y int) {
    text.Draw(screen, label, basicfont.Face7x13, x, y, color.Black)
}

// Layout define el tamaño de la ventana.
func (gui *GUI) Layout(outsideWidth, outsideHeight int) (int, int) {
    return gui.windowWidth, gui.windowHeight
}

// Run inicia el bucle principal de Ebiten.
func (gui *GUI) Run() {
    ebiten.SetWindowSize(gui.windowWidth, gui.windowHeight)
    ebiten.SetWindowTitle("Simulador de Estacionamiento")

    if err := ebiten.RunGame(gui); err != nil {
        panic(err)
    }
}
