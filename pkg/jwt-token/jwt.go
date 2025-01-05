package jwttoken

import (
	"fmt"
	"orderly/internal/infrastructure/config"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrTokenExpired        = jwt.ErrTokenExpired
	ErrTokenIsInvalid      = fmt.Errorf("jwt token is invalid")
	ErrNoCustomClaimsFound = fmt.Errorf("error while parsing token, no custom claims found")
)

type tokenData struct {
	Id       uuid.UUID
	Role     string
	AddlInfo interface{}
}

type CustomClaims struct {
	tokenData
	jwt.RegisteredClaims
}

func GenerateToken(id uuid.UUID, role string, addlInfo map[string]interface{}, JwtExpTimeInMinutes time.Duration) (string, error) {

	//create a custom claim
	claims := &CustomClaims{
		tokenData: tokenData{
			Id:       id,
			Role:     role,
			AddlInfo: addlInfo,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JwtExpTimeInMinutes)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Configs.Env.JwtSecretKey))

	return tokenString, err
}

func GetDataFromToken(tokenString string) (tokenData *tokenData, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Configs.Env.JwtSecretKey), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), jwt.ErrTokenExpired.Error()) {
			return nil, ErrTokenExpired
		}
		return nil, fmt.Errorf("jwt token is invalid or error in parsing. (error: %v)", err)

	}
	if !token.Valid {
		return nil, ErrTokenIsInvalid
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		//check if token expired
		if claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, ErrTokenExpired
		}
		return &claims.tokenData, nil
	} else {
		return nil, ErrNoCustomClaimsFound
	}
}
