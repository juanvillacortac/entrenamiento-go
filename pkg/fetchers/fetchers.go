package fetchers

import "github.com/juanvillacortac/entrenamiento-go/pkg/entities"

type ApiResponse interface {
	Transform() []entities.Song
}

type Params struct {
	Name   string `form:"name" json:"name"`
	Artist string `form:"artist" json:"artist"`
	Album  string `form:"album" json:"album"`
}

type Fetcher func(params Params) (ApiResponse, error)

var RegisteredFetchers []Fetcher = []Fetcher{
	FetchFromItunes,
}

func RetrieveFromApis(params Params) ([]entities.Song, error) {
	songs := make([]entities.Song, 0)
	for _, fetcher := range RegisteredFetchers {
		response, err := fetcher(params)
		if err != nil {
			return songs, err
		}
		songs = append(songs, response.Transform()...)
	}
	return songs, nil
}
