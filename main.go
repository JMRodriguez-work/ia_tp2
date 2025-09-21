package main

import (
	"fmt"

	"github.com/JMRodriguez-work/ia_tp2/robot"
	"github.com/JMRodriguez-work/ia_tp2/utils"
)

func runExample(name string, start, target utils.Position, r robot.Robot, maxSteps int) {
	fmt.Printf("\n--- %s ---\n", name)
	fmt.Printf("Inicio: (%.2f, %.2f, %.2f)\n", start.X, start.Y, start.Z)
	fmt.Printf("Objetivo: (%.2f, %.2f, %.2f)\n", target.X, target.Y, target.Z)
	fmt.Printf("Delta: %.2f, Máx pasos: %d\n\n", r.Delta, maxSteps)

	fmt.Println("--- Búsqueda Exhaustiva (BFS) ---")
	nodeBFS := r.BFS(start, target, maxSteps)
	if nodeBFS != nil {
		pathBFS := robot.ReconstructPath(nodeBFS)
		for i, pos := range pathBFS {
			fmt.Printf("Paso %d: X=%.2f Y=%.2f Z=%.2f\n", i, pos.X, pos.Y, pos.Z)
		}
		fmt.Printf("Total pasos BFS: %d\n", len(pathBFS)-1)
	}

	fmt.Println("\n--- Búsqueda Heurística (A*) ---")
	nodeAStar := r.AStar(start, target, maxSteps)
	if nodeAStar != nil {
		pathAStar := robot.ReconstructPath(nodeAStar)
		for i, pos := range pathAStar {
			fmt.Printf("Paso %d: X=%.2f Y=%.2f Z=%.2f\n", i, pos.X, pos.Y, pos.Z)
		}
		fmt.Printf("Total pasos A*: %d\n", len(pathAStar)-1)
	}
}

func main() {
	runExample(
		"Ejemplo 1: Movimiento lineal simple",
		utils.Position{X: 0, Y: 0, Z: 0},
		utils.Position{X: 3, Y: 0, Z: 0},
		robot.Robot{Delta: 1.0},
		100,
	)

	runExample(
		"Ejemplo 2: Movimiento diagonal",
		utils.Position{X: 0, Y: 0, Z: 0},
		utils.Position{X: 2, Y: 0, Z: 3},
		robot.Robot{Delta: 1.0},
		100,
	)

	runExample(
		"Ejemplo 3: Movimientos con más detalle (Delta=0.5)",
		utils.Position{X: 0, Y: 0, Z: 0},
		utils.Position{X: 2, Y: 0, Z: 1.5},
		robot.Robot{Delta: 0.5},
		150,
	)

	fmt.Println("\nPresiona Enter para salir...")
	fmt.Scanln()
}
