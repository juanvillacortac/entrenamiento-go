package query

import (
	"fmt"

	"github.com/juanvillacortac/entrenamiento-go/pkg/api"
	"github.com/juanvillacortac/entrenamiento-go/pkg/controllers"
	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
	"gorm.io/gorm/clause"
)

func QuerySongs(params api.Params, useCache bool) (entities.Songs, error) {
	if useCache {
		songs, err := QuerySongsWithCache(params)
		return *songs, err
	}
	songs := controllers.GetSongsByParams(params)
	if len(songs) == 0 {
		songs, err := SyncSongsDB(params)
		return songs, err
	}
	go SyncSongsDB(params)
	return songs, nil
}

func SyncSongsDB(params api.Params) (entities.Songs, error) {
	fmt.Println("Syncing database with APIs...")
	songs, err := api.RetrieveFromApis(params)
	if err != nil {
		return entities.Songs{}, err
	}
	db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&songs)
	fmt.Println("Database synced")
	return songs, err
}

func CacheSongs(params api.Params) (entities.Songs, error) {
	songs, err := QuerySongs(params, false)
	if err != nil {
		return entities.Songs{}, err
	}
	db.RDB.Set(db.RCtx, params.Hash(), songs.String(), 0)
	return songs, nil
}

func QuerySongsWithCache(params api.Params) (*entities.Songs, error) {
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
