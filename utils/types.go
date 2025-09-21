package utils

type Position struct {
	X float64
	Y float64
	Z float64
}

type Node struct {
	Pos    Position
	Parent *Node
	G      float64 // Es el costo acumulado para el algoritmo A*
	H      float64 // Para heurística
	F      float64 // F = G + H (A*)
}
