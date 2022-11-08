package controllers

import (
	"strings"

	"github.com/juanvillacortac/entrenamiento-go/pkg/api"
	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
	"gorm.io/gorm/clause"
)

func GetSongsByParams(params api.Params) []entities.Song {
	songs := []entities.Song{}
	query := db.DB

	if params.Name == "" && params.Album == "" && params.Artist == "" {
		return songs
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

	return songs
}
