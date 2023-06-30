package main

import "fmt"

type Triangulo struct {
	Base   float64
	Altura float64
}

func main() {
	t := Triangulo{Base: 10.0, Altura: 5.0}
	area := calcularAreaTriangulo(t)
	fmt.Println("Área do triângulo:", area)
}

func calcularAreaTriangulo(t Triangulo) float64 {
	return (t.Base * t.Altura) / 2
}
