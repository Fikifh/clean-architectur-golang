package course_repository_v1

import (
	"errors"
	"time"

	irepository "incentrick-restful-api/app/repository"
	"incentrick-restful-api/entity"

	"gorm.io/gorm"
)

type repository struct {
	db        *gorm.DB
	tableName string
}

func New(db *gorm.DB) irepository.CourseRepository {
	return &repository{db, "courses"}
}

func (r *repository) Get(id int) (*entity.Courses, error) {
	resp := CourseModel{}
	err := r.db.Table(r.tableName).Where("id = ?", id).First(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, irepository.ErrCourseNotFound
		}
		return nil, err
	}
	return resp.ToCourseEntity(), nil
}

func (r *repository) GetAll() ([]*entity.Courses, error) {
	datas := []CourseModel{}
	err := r.db.Table(r.tableName).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	resp := []*entity.Courses{}

	for _, data := range datas {
		resp = append(resp, data.ToCourseEntity())
	}
	return resp, nil

}
func (r *repository) Create(in entity.Courses) (*entity.Courses, error) {
	CourseModel := CourseModel{}.FromCourseEntity(in)

	timeNow := time.Now()
	CourseModel.CreatedAt = &timeNow
	CourseModel.UpdatedAt = &timeNow

	err := r.db.Table(r.tableName).Create(&CourseModel).Error
	if err != nil {
		return nil, err
	}
	return CourseModel.ToCourseEntity(), nil

}
func (r *repository) Update(in entity.Courses) (*entity.Courses, error) {
	CourseModel := CourseModel{}.FromCourseEntity(in)
	_, err := r.Get(in.Id)
	if errors.Is(err, irepository.ErrCourseNotFound) {
		return nil, nil
	}

	timeNow := time.Now()
	CourseModel.CreatedAt = nil
	CourseModel.UpdatedAt = &timeNow
	err = r.db.Table(r.tableName).Where("id = ?", in.Id).Updates(&CourseModel).Error
	if err != nil {
		return nil, err
	}

	return CourseModel.ToCourseEntity(), nil

}
func (r *repository) Delete(id int) error {
	return r.db.Table(r.tableName).Delete(&CourseModel{}, "id = ?", id).Error
}
