package repository

import (
	"errors"

	"incentrick-restful-api/entity"
)

var ErrVideoNotFound = errors.New("Video not found")

type VideoRepository interface {
	Get(id int) (*entity.Video, error)
	GetAll() ([]*entity.Video, error)
	Create(in entity.Video) (*entity.Video, error)
	Update(in entity.Video) (*entity.Video, error)
	Delete(id int) error
}
