package user_repository

import (
	"gorm.io/gorm"
)

type repositoryStruct struct {
	db        *gorm.DB
	tableName string
}

// func New(db *gorm.DB) repository.UserRepository {
// 	return &repositoryStruct{db, "users"}
// }

// func (r *repositoryStruct) Get(id int) (*entity.User, error) {
// 	resp := UserModel{}
// 	err := r.db.Table(r.tableName).Where("id = ?", id).First(&resp).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, irepository.ErrUserNotFound
// 		}
// 		return nil, err
// 	}
// 	return resp.ToUserEntity(), nil
// }

// func (r *repository) Create(in entity.Video) (*entity.Video, error) {
// 	videoModel := VideoModel{}.FromVideoEntity(in)

// 	timeNow := time.Now()
// 	videoModel.CreatedAt = &timeNow
// 	videoModel.UpdatedAt = &timeNow

// 	err := r.db.Table(r.tableName).Create(&videoModel).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return videoModel.ToVideoEntity(), nil

// }
// func (r *repository) Update(in entity.Video) (*entity.Video, error) {
// 	videoModel := VideoModel{}.FromVideoEntity(in)
// 	_, err := r.Get(in.Id)
// 	if errors.Is(err, irepository.ErrVideoNotFound) {
// 		return nil, nil
// 	}

// 	timeNow := time.Now()
// 	videoModel.CreatedAt = nil
// 	videoModel.UpdatedAt = &timeNow
// 	err = r.db.Table(r.tableName).Where("id = ?", in.Id).Updates(&videoModel).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return videoModel.ToVideoEntity(), nil

// }
