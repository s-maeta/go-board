package service

import (
	"board/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, error) {
	fmt.Printf("userid:%s", userId)
	// configファイルからシークレットキーを取得する
	config := config.GetConfig()
	secretKey := config.Auth.SecretKey

	// トークンの有効期限を取得する
	tokenLifeTime := config.Auth.TokenLifetime

	//クレームを生成する
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Unix(),
	}

	// トークンを生成し、文字列に変換する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
