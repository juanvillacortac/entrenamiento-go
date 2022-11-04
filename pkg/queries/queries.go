package queries

import (
	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
	"github.com/juanvillacortac/entrenamiento-go/pkg/fetchers"
	"gorm.io/gorm/clause"
)

func QuerySongs(params fetchers.Params) []entities.Song {
	songs := []entities.Song{}
	query := db.DB.Preload("Songs")
	if params.Name == "" && params.Album == "" && params.Artist == "" {
		return songs
	}
	if params.Name != "" {
		query = query.Where("Name LIKE ?", params.Name+"%")
	}
	if params.Album != "" {
		query = query.Where("album LIKE ?", params.Album+"%")
	}
	if params.Artist != "" {
		query = query.Where("artist LIKE ?", params.Artist+"%")
	}
	query.Find(&songs)
	if len(songs) == 0 {
		songs, _ = fetchers.RetrieveFromApis(params)
		db.DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&songs)
		query.Find(&songs)
	}
	return songs
}
