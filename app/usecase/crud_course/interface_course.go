package crud_course

import "incentrick-restful-api/entity"

type UseCase interface {
	Get(id int) (*entity.Courses, error)
	GetAll() ([]*entity.Courses, error)
	Create(in entity.Courses) (*entity.Courses, error)
	Update(in entity.Courses) (*entity.Courses, error)
	Delete(id int) error
}
