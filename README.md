# TP2 - Inteligencia Artificial

**Implementación en GO de algoritmos de búsqueda (BFS y A\*)**

Se implementan dos algoritmos de búsqueda aplicados al problema planteado: un **robot manipulador** que debe mover piezas de un motor hasta alvanzar una posición objetivo.

Tanto el algoritmo exhaustivo (BFS) como el heurístico (A\*) se aplican sobre un **espacio de estados** definido por las posiciones `(X, Y, Z)`que representan el movimiento y la rotación de la pieza en el espacio.

## Modos de ejecución

Se puede ejecutar el binario o correr el proyecto localmente con GO

### Para el binario dependiendo del SO

En la sección **[Releases](./releases/)** de este repositorio se encuentran los archivos, descargar el correspondiente y hacer doble click en el mismo.

### Requisitos locales

- Tener instalado [Go](https://go.dev/dl/)

#### Pasos

1. Clonar este repositorio o descargar los archivos:
   ```bash
   git clone https://github.com/JMRodriguez-work/ia_tp2
   cd ia_tp2
   ```
2. Ejecutar el programa
   ```go
   go run main.go
   ```
