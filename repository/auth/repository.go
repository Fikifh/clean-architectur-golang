package auth

import (
	"encoding/json"
	"errors"
	repository_auth "incentrick-restful-api/app/repository/authentication"
	"incentrick-restful-api/entity"
	entity_authentication "incentrick-restful-api/entity/Authentication"
	"io"
	"net/http"

	"gorm.io/gorm"
)

type repository struct {
	db        *gorm.DB
	tableName string
}

func New(db *gorm.DB) repository_auth.AuthRepository {
	return &repository{db, "users"}
}

func (r *repository) Login(googleClaims *entity_authentication.GoogleClaims) (*entity.UserWithAuth, error) {
	user := UserWithAuthModel{}
	err := r.db.Table(r.tableName).Where("email = ?", googleClaims.Email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user.ToAuthUser(), nil
}

func (r *repository) Register(googleClaims *entity_authentication.GoogleClaims) (*entity.UserWithAuth, error) {
	userAuth := UserWithAuthModel{}
	userId := r.db.Table(r.tableName).Create(&entity.User{Email: googleClaims.Email, Name: googleClaims.FirstName + googleClaims.LastName})
	err := r.db.Table(r.tableName).Where("id = ?", userId).First(&userAuth).Error
	if err != nil {
		return nil, err
	}
	return userAuth.ToAuthUser(), nil
}

func (r *repository) CreateOrUpdateAuth(accessToken string, refreshToken string, agent string, userid int) (*entity.UserWithAuth, error) {
	authModel := &entity_authentication.Authentication{}
	r.db.Where("platform = ? AND user_id", agent, userid).First(&authModel)
	if authModel == nil {
		err := r.db.Create(&entity_authentication.Authentication{AccessToken: accessToken, RefreshToken: refreshToken, Platform: agent}).Error
		if err != nil {
			return nil, err
		}
		authModel = &entity_authentication.Authentication{}
	} else {
		authModel.AccessToken = accessToken
		authModel.RefreshToken = refreshToken
		authModel.IsRevoked = false
		authModel.Platform = agent
		r.db.Save(&authModel)
	}
	r.db.Where("id = ?", userid).First(&entity.UserWithAuth{})
	return &entity.UserWithAuth{}, nil
}

func (r *repository) GetGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}
