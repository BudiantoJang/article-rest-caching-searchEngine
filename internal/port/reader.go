package port

import (
	"jang-article/internal/model"
)

type ConfigReader interface {
	Read() (model.Config, error)
}
