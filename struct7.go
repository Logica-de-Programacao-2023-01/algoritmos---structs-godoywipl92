package main

import "fmt"

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func (a *Animal) ModificarSom(novoSom string) {
	a.Som = novoSom
}

func (a Animal) ImprimirInformacoes() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Espécie:", a.Especie)
	fmt.Println("Idade:", a.Idade)
	fmt.Println("Som:", a.Som)
}

func main() {
	animal := Animal{
		Nome:    "Rex",
		Especie: "Cachorro",
		Idade:   5,
		Som:     "Au Au",
	}

	animal.ImprimirInformacoes()

	fmt.Println("---")

	animal.ModificarSom("Miau")

	animal.ImprimirInformacoes()
}
