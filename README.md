# Simulador de Estacionamiento en Go usando Goroutines y Ebiten

## Descripción del Proyecto

Este proyecto consiste en un pequeño simulador de un estacionamiento, donde los automóviles verifican si hay un espacio disponible para estacionarse. La lógica es la siguiente:

- Si hay un espacio disponible, el automóvil se estaciona inmediatamente.
- Si todos los lugares del estacionamiento están ocupados, los automóviles se encolan y esperan hasta que un lugar se desocupe para poder estacionarse.

El simulador utiliza goroutines para gestionar la concurrencia entre los automóviles y la biblioteca Ebiten para la visualización gráfica.

## Estructura de carpetas del proyecto
```markdown
- src/
  - controller/
  - core/
  - utils/
  - view/
    - estacionamiento/
      - elements/
      - draw.go
      - gui.go
      - layout.go
      - processCar.go
      - update.go
- app/
```
### Descripciónes de algunas carpetas principales

### Controller
- **config.go**: Contiene la configuración y parámetros generales para el simulador.
### Core
#### Models
- **car.go**: Define el modelo de un automóvi asi como varios estados del vehiculo (searching,parked,waiting,existing)
- **parkingSpace.go**: Modelo de cada espacio de estacionamiento, indicando si está ocupado o disponible esta es la parte donde aplica el uso de semaforos.

#### Observer
- **observerCar.go**: Implementa el patrón de diseño Observer para gestionar la creacion de vehiculos.
- **singletonGeneratorCar.go**: Implementa un generador Singleton para crear instancias de automóviles, garantizando que haya una única instancia del generador en la aplicación.

#### Services
- **servicesCar.go**:Es un servicio que nos permitira la generacion de vehiculos

### View
- **estacionamiento/**: Contiene los archivos relacionados con la visualización del estacionamiento en la interfaz gráfica, utilizando Ebiten para representar gráficamente el estado del estacionamiento y los automóviles
- asi mismo se intento realizar lo mas modular posible.

## Instalación del Simulador de Estacionamiento

1. Clona el repositorio desde GitHub:
   ```bash
   git clone https://github.com/CarlosMario123/simuladorEstacionamiento.git
   cd simuladorEstacionamiento
   go mod tidy
   go run cmd/app/main.go

   ```
   

