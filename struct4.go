package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Titulo   string
	Artista  string
	Duracao  time.Duration
}

type Playlist struct {
	Nome     string
	Musicas  []Musica
}

func main() {
	p := Playlist{
		Nome: "Minha Playlist",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3 * time.Minute},
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4 * time.Minute},
			{Titulo: "Música 3", Artista: "Artista 3", Duracao: 2 * time.Minute},
		},
	}
	imprimirPlaylist(p)
}

func imprimirPlaylist(p Playlist) {
	fmt.Println("Playlist:", p.Nome)

	totalDuracao := calcularDuracaoTotal(p.Musicas)
	fmt.Println("Tempo total:", totalDuracao)

	fmt.Println("Músicas:")
	for _, musica := range p.Musicas {
		fmt.Printf("- %s, por %s, duração: %s\n", musica.Titulo, musica.Artista, musica.Duracao)
	}
}

func calcularDuracaoTotal(musicas []Musica) time.Duration {
	var duracaoTotal time.Duration
	for _, musica := range musicas {
		duracaoTotal += musica.Duracao
	}
	return duracaoTotal
}
