package main

import "fmt"

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func ViagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		return Viagem{}
	}

	viagemMaisCara := viagens[0]

	for _, viagem := range viagens {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{Origem: "São Paulo", Destino: "Rio de Janeiro", Data: "2023-07-10", Preco: 500.0},
		{Origem: "Paris", Destino: "Londres", Data: "2023-08-15", Preco: 800.0},
		{Origem: "Nova York", Destino: "Los Angeles", Data: "2023-09-20", Preco: 1000.0},
	}

	viagemMaisCara := ViagemMaisCara(viagens)

	fmt.Println("Viagem mais cara:")
	fmt.Println("Origem:", viagemMaisCara.Origem)
	fmt.Println("Destino:", viagemMaisCara.Destino)
	fmt.Println("Data:", viagemMaisCara.Data)
	fmt.Println("Preço:", viagemMaisCara.Preco)
}
