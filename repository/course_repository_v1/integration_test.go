package course_repository_v1

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"incentrick-restful-api/entity"
	"incentrick-restful-api/pkg/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	fmt.Println("init db for test")
	var err error

	loop := 0
	maxLoop := 100
	for {
		db, err = mysql.Connect("localhost", "33061", "root", "password", "gca", logger.LogLevel(3))
		if err == nil {
			break
		}
		loop++
		if loop > maxLoop {
			panic("Failed to init db, 100 times test already")
		}
		time.Sleep(time.Second)
	}
}

func Test_repository_Get(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		tableName string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Courses
		wantErr bool
	}{
		{
			name:   "success: standart scenario",
			fields: fields{db: db, tableName: "courses"},
			args:   args{id: 1},
			want: &entity.Courses{
				Id:   1,
				Name: "Learn Go Programming - Golang Tutorial for Beginners",
			},
			wantErr: false,
		},
		{
			name:    "failed: not found",
			fields:  fields{db: db, tableName: "courses"},
			args:    args{id: 100},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "failed: wrong tablename (human/dev error)",
			fields:  fields{db: db, tableName: "courses2"},
			args:    args{id: 100},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				tableName: tt.fields.tableName,
			}
			got, err := r.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetAll(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		tableName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Courses
		wantErr bool
	}{
		{
			name:   "success: standart scenario",
			fields: fields{db: db, tableName: "courses"},
			want: []*entity.Courses{
				&entity.Courses{Id: 1, Name: "Learn Go Programming - Golang Tutorial for Beginners"},
				&entity.Courses{Id: 2, Name: "7 Habits of Highly Effective Programmers (ft. ex-Google TechLead)"},
				&entity.Courses{Id: 3, Name: "Building a Bank with Go"},
				&entity.Courses{Id: 4, Name: "ITkonekt 2019 | Robert C. Martin (Uncle Bob), Clean Architecture and Design"},
				&entity.Courses{Id: 5, Name: "The Principles of Clean Architecture by Uncle Bob Martin"},
				&entity.Courses{Id: 6, Name: "Making Architecture Matter - Martin Fowler Keynote"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				tableName: tt.fields.tableName,
			}
			got, err := r.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, g := range got {
				fmt.Println(*g)
				// if !reflect.DeepEqual(g, tt.want[i]) {
				// 	t.Errorf("repository.GetAll() = %v, want %v", got, tt.want)
				// }
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Create(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		tableName string
	}
	type args struct {
		in entity.Courses
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Courses
		wantErr bool
	}{
		{
			name:    "success: standart scenario",
			fields:  fields{db: db, tableName: "courses"},
			args:    args{in: entity.Courses{Name: "Test Video"}},
			want:    &entity.Courses{Id: 7, Name: "Test Video"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				tableName: tt.fields.tableName,
			}
			got, err := r.Create(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Update(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		tableName string
	}
	type args struct {
		in entity.Courses
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Courses
		wantErr bool
	}{
		{
			name:    "success: standart scenario",
			fields:  fields{db: db, tableName: "courses"},
			args:    args{in: entity.Courses{Id: 2, Name: "Test Update Video"}},
			want:    &entity.Courses{Id: 2, Name: "Test Update Video"},
			wantErr: false,
		},
		{
			name:    "success: nothing to update",
			fields:  fields{db: db, tableName: "courses"},
			args:    args{in: entity.Courses{Id: 101, Name: "Test Update Video"}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				tableName: tt.fields.tableName,
			}
			got, err := r.Update(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Delete(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		tableName string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success: standart scenario",
			fields:  fields{db: db, tableName: "courses"},
			args:    args{id: 6},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				tableName: tt.fields.tableName,
			}
			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("repository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
