package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sktandon0121/backend/models"
)

const (
	SECRET string = "KPM1XW0Qkcpxw8aUkYFH3e3XZOTVs5lp"
)

func GetJWTTokenInstance(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method confirm to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func VerifyToken(tokenString string) (bool, error) {
	tokenInstance, err := GetJWTTokenInstance(tokenString)
	if err != nil {
		return false, err
	}
	if err := tokenInstance.Claims.Valid(); err != nil {
		return false, err
	}
	return true, nil
}

func GetTokenMetadata(tokenString string) (*models.Claim, error) {
	token, err := GetJWTTokenInstance(tokenString)
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(*models.Claim)
	if !ok && !token.Valid {
		return nil, errors.New("error getting token metadata")
	}
	return claim, nil
}

func CreateToken(userid int) (string, error) {
	var err error
	//Creating Access Token
	atClaims := models.Claim{
		Authorized: true,
		UserId:     userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    "Subodh Tandon",
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return token, nil
}

func AddDataToContext(ctx context.Context, data *models.Claim, token string) context.Context {
	ctx1 := context.WithValue(ctx, "userId", data.UserId)
	ctx2 := context.WithValue(ctx1, "token", token)
	return ctx2
}

func GetUserFromContext(ctx context.Context) int {
	var userId int
	if ctx == nil {
		return userId
	}
	value := ctx.Value("userId")
	if value == nil {
		return userId
	}
	userId = value.(int)
	return userId
}
