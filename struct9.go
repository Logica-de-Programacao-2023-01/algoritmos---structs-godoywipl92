package main

import (
	"fmt"
)

type Aluno struct {
	Nome   string
	Idade  int
	Notas  []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverNota(indice int) {
	if indice >= 0 && indice < len(a.Notas) {
		a.Notas = append(a.Notas[:indice], a.Notas[indice+1:]...)
	}
}

func (a *Aluno) CalcularMedia() float64 {
	total := 0.0
	for _, nota := range a.Notas {
		total += nota
	}
	return total / float64(len(a.Notas))
}

func (a *Aluno) ImprimirInformacoes() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Idade:", a.Idade)
	fmt.Println("Média:", a.CalcularMedia())
}

func main() {
	aluno := Aluno{
		Nome:  "João",
		Idade: 20,
		Notas: []float64{8.5, 7.9, 9.2},
	}

	aluno.ImprimirInformacoes()

	fmt.Println("Adicionando nota 6.8")
	aluno.AdicionarNota(6.8)

	fmt.Println("Removendo nota de índice 1")
	aluno.RemoverNota(1)

	aluno.ImprimirInformacoes()
}
