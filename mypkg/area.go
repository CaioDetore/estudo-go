package mypkg

import "math"

type Area struct{}

const Pi = 3.1416

// Circ é responsável por calcular a área da circunferência
func (a *Area) Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

func (a *Area) Rect(base, altura float64) float64 {
	return base * altura
}

func (a *Area) TrianguloEq(base, altura float64) float64 {
	return base * altura
}
