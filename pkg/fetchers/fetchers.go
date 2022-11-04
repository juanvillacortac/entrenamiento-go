package fetchers

import (
	"strings"

	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
	"github.com/juanvillacortac/entrenamiento-go/pkg/utils"
)

type ApiResponse interface {
	Transform() []entities.Song
}

type Params struct {
	Name   string `form:"name" json:"name"`
	Artist string `form:"artist" json:"artist"`
	Album  string `form:"album" json:"album"`
}

func (p Params) String() string {
	return p.StringWithDelimiter("-")
}

func (p Params) Hash() string {
	return utils.Btoa(p.String())
}

func (p Params) StringWithDelimiter(delimiter string) string {
	var terms []string
	for _, s := range []string{p.Name, p.Album, p.Artist} {
		if strings.TrimSpace(s) != "" {
			terms = append(terms, s)
		}
	}
	return strings.Join(terms, delimiter)
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
