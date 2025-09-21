package robot

import (
	"fmt"
	"sort"

	"github.com/JMRodriguez-work/ia_tp2/utils"
)

/*
	Para la búsqueda A* como es un algoritmo de búsqueda INFORMADO que usa una función heurística para ir al objetivo
	de manera eficiente

	- Usamos una lista abierta ordenada por f(n) = g(n) + h(n)
	  - g(n): Costo real desde el inicio hasta el nodo n
		- h(n): Estimación heurística desde n hasta el objetivo
		- f(n): Estimación del costo total del camino que pasa por n
	
	- Siempre expande al nodo con menor f(n)
*/

func (r *Robot) AStar(start, target utils.Position, maxSteps int) *utils.Node {
	openList := []*utils.Node{{Pos: start, Parent: nil, G: 0, H: r.Distance(start, target), F: r.Distance(start, target)}}
	visited := make(map[string]bool)

	serialize := func(pos utils.Position) string {
		return fmt.Sprintf("%.2f_%.2f_%.2f", pos.X, pos.Y, pos.Z)
	}

	for step := 0; step < maxSteps && len(openList) > 0; step++ {
		sort.Slice(openList, func(i, j int) bool {
			return openList[i].F < openList[j].F
		})
		current := openList[0]
		openList = openList[1:]

		key := serialize(current.Pos)
		if visited[key] {
			continue
		}
		visited[key] = true

		if r.CheckPosition(current.Pos, target) {
			fmt.Println("Posición encontrada con A*")
			return current
		}

		for _, neighborPos := range r.GenerateNeighbours(current.Pos) {
			g := current.G + r.Delta
			h := r.Distance(neighborPos, target)
			f := g + h
			openList = append(openList, &utils.Node{Pos: neighborPos, Parent: current, G: g, H: h, F: f})
		}
	}

	fmt.Println("A*: No se encontró la posición dentro del límite de pasos")
	return nil
}
