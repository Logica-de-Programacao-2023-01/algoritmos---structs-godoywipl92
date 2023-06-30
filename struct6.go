package main

import (
	"fmt"
	"time"
)

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionario) AumentarSalario(percentual float64) {
	f.Salario += f.Salario * (percentual / 100)
}

func (f *Funcionario) DiminuirSalario(percentual float64) {
	f.Salario -= f.Salario * (percentual / 100)
}

func (f *Funcionario) TempoServico() int {
	idadeInicioTrabalho := 18
	idadeAtual := time.Now().Year() - f.Idade
	tempoServico := idadeAtual - idadeInicioTrabalho
	return tempoServico
}

func main() {
	funcionario := Funcionario{
		Nome:    "João",
		Salario: 3000.0,
		Idade:   25,
	}

	fmt.Println("Salário inicial:", funcionario.Salario)
	funcionario.AumentarSalario(10)
	fmt.Println("Salário após aumento:", funcionario.Salario)
	funcionario.DiminuirSalario(5)
	fmt.Println("Salário após diminuição:", funcionario.Salario)

	tempoServico := funcionario.TempoServico()
	fmt.Println("Tempo de serviço:", tempoServico, "anos")
}
