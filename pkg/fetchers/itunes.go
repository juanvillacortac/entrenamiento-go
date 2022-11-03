package fetchers

import (
	"encoding/json"
	"fmt"
	"time"

	client "github.com/bozd4g/go-http-client"
)

type ItunesResponse struct {
	ResultCount int                    `json:"resultCount"`
	Results     []ItunesResponseResult `json:"results"`
}

type ItunesResponseResult struct {
	WrapperType             string    `json:"wrapperType"`
	Kind                    string    `json:"kind"`
	ArtistID                int       `json:"artistId,omitempty"`
	CollectionID            int       `json:"collectionId,omitempty"`
	TrackID                 int       `json:"trackId"`
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
	DiscCount               int       `json:"discCount,omitempty"`
	DiscNumber              int       `json:"discNumber,omitempty"`
	TrackCount              int       `json:"trackCount,omitempty"`
	TrackNumber             int       `json:"trackNumber,omitempty"`
	TrackTimeMillis         int       `json:"trackTimeMillis"`
	Country                 string    `json:"country"`
	Currency                string    `json:"currency"`
	PrimaryGenreName        string    `json:"primaryGenreName"`
	IsStreamable            bool      `json:"isStreamable,omitempty"`
	CollectionArtistID      int       `json:"collectionArtistId,omitempty"`
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

// func (r ItunesResponse) Transform() []entities.Song {
// }

func FetchFromItunes(query string) (*ItunesResponse, error) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")

	request, err := httpClient.Get(fmt.Sprintf("/search?term=%s", query))
	if err != nil {
		return nil, err
	}

	httpResponse, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	var response *ItunesResponse
	err = json.Unmarshal(httpResponse.Get().Body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
