package robot

import "github.com/JMRodriguez-work/ia_tp2/utils"

type Robot struct {
	Delta float64 // Incremento de movimiento ΔH
}

func (r *Robot) CheckPosition(pos, target utils.Position) bool {
	const epsilon = 0.01
	return (pos.X >= target.X-epsilon && pos.X <= target.X+epsilon &&
		pos.Y >= target.Y-epsilon && pos.Y <= target.Y+epsilon &&
		pos.Z >= target.Z-epsilon && pos.Z <= target.Z+epsilon)
}

func (r *Robot) GenerateNeighbours(pos utils.Position) []utils.Position {
	return []utils.Position{
		{X: pos.X + r.Delta, Y: pos.Y, Z: pos.Z},
		{X: pos.X - r.Delta, Y: pos.Y, Z: pos.Z},
		{X: pos.X, Y: pos.Y, Z: pos.Z + r.Delta},
		{X: pos.X, Y: pos.Y, Z: pos.Z - r.Delta},
	}
}

func (r *Robot) Distance(a, b utils.Position) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return dx*dx + dy*dy + dz*dz
}
