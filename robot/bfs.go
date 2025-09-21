package robot

import (
	"fmt"

	"github.com/JMRodriguez-work/ia_tp2/utils"
)

/*
	Para la búsqueda BFS como es un algoritmo de búsqueda NO INFORMADO explora el espacio de estados nivel por nivel,
	garantizando encontrar la solución óptima

	- Usamos una cola FIFO para mantener los nodos por explorar
	- Explora todos los nodos a distancia 1, despues TODOS a distancia 2, etc.
	- Marcamos nodos visitados para evitar ciclos infinitos
	- Puede ser bastante lento en problemas complejos
*/

func (r *Robot) BFS(start, target utils.Position, maxSteps int) *utils.Node {
	queue := []*utils.Node{{Pos: start, Parent: nil}}
	visited := make(map[string]bool)

	serialize := func(pos utils.Position) string {
		return fmt.Sprintf("%.2f_%.2f_%.2f", pos.X, pos.Y, pos.Z)
	}

	visited[serialize(start)] = true
	nodesProcessed := 0

	for len(queue) > 0 && nodesProcessed < maxSteps {
		current := queue[0]
		queue = queue[1:]
		nodesProcessed++

		if r.CheckPosition(current.Pos, target) {
			fmt.Println("Posición encontrada con BFS")
			return current
		}

		for _, neighbor := range r.GenerateNeighbours(current.Pos) {
			key := serialize(neighbor)
			if !visited[key] {
				visited[key] = true
				queue = append(queue, &utils.Node{Pos: neighbor, Parent: current})
			}
		}
	}

	fmt.Println("BFS: No se encontró la posición dentro del límite de pasos")
	return nil
}

func ReconstructPath(node *utils.Node) []utils.Position {
	if node == nil {
		return nil
	}
	var path []utils.Position
	for n := node; n != nil; n = n.Parent {
		path = append([]utils.Position{n.Pos}, path...)
	}
	return path
}
