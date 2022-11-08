package entities

import "encoding/json"

type SongOrigin string

const (
	Itunes      SongOrigin = "itunes"
	ChartLyrics SongOrigin = "chartlyrics"
)

type Song struct {
	ID       string     `json:"id" gorm:"primary_key"`
	Name     string     `json:"name"`
	Artist   string     `json:"artist"`
	Duration string     `json:"duration"`
	Album    string     `json:"album"`
	Artwork  string     `json:"artwork"`
	Price    string     `json:"price"`
	Origin   SongOrigin `json:"origin"`
}

type Songs []Song

func (songs Songs) String() string {
	b, _ := json.Marshal(songs)
	return string(b)
}

func UnmarshalSongs(str string) (Songs, error) {
	var songs Songs
	err := json.Unmarshal([]byte(str), &songs)
	return songs, err
}
