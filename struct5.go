package main

import (
	"fmt"
)
type Musica struct {
	Titulo   string
	Artista  string
	Duracao  string
}
type Playlist struct {
	Nome     string
	Musicas  []Musica
}
func main() {
	playlist1 := Playlist{
		Nome: "Playlist 1",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista 1", Duracao: "3m30s"},
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: "4m15s"},
		},
	}

	playlist2 := Playlist{
		Nome: "Playlist 2",
		Musicas: []Musica{
			{Titulo: "Música 3", Artista: "Artista 3", Duracao: "2m45s"},
			{Titulo: "Música 4", Artista: "Artista 4", Duracao: "3m10s"},
		},
	}

	playlist3 := Playlist{
		Nome: "Playlist 3",
		Musicas: []Musica{
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: "4m15s"},
			{Titulo: "Música 4", Artista: "Artista 4", Duracao: "3m10s"},
		},
	}

	playlists := []Playlist{playlist1, playlist2, playlist3}

	tituloBuscado := "Música 2"
	resultado := buscarPlaylistsPorTitulo(playlists, tituloBuscado)

	if len(resultado) > 0 {
		fmt.Printf("Playlists com a música '%s':\n", tituloBuscado)
		for _, playlist := range resultado {
			fmt.Println(playlist.Nome)
		}
	} else {
		fmt.Printf("Nenhuma playlist encontrada com a música '%s'.\n", tituloBuscado)
	}
}

func buscarPlaylistsPorTitulo(playlists []Playlist, titulo string) []Playlist {
	var playlistsEncontradas []Playlist
	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == titulo {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}
	return playlistsEncontradas
}
