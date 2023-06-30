package main

import (
	"fmt"
)

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func (f *Filme) AdicionarAvaliacao(avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func (f *Filme) RemoverAvaliacao(indice int) {
	if indice >= 0 && indice < len(f.Avaliacoes) {
		f.Avaliacoes = append(f.Avaliacoes[:indice], f.Avaliacoes[indice+1:]...)
	}
}

func (f *Filme) CalcularMediaAvaliacoes() float64 {
	total := 0
	for _, nota := range f.Avaliacoes {
		total += nota
	}
	return float64(total) / float64(len(f.Avaliacoes))
}

func (f *Filme) ImprimirInformacoes() {
	fmt.Println("Título:", f.Titulo)
	fmt.Println("Diretor:", f.Diretor)
	fmt.Println("Ano:", f.Ano)
	fmt.Println("Média de Avaliações:", f.CalcularMediaAvaliacoes())
}

func main() {
	filme := Filme{
		Titulo:     "O Poderoso Chefão",
		Diretor:    "Francis Ford Coppola",
		Ano:        1972,
		Avaliacoes: []int{9, 10, 8, 9, 9},
	}

	filme.ImprimirInformacoes()

	fmt.Println("Adicionando avaliação 7")
	filme.AdicionarAvaliacao(7)

	fmt.Println("Removendo avaliação de índice 1")
	filme.RemoverAvaliacao(1)

	filme.ImprimirInformacoes()
}
