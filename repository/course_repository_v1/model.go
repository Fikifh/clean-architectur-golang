package course_repository_v1

import (
	"time"

	"incentrick-restful-api/entity"
)

type CourseModel struct {
	Id        int        `gorm:"primary_key;column:id"`
	Name      string     `gorm:"column:name"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (CourseModel) FromCourseEntity(v entity.Courses) *CourseModel {
	return &CourseModel{
		Id:   v.Id,
		Name: v.Name,
	}
}

func (m *CourseModel) ToCourseEntity() *entity.Courses {
	return &entity.Courses{
		Id:   m.Id,
		Name: m.Name,
	}
}
