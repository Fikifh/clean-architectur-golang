package repository

import (
	"errors"

	"incentrick-restful-api/entity"
)

var ErrCourseNotFound = errors.New("Course not found")

type CourseRepository interface {
	Get(id int) (*entity.Courses, error)
	GetAll() ([]*entity.Courses, error)
	Create(in entity.Courses) (*entity.Courses, error)
	Update(in entity.Courses) (*entity.Courses, error)
	Delete(id int) error
}
