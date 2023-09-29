package authentication

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	repository_auth "incentrick-restful-api/app/repository/authentication"
	"incentrick-restful-api/config"
	"incentrick-restful-api/entity"
	entity_authentication "incentrick-restful-api/entity/Authentication"
	utils "incentrick-restful-api/handler/Utils"
	"io"
	"log"
	"net/http"
	"time"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

type usecase struct {
	authRepo repository_auth.AuthRepository
}

func NewUseCase(authRepo repository_auth.AuthRepository) UseCase {
	return &usecase{authRepo: authRepo}
}

func (uc *usecase) Login(idToken string, token string, platform string, w http.ResponseWriter, r *http.Request) (*entity.UserWithAuth, error) {
	googleClaims, err := uc.ValidateGoogleJWTs(idToken)
	if err != nil {
		return nil, err
	}
	login, err := uc.authRepo.Login(googleClaims)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			login, err := uc.authRepo.Register(googleClaims)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
			}
			return login, nil
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	tokens, err := generateJWT(login)
	fmt.Printf("token %v", tokens)
	createToken, err := uc.authRepo.CreateOrUpdateAuth(tokens.AccessToken, tokens.RefreshToken, platform, login.ID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	login = createToken
	return login, nil
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func generateJWT(user *entity.UserWithAuth) (*Tokens, error) {
	var sampleSecretKey = []byte("SecretYouShouldHide")
	accesToken, err := generateToken(user, string(sampleSecretKey), "access_token")
	if err != nil {
		return nil, err
	}
	refreshToken, err := generateToken(user, string(sampleSecretKey), "refresh_token")
	if err != nil {
		return nil, err
	}
	return &Tokens{AccessToken: accesToken, RefreshToken: refreshToken}, nil
}

func generateToken(user *entity.UserWithAuth, secret string, tokenType string) (string, error) {
	var expiryToken int
	if tokenType == "access_token" {
		expiryToken = 24
	} else {
		expiryToken = 7 * 24
	}
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(expiryToken) * time.Hour)
	claims["authorized"] = true
	claims["user"] = user
	claims["token_type"] = tokenType
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uc *usecase) ValidateGoogleJWTs(idToken string) (*entity_authentication.GoogleClaims, error) {
	v := googleAuthIDTokenVerifier.Verifier{}
	aud := utils.GoDotEnvVariable("GOOGLE_CLIENT_ID")
	err := v.VerifyIDToken(idToken, []string{
		aud,
	})

	fmt.Println("client di", aud)
	if err == nil {
		claimSet, err := googleAuthIDTokenVerifier.Decode(idToken)
		if err != nil {
			return nil, errors.New("error di claimset ")
		}
		fmt.Println(claimSet)
		return &entity_authentication.GoogleClaims{
			Email:          "", //payload.Claims["email"].(string),
			FirstName:      "", //payload.Claims["name"].(string),
			LastName:       "",
			Picture:        "",
			StandardClaims: jwt.StandardClaims{},
		}, nil
	}
	return nil, err
}

func verifyIdToken(idToken string) (*entity_authentication.GoogleClaims, error) {
	var token string // this comes from your web or mobile app maybe
	token = idToken
	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Println(token)
	payload, err := tokenValidator.Validate(context.Background(), token, "")
	if err != nil {
		return nil, err
	}
	fmt.Println("PAYLOAD", payload)
	return &entity_authentication.GoogleClaims{
		Email:          "", //payload.Claims["email"].(string),
		FirstName:      "", //payload.Claims["name"].(string),
		LastName:       "",
		Picture:        "",
		StandardClaims: jwt.StandardClaims{},
	}, nil
	// email := payload.Claims["email"]
	// name := payload.Claims["name"]
}

func (uc *usecase) OauthGoogleLogin(c *gin.Context) {

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(c.Writer)
	u := config.GoogleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(c.Writer, c.Request, u, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func (uc *usecase) OauthGoogleCallback(c *gin.Context) {
	// Read oauthState from Cookie
	oauthState, _ := c.Request.Cookie("oauthstate")

	if c.Request.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(c.Request.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	fmt.Fprintf(c.Writer, "UserInfo: %s\n", data)
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := config.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

// func (uc *usecase) Register(*entity.User) {

// }

// func (uc *usecase) Logout(*entity.User) {

// }

// func (uc *usecase) CreateToken(user *entity.User, expiry int) string {
// 	mySigningKey := []byte("AllYourBase")
// 	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
// 	type TokenClaims struct {
// 		Name string `json:"name"`
// 		ID   int    `json:"id"`
// 		jwt.StandardClaims
// 	}

// 	// Create the Claims
// 	claims := &TokenClaims{
// 		Name: user.Name,
// 		ID:   user.ID,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: exp,
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	ss, err := token.SignedString(mySigningKey)
// 	fmt.Printf("%v %v", ss, err)
// 	return ss
// }
