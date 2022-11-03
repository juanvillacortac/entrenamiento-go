package fetchers

import "github.com/juanvillacortac/entrenamiento-go/pkg/entities"

type ApiResponse interface {
	Transform() []entities.Song
}
