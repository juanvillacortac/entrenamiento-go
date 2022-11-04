package entities

type SongOrigin string

const (
	Itunes SongOrigin = "itunes"
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
