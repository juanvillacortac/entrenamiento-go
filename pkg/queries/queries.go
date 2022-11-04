package queries

import (
	"time"

	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
	"github.com/juanvillacortac/entrenamiento-go/pkg/fetchers"
	"gorm.io/gorm/clause"
)

func QuerySongs(params fetchers.Params, useCache bool) (entities.Songs, error) {
	if useCache {
		songs, err := QuerySongsWithCache(params)
		return *songs, err
	}
	songs := []entities.Song{}
	query := db.DB.Preload("Songs")
	if params.Name == "" && params.Album == "" && params.Artist == "" {
		return songs, nil
	}
	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if params.Album != "" {
		query = query.Where("album LIKE ?", "%"+params.Album+"%")
	}
	if params.Artist != "" {
		query = query.Where("artist LIKE ?", "%"+params.Artist+"%")
	}
	query.Find(&songs)
	if len(songs) == 0 {
		songs, err := fetchers.RetrieveFromApis(params)
		if err != nil {
			return songs, err
		}
		db.DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&songs)
		query.Find(&songs)
	}
	return songs, nil
}

func QuerySongsWithCache(params fetchers.Params) (*entities.Songs, error) {
	hash := params.Hash()
	songs := entities.Songs{}
	cached, _ := db.RDB.Get(db.RCtx, hash).Result()
	if cached != "" {
		songs, err := entities.UnmarshalSongs(cached)
		if err != nil {
			return &songs, err
		}
		return &songs, nil
	}
	songs, err := QuerySongs(params, false)
	if err != nil {
		return &songs, err
	}
	db.RDB.Set(db.RCtx, hash, songs.String(), 60*time.Minute)
	return &songs, nil
}
