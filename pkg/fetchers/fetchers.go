package fetchers

import "github.com/juanvillacortac/entrenamiento-go/pkg/entities"

type ApiResponse interface {
	Transform() []entities.Song
}

type Params struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

type Fetcher func(query string) (ApiResponse, error)

var RegisteredFetchers []Fetcher = []Fetcher{
	FetchFromItunes,
}
