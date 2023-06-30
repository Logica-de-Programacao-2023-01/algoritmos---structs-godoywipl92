package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func main() {
	c := Circulo{raio: 5.0}
	area := calcularArea(c)
	fmt.Println("Área do círculo:", area)
}

func calcularArea(c Circulo) float64 {
	return math.Pi * c.raio * c.raio
}
