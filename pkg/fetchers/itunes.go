package fetchers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	client "github.com/bozd4g/go-http-client"
	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
	"github.com/juanvillacortac/entrenamiento-go/pkg/utils"
)

type ItunesResponse struct {
	ResultCount int                    `json:"resultCount"`
	Results     []ItunesResponseResult `json:"results"`
}

type ItunesResponseResult struct {
	WrapperType             string    `json:"wrapperType"`
	Kind                    string    `json:"kind"`
	ArtistID                int64     `json:"artistId,omitempty"`
	CollectionID            int64     `json:"collectionId,omitempty"`
	TrackID                 int64     `json:"trackId"`
	ArtistName              string    `json:"artistName"`
	CollectionName          string    `json:"collectionName,omitempty"`
	TrackName               string    `json:"trackName"`
	CollectionCensoredName  string    `json:"collectionCensoredName,omitempty"`
	TrackCensoredName       string    `json:"trackCensoredName"`
	ArtistViewURL           string    `json:"artistViewUrl,omitempty"`
	CollectionViewURL       string    `json:"collectionViewUrl,omitempty"`
	TrackViewURL            string    `json:"trackViewUrl"`
	PreviewURL              string    `json:"previewUrl"`
	ArtworkURL30            string    `json:"artworkUrl30"`
	ArtworkURL60            string    `json:"artworkUrl60"`
	ArtworkURL100           string    `json:"artworkUrl100"`
	CollectionPrice         float64   `json:"collectionPrice"`
	TrackPrice              float64   `json:"trackPrice"`
	ReleaseDate             time.Time `json:"releaseDate"`
	CollectionExplicitness  string    `json:"collectionExplicitness"`
	TrackExplicitness       string    `json:"trackExplicitness"`
	DiscCount               int64     `json:"discCount,omitempty"`
	DiscNumber              int64     `json:"discNumber,omitempty"`
	TrackCount              int64     `json:"trackCount,omitempty"`
	TrackNumber             int64     `json:"trackNumber,omitempty"`
	TrackTimeMillis         int64     `json:"trackTimeMillis"`
	Country                 string    `json:"country"`
	Currency                string    `json:"currency"`
	PrimaryGenreName        string    `json:"primaryGenreName"`
	IsStreamable            bool      `json:"isStreamable,omitempty"`
	CollectionArtistID      int64     `json:"collectionArtistId,omitempty"`
	CollectionArtistViewURL string    `json:"collectionArtistViewUrl,omitempty"`
	TrackRentalPrice        float64   `json:"trackRentalPrice,omitempty"`
	CollectionHdPrice       float64   `json:"collectionHdPrice,omitempty"`
	TrackHdPrice            float64   `json:"trackHdPrice,omitempty"`
	TrackHdRentalPrice      float64   `json:"trackHdRentalPrice,omitempty"`
	ContentAdvisoryRating   string    `json:"contentAdvisoryRating,omitempty"`
	ShortDescription        string    `json:"shortDescription,omitempty"`
	LongDescription         string    `json:"longDescription,omitempty"`
	HasITunesExtras         bool      `json:"hasITunesExtras,omitempty"`
	CollectionArtistName    string    `json:"collectionArtistName,omitempty"`
}

func (r *ItunesResponseResult) GenerateId() string {
	return utils.Btoa(fmt.Sprintf("%s_%d", entities.Itunes, r.TrackID))
}

func (r *ItunesResponse) Transform() []entities.Song {
	songs := make([]entities.Song, 0)
	for _, result := range r.Results {
		songs = append(songs, entities.Song{
			ID:       result.GenerateId(),
			Name:     result.TrackName,
			Artist:   result.ArtistName,
			Duration: utils.FormatDurationFromMilliseconds(result.TrackTimeMillis),
			Album:    result.CollectionName,
			Artwork:  result.ArtworkURL100,
			Price:    fmt.Sprintf("%s %.2f", result.Currency, result.TrackPrice),
			Origin:   entities.Itunes,
		})
	}
	return songs
}

func FetchFromItunes(params Params) (ApiResponse, error) {
	httpClient := client.New("https://itunes.apple.com")

	request, err := httpClient.Get(fmt.Sprintf("/search?entity=song&term=%s", url.QueryEscape(params.StringWithDelimiter("-"))))
	if err != nil {
		return nil, err
	}

	httpResponse, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	var response ItunesResponse
	err = json.Unmarshal(httpResponse.Get().Body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
