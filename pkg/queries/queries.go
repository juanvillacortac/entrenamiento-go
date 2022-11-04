package queries

import (
	"strings"

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
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(params.Name)+"%")
	}
	if params.Album != "" {
		query = query.Where("LOWER(album) LIKE ?", "%"+strings.ToLower(params.Album)+"%")
	}
	if params.Artist != "" {
		query = query.Where("LOWER(artist) LIKE ?", "%"+strings.ToLower(params.Artist)+"%")
	}
	query.Order(clause.OrderByColumn{
		Column: clause.Column{
			Name: "name",
		},
		Desc: true,
	}).Find(&songs)
	go SyncSongsDB(params)
	return songs, nil
}

func SyncSongsDB(params fetchers.Params) (entities.Songs, error) {
	songs, err := fetchers.RetrieveFromApis(params)
	if err != nil {
		return entities.Songs{}, err
	}
	db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&songs)
	return songs, err
}

func CacheSongs(params fetchers.Params) (entities.Songs, error) {
	songs, err := QuerySongs(params, false)
	if err != nil {
		return entities.Songs{}, err
	}
	db.RDB.Set(db.RCtx, params.Hash(), songs.String(), 0)
	return songs, nil
}

func QuerySongsWithCache(params fetchers.Params) (*entities.Songs, error) {
	cached, _ := db.RDB.Get(db.RCtx, params.Hash()).Result()
	if cached != "" {
		songs := entities.Songs{}
		songs, err := entities.UnmarshalSongs(cached)
		if err != nil {
			return &songs, err
		}
		go CacheSongs(params)
		return &songs, nil
	}
	songs, err := CacheSongs(params)
	return &songs, err
}
