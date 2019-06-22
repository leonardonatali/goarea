package goarea

import "math"

//Pi é uma proporção numérica definida pela relação entre
//o perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

//Circ - Função responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect - Função responsável por calcular a área do retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Função privada
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
