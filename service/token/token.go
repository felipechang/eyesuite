package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/hardcake/eyesuite/service/storage"
	"net/http"
	"strings"
	"time"
)

type Token interface {
	ExtractHeaderToken(r *http.Request) string
	ExtractUserKey(accessToken string, accessSecret string) (string, error)
	ApplyTokenValues(key string, user *storage.User) error
}

type token struct {
	AccessSecret  string
	RefreshSecret string
}

func NewToken(AccessSecret string, RefreshSecret string) Token {
	return &token{
		AccessSecret:  AccessSecret,
		RefreshSecret: RefreshSecret,
	}
}

type AccessClaim struct {
	AccessUuid string `json:"access_uuid"`
	UserId     string `json:"user_id"`
	Exp        int64  `json:"exp"`
	jwt.MapClaims
}

type RefreshClaim struct {
	RefreshUuid string `json:"refresh_uuid"`
	UserId      string `json:"user_id"`
	Exp         int64  `json:"exp"`
	jwt.MapClaims
}

type Pair struct {
	AccessToken  string `form:"access_token" json:"access_token"`
	RefreshToken string `form:"refresh_token" json:"refresh_token"`
}

// ExtractHeaderToken get Authorization key from the request header
func (t *token) ExtractHeaderToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// ExtractUserKey get user key from access token
func (t *token) ExtractUserKey(accessToken string, accessSecret string) (string, error) {
	tk, err := verifyToken(accessToken, accessSecret)
	if err != nil {
		return "", err
	}
	s, ok := tk.Claims.(jwt.MapClaims)
	if ok == false {
		return "", errors.New("could not parse claim")
	}
	err = s.Valid()
	if err != nil {
		return "", err
	}
	return s["user_id"].(string), nil
}

// ApplyTokenValues set new token values on user
func (t *token) ApplyTokenValues(key string, user *storage.User) error {
	user.AtExpires = time.Now().Add(time.Hour * 10).Unix()
	user.TokenUuid = uuid.NewV4().String()

	user.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	user.RefreshUuid = user.TokenUuid + "++" + key

	//Creating Access Claim
	ac := AccessClaim{}
	ac.AccessUuid = user.TokenUuid
	ac.UserId = key
	ac.Exp = user.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
	accessToken, err := at.SignedString([]byte(t.AccessSecret))
	if err != nil {
		return err
	}
	user.AccessToken = accessToken

	//Creating Refresh Claim
	user.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	user.RefreshUuid = user.TokenUuid + "++" + key

	rc := RefreshClaim{}
	rc.RefreshUuid = user.RefreshUuid
	rc.UserId = key
	rc.Exp = user.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rc)

	refreshToken, err := rt.SignedString([]byte(t.RefreshSecret))
	if err != nil {
		return err
	}
	user.RefreshToken = refreshToken

	return nil
}

// verifyToken token has a valid signature
func verifyToken(tokenString string, secret string) (*jwt.Token, error) {
	tk, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return tk, nil
}
