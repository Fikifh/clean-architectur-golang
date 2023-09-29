package crud_course

import (
	"errors"
	"fmt"

	"incentrick-restful-api/app/repository"
	"incentrick-restful-api/entity"
)

type usecase struct {
	course_repo repository.CourseRepository
}

func NewUseCase(course_repo repository.CourseRepository) UseCase {
	return &usecase{course_repo: course_repo}
}

func (uc *usecase) Get(id int) (*entity.Courses, error) {
	data, err := uc.course_repo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrCourseNotFound) {
			return nil, repository.ErrCourseNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

func (uc *usecase) GetAll() ([]*entity.Courses, error) {
	data, err := uc.course_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}
func (uc *usecase) Create(in entity.Courses) (*entity.Courses, error) {
	data, err := uc.course_repo.Create(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

func (uc *usecase) Update(in entity.Courses) (*entity.Courses, error) {
	data, err := uc.course_repo.Update(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil

}

func (uc *usecase) Delete(id int) error {
	err := uc.course_repo.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return err
}
