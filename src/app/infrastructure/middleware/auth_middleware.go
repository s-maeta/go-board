package middleware

import (
	"board/app/infrastructure/repository"
	"board/config"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LoginCheckMiddleware(c *gin.Context) {
	// configファイルからシークレットキーを取得する
	config := config.GetConfig()
	secretKey := config.Auth.SecretKey
	// Authorization のヘッダーを取得
	authorizationHeader := c.Request.Header.Get("Authorization")
	// authorizationヘッダーが空白の場合にはリターン
	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
		c.Next()
		return
	}

	ary := strings.Split(authorizationHeader, " ")
	if len(ary) == 2 {
		if ary[0] == "Bearer" {
			// ここまででトークンと思われるものを取り出せたので、解析する
			// ここのロジックは公式サイト参照
			//　https://pkg.go.dev/github.com/golang-jwt/jwt#example-Parse-Hmac
			// token, err := jwt.Parse(ary[1], func(token *jwt.Token) (interface{}, error) {
			// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// 		c.JSON(http.StatusUnauthorized, gin.H{"errors": fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])})
			// 	}
			// 	return []byte(secretKey), nil
			// })

			// トークンを検証
			claims := jwt.MapClaims{}
			// シークレットキーを使用してトークンを検証
			token, err := jwt.ParseWithClaims(ary[1], claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
				return
			}
			// トークンの有効期限を検証
			if !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"errors": "Invalid token"})
				return
			}

			// トークンの有効期限を確認
			expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
			if time.Now().After(expirationTime) {
				c.JSON(http.StatusUnauthorized, gin.H{"errors": "Token has expired"})
				return
			}
			// // 解析に成功したら、ユーザIDを取り出す
			// 	userIDStr := claims["sub"].(string)
			// 	// ユーザIDを使って新しいトークンを取り出す（有効期限切れ対策）
			// 	newToken, err := service.GenerateToken(userIDStr)
			// ユーザーIDを取得
			userID, ok := claims["user_id"].(string)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"errors": "User ID not found in token"})
				return
			}
			// ユーザの情報を取り出す
			userRepository := repository.NewUserRepository()
			user := userRepository.FindForUniqueId(userID)
			if user == nil {
				c.JSON(http.StatusUnauthorized, gin.H{"errors": "User is not found"})
				return
			}
			// 新しいトークンとユーザの情報をコンテキストに保存
			// c.Set("token", newToken)
			c.Set("user", user)
		}
	}
	c.Next()
}
